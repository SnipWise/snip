package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/server"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/toolbox/logger"
	"github.com/snipwise/nova/nova-sdk/ui/display"
	"github.com/snipwise/nova/nova-sdk/ui/prompt"
)

// watchStoreFile monitors the store file for changes and reloads it when modified
func watchStoreFile(storePathFile, dataPath, storePath, storeFile string, ragAgent *rag.Agent, metadataExtractorAgent *structured.Agent[KeywordMetadata], checkInterval int) {
	var lastModTime time.Time

	// Get initial modification time
	fileInfo, err := os.Stat(storePathFile)
	if err != nil {
		display.Errorf("❌ Error getting file info for %s: %v", storePathFile, err)
		return
	}
	lastModTime = fileInfo.ModTime()

	display.Infof("👀 Watching store file for changes: %s (checking every %d seconds)", storePathFile, checkInterval)

	ticker := time.NewTicker(time.Duration(checkInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fileInfo, err := os.Stat(storePathFile)
		if err != nil {
			display.Errorf("❌ Error checking file %s: %v", storePathFile, err)
			continue
		}

		if fileInfo.ModTime().After(lastModTime) {
			display.Infof("🔄 Store file changed, reloading: %s", storePathFile)
			lastModTime = fileInfo.ModTime()

			// Reload the store
			err = LoadSnippetData(dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent)
			if err != nil {
				display.Errorf("❌ Error reloading store: %v", err)
			} else {
				display.Infof("✅ Store successfully reloaded")
			}
		}
	}
}

// initializeServerAgent initializes and configures the server agent with all dependencies
func initializeServerAgent() (*server.ServerAgent, error) {
	// Create logger from environment variable
	log := logger.GetLoggerFromEnv()

	envFile := ".env"
	// Load environment variables from env file (optional in Docker)
	if err := godotenv.Load(envFile); err != nil {
		log.Info("Note: .env file not found (using Docker environment variables)\n")
	}

	ctx := context.Background()

	storePath := env.GetEnvOrDefault("STORE_PATH", "./store")
	storeFile := env.GetEnvOrDefault("STORE_FILE", "golang.snippets.store.json")
	dataPath := env.GetEnvOrDefault("DATA_PATH", "./data")

	httpPort := env.GetEnvOrDefault("HTTP_PORT", "3500")

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:12434/engines/llama.cpp/v1")
	ragModelId := env.GetEnvOrDefault("RAG_EMBEDDING_MODEL_ID", "ai/embeddinggemma:latest")
	metadataModelId := env.GetEnvOrDefault("METADATA_MODEL_ID", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	ragAgent, err := CreateRagAgent(ctx, engineURL, ragModelId)
	if err != nil {
		return nil, err
	}

	metadataExtractorAgent, err := CreateMetadataExtractorAgent(ctx, engineURL, metadataModelId)
	if err != nil {
		return nil, err
	}

	err = LoadSnippetData(dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent)
	if err != nil {
		return nil, err
	}

	// Start watching the store file for changes in a goroutine if enabled
	enableStoreWatch := env.GetEnvBoolOrDefault("ENABLE_STORE_WATCH", true)
	if enableStoreWatch {
		storePathFile := filepath.Join(storePath, storeFile)
		storeWatchInterval := env.GetEnvIntOrDefault("STORE_WATCH_INTERVAL", 5)
		go watchStoreFile(storePathFile, dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent, storeWatchInterval)
	}

	compressorAgent, err := CreateCompressorAgent(ctx, engineURL)
	if err != nil {
		return nil, err
	}

	// ------------------------------------------------
	// Create the server agent
	// ------------------------------------------------
	temperature := env.GetEnvFloatOrDefault("CODER_TEMPERATURE", 0.8)
	topP := env.GetEnvFloatOrDefault("CODER_TOP_P", 0.9)
	coderAgentModelID := env.GetEnvOrDefault("CODER_MODEL_ID", "hf.co/menlo/jan-nano-gguf:q4_k_m")
	directivesPathFile := env.GetEnvOrDefault("CODER_DIRECTIVES_PATH_FILE", "./directives/coder.agent.system.instructions.md")
	coderAgentSystemInstructionsContent, err := files.ReadTextFile(directivesPathFile)
	if err != nil {
		return nil, err
	}

	serverAgent, err := server.NewAgent(
		ctx,
		agents.Config{
			Name:                    "coder",
			EngineURL:               engineURL,
			SystemInstructions:      coderAgentSystemInstructionsContent,
			KeepConversationHistory: true,
		},
		models.Config{
			Name:        coderAgentModelID,
			Temperature: models.Float64(temperature),
			TopP:        models.Float64(topP),
		},
		":"+httpPort,
	)
	if err != nil {
		display.Errorf("❌ Error creating server agent: %v", err)
		return nil, err
	}
	display.Infof("🚀 SNIP server agent created and listening on port %s", serverAgent.GetPort())

	contextSizeLimit := env.GetEnvIntOrDefault("CONTEXT_SIZE_LIMIT", 32000)

	serverAgent.SetCompressorAgent(compressorAgent)
	serverAgent.SetContextSizeLimit(contextSizeLimit)

	serverAgent.SetRagAgent(ragAgent)

	similarityLimit := env.GetEnvFloatOrDefault("SIMILARITY_LIMIT", 0.4)
	maxSimilarities := env.GetEnvIntOrDefault("MAX_SIMILARITIES", 7)

	serverAgent.SetSimilarityLimit(similarityLimit)
	serverAgent.SetMaxSimilarities(maxSimilarities)

	return serverAgent, nil
}

func chatWithAgent() {
	chatAgent, err := initializeServerAgent()
	if err != nil {
		return
	}

	for {

		markdownParser := display.NewMarkdownChunkParser()

		input := prompt.NewWithColor("🤖 Ask me something? [" + chatAgent.GetName() + "]")
		question, err := input.RunWithEdit()

		if err != nil {
			display.Errorf("failed to get input: %v", err)
			return
		}
		if strings.HasPrefix(question, "/bye") {
			display.Infof("👋 Goodbye!")
			break
		}

		if strings.HasPrefix(question, "/messages") {
			display.Infof("💬 Current conversation messages:")
			for i, msg := range chatAgent.GetMessages() {
				display.Infof("Message %d - Role: %s, Content: \n%s", i, msg.Role, msg.Content)
				display.Separator()
			}
			continue
		}

		if strings.HasPrefix(question, "/reset") {
			display.Infof("🔄 Resetting %s context", chatAgent.GetName())
			chatAgent.ResetMessages()
			continue
		}

		display.NewLine()

		result, err := chatAgent.StreamCompletion(question,
			func(chunk string, finishReason string) error {

				// Use markdown chunk parser for colorized streaming output
				if chunk != "" {
					display.MarkdownChunk(markdownParser, chunk)
				}
				if finishReason == "stop" {
					markdownParser.Flush()
					markdownParser.Reset()
					//markdownParser.Flush()
					display.NewLine()
				}
				return nil
			})

		if err != nil {
			display.Errorf("[%s][%v]failed to get completion: %v", chatAgent.GetName(), chatAgent.GetContextSize(), err)
			return
		}

		display.NewLine()
		display.Separator()
		display.KeyValue("Finish reason", result.FinishReason)
		display.KeyValue("Context size", fmt.Sprintf("%d characters", chatAgent.GetContextSize()))
		display.Separator()
	}
}

func indexDocuments(dataPath, storeFilePath string) {
	// ex of storeFilePath: ./store/golang.snippets.store.json
	ctx := context.Background()

	storePath := filepath.Dir(storeFilePath)
	storeFile := filepath.Base(storeFilePath)

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:12434/engines/llama.cpp/v1")
	ragModelId := env.GetEnvOrDefault("RAG_EMBEDDING_MODEL_ID", "ai/embeddinggemma:latest")
	metadataModelId := env.GetEnvOrDefault("METADATA_MODEL_ID", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	ragAgent, err := CreateRagAgent(ctx, engineURL, ragModelId)
	if err != nil {
		display.Errorf("❌ Error creating RAG agent: %v", err)
		return
	}

	metadataExtractorAgent, err := CreateMetadataExtractorAgent(ctx, engineURL, metadataModelId)
	if err != nil {
		display.Errorf("❌ Error creating metadata extractor agent: %v", err)
		return
	}

	err = LoadSnippetData(dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent)
	if err != nil {
		display.Errorf("❌ Error loading snippet data: %v", err)
		return
	}

}

// addToIndex adds a single document to the existing index
func addToIndex(dataFilePath, storeFilePath string) {
	// dataFilePath: ./data/snippet1.md
	// ex of storeFilePath: ./store/golang.snippets.store.json
	ctx := context.Background()

	storePath := filepath.Dir(storeFilePath)
	storeFile := filepath.Base(storeFilePath)

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:12434/engines/llama.cpp/v1")
	ragModelId := env.GetEnvOrDefault("RAG_EMBEDDING_MODEL_ID", "ai/embeddinggemma:latest")
	metadataModelId := env.GetEnvOrDefault("METADATA_MODEL_ID", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	ragAgent, err := CreateRagAgent(ctx, engineURL, ragModelId)
	if err != nil {
		display.Errorf("❌ Error creating RAG agent: %v", err)
		return
	}

	metadataExtractorAgent, err := CreateMetadataExtractorAgent(ctx, engineURL, metadataModelId)
	if err != nil {
		display.Errorf("❌ Error creating metadata extractor agent: %v", err)
		return
	}

	// Add the document to the existing store
	err = AddToSnippetData(dataFilePath, storePath, storeFile, ragAgent, metadataExtractorAgent)
	if err != nil {
		display.Errorf("❌ Error adding document to index: %v", err)
		return
	}

	display.Infof("✅ Document successfully added to index: %s", dataFilePath)
}

func serveAgent() {
	serverAgent, err := initializeServerAgent()
	if err != nil {
		return
	}

	// Start the server agent
	if err := serverAgent.StartServer(); err != nil {
		display.Errorf("❌ Error starting crew server agent: %v", err)
		return
	}
}

func main() {
	// Define command-line flags
	serveCmd := flag.Bool("serve", false, "Start the server agent")
	chatCmd := flag.Bool("chat", false, "Start chat mode")
	indexCmd := flag.Bool("index", false, "Index documents")
	addToIndexCmd := flag.Bool("add-to-index", false, "Add a document to the index")

	flag.Parse()

	// Get remaining arguments after flags
	args := flag.Args()

	// Determine which command to execute
	switch {
	case *chatCmd:
		chatWithAgent()
	case *indexCmd:
		dataPath := "./data"
		storeFilePath := "./store/golang.snippets.store.json"
		if len(args) >= 1 {
			dataPath = args[0]
		}
		if len(args) >= 2 {
			storeFilePath = args[1]
		}
		indexDocuments(dataPath, storeFilePath)
	case *addToIndexCmd:
		if len(args) < 2 {
			fmt.Println("Usage: --add-to-index <dataFilePath> <storeFilePath>")
			os.Exit(1)
		}
		dataFilePath := args[0]
		storeFilePath := args[1]
		addToIndex(dataFilePath, storeFilePath)
	case *serveCmd:
		serveAgent()
	default:
		// Default behavior is to serve
		serveAgent()
	}
}
