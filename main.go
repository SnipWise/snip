package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/agents/compressor"
	"github.com/snipwise/nova/nova-sdk/agents/crew"
	"github.com/snipwise/nova/nova-sdk/agents/crewserver"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
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

func initializeSideKickAgents(ctx context.Context, engineURL, dataPath, storePath, storeFile string) (*rag.Agent, *structured.Agent[KeywordMetadata], *compressor.Agent, error) {
	// Create logger from environment variable
	log := logger.GetLoggerFromEnv()

	ragModelId := env.GetEnvOrDefault("RAG_EMBEDDING_MODEL_ID", "ai/embeddinggemma:latest")
	metadataModelId := env.GetEnvOrDefault("METADATA_MODEL_ID", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	ragAgent, err := CreateRagAgent(ctx, engineURL, ragModelId)

	metadataExtractorAgent, err := CreateMetadataExtractorAgent(ctx, engineURL, metadataModelId)

	compressorAgent, err := CreateCompressorAgent(ctx, engineURL)

	if err != nil {
		log.Error("❌ Error creating sidekick agents: %v", err)
		//display.Errorf("❌ Error creating sidekick agents: %v", err)
		return nil, nil, nil, err
	}
	err = LoadSnippetData(dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent)
	if err != nil {
		log.Error("❌ Error loading snippet data: %v", err)
		//display.Errorf("❌ Error loading snippet data: %v", err)
		return nil, nil, nil, err
	}
	return ragAgent, metadataExtractorAgent, compressorAgent, nil

}

func initializeCoderChatAgent(ctx context.Context, engineURL string) (*chat.Agent, error) {
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

	userMessagePreDirectives := env.GetEnvOrDefault("USER_MESSAGE_PRE_DIRECTIVES", "")
	userMessagePostDirectives := env.GetEnvOrDefault("USER_MESSAGE_POST_DIRECTIVES", "")


	coderChatAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:                    "coder-chat",
			EngineURL:               engineURL,
			SystemInstructions:      coderAgentSystemInstructionsContent,
			KeepConversationHistory: true,
		},
		models.Config{
			Name:        coderAgentModelID,
			Temperature: models.Float64(temperature),
			TopP:        models.Float64(topP),
		},
	)
	if err != nil {
		display.Errorf("❌ Error creating coder chat agent: %v", err)
		return nil, err
	}
	if userMessagePreDirectives != "" {
		coderChatAgent.SetUserMessagePreDirectives(userMessagePreDirectives)
		display.Infof("📘 Set user message pre directives: %s", userMessagePreDirectives)
	}
	if userMessagePostDirectives != "" {
		coderChatAgent.SetUserMessagePostDirectives(userMessagePostDirectives)
		display.Infof("📙 Set user message post directives: %s", userMessagePostDirectives)
	}

	display.Infof("🚀 Coder chat agent created")

	return coderChatAgent, nil
}

type Configuration struct {
	EngineURL          string
	StorePath          string
	StoreFile          string
	DataPath           string
	HTTPPort           int
	EnableStoreWatch   bool
	StoreWatchInterval int
	SimilarityLimit    float64
	MaxSimilarities    int
	ContextSizeLimit   int
}

func loadConfiguration() Configuration {
	return Configuration{
		EngineURL:          env.GetEnvOrDefault("ENGINE_URL", "http://localhost:12434/engines/llama.cpp/v1"),
		StorePath:          env.GetEnvOrDefault("STORE_PATH", "./store"),
		StoreFile:          env.GetEnvOrDefault("STORE_FILE", "golang.snippets.store.json"),
		DataPath:           env.GetEnvOrDefault("DATA_PATH", "./data"),
		HTTPPort:           env.GetEnvIntOrDefault("HTTP_PORT", 3500),
		EnableStoreWatch:   env.GetEnvBoolOrDefault("ENABLE_STORE_WATCH", true),
		StoreWatchInterval: env.GetEnvIntOrDefault("STORE_WATCH_INTERVAL", 5),
		SimilarityLimit:    env.GetEnvFloatOrDefault("SIMILARITY_LIMIT", 0.4),
		MaxSimilarities:    env.GetEnvIntOrDefault("MAX_SIMILARITIES", 7),
		ContextSizeLimit:   env.GetEnvIntOrDefault("CONTEXT_SIZE_LIMIT", 32000),
	}
}

// initializeServerAgent initializes and configures the server agent with all dependencies
func initializeServerAgent() (*crewserver.CrewServerAgent, error) {
	// Create logger from environment variable
	log := logger.GetLoggerFromEnv()
	config := loadConfiguration()

	ctx := context.Background()

	httpPort := config.HTTPPort
	engineURL := config.EngineURL

	storePath := config.StorePath
	storeFile := config.StoreFile
	dataPath := config.DataPath

	ragAgent, metadataExtractorAgent, compressorAgent, err := initializeSideKickAgents(ctx, engineURL, dataPath, storePath, storeFile)
	if err != nil {
		log.Error("❌ Error creating sidekick agents: %v", err)
		return nil, err
	}

	coderAgent, err := initializeCoderChatAgent(ctx, engineURL)
	if err != nil {
		log.Error("❌ Error creating coder chat agent: %v", err)
		return nil, err
	}

	// Start watching the store file for changes in a goroutine if enabled
	enableStoreWatch := config.EnableStoreWatch
	if enableStoreWatch {
		storePathFile := filepath.Join(storePath, storeFile)
		storeWatchInterval := config.StoreWatchInterval
		go watchStoreFile(storePathFile, dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent, storeWatchInterval)
	}

	similarityLimit := config.SimilarityLimit
	maxSimilarities := config.MaxSimilarities
	contextSizeLimit := config.ContextSizeLimit

	serverAgent, err := crewserver.NewAgent(
		ctx,
		crewserver.WithSingleAgent(coderAgent),
		crewserver.WithPort(httpPort),
		crewserver.WithRagAgentAndSimilarityConfig(ragAgent, similarityLimit, maxSimilarities),
		crewserver.WithCompressorAgentAndContextSize(compressorAgent, contextSizeLimit),
	)

	if err != nil {
		display.Errorf("❌ Error creating server agent: %v", err)
		return nil, err
	}
	display.Infof("🚀 SNIP server agent created and listening on port %s", serverAgent.GetPort())

	return serverAgent, nil
}

func chatWithAgent() {
	// Create logger from environment variable
	log := logger.GetLoggerFromEnv()
	config := loadConfiguration()

	ctx := context.Background()

	//httpPort := config.HTTPPort
	engineURL := config.EngineURL

	storePath := config.StorePath
	storeFile := config.StoreFile
	dataPath := config.DataPath

	ragAgent, metadataExtractorAgent, compressorAgent, err := initializeSideKickAgents(ctx, engineURL, dataPath, storePath, storeFile)
	if err != nil {
		log.Error("❌ Error creating sidekick agents: %v", err)
		return
	}

	coderAgent, err := initializeCoderChatAgent(ctx, engineURL)
	if err != nil {
		log.Error("❌ Error creating coder chat agent: %v", err)
		return
	}
	// Start watching the store file for changes in a goroutine if enabled
	enableStoreWatch := config.EnableStoreWatch
	if enableStoreWatch {
		storePathFile := filepath.Join(storePath, storeFile)
		storeWatchInterval := config.StoreWatchInterval
		go watchStoreFile(storePathFile, dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent, storeWatchInterval)
	}

	similarityLimit := config.SimilarityLimit
	maxSimilarities := config.MaxSimilarities
	contextSizeLimit := config.ContextSizeLimit

	crewAgent, err := crew.NewAgent(
		ctx,
		crew.WithSingleAgent(coderAgent),
		crew.WithRagAgentAndSimilarityConfig(ragAgent, similarityLimit, maxSimilarities),
		crew.WithCompressorAgentAndContextSize(compressorAgent, contextSizeLimit),
	)
	if err != nil {
		display.Errorf("❌ Error creating crew agent: %v", err)
		return
	}
	display.Infof("🤖 SNIP crew agent created for chat")

	for {

		markdownParser := display.NewMarkdownChunkParser()

		input := prompt.NewWithColor("🤖 Ask me something? [" + crewAgent.GetName() + "]")
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
			for i, msg := range crewAgent.GetMessages() {
				display.Infof("Message %d - Role: %s, Content: \n%s", i, msg.Role, msg.Content)
				display.Separator()
			}
			continue
		}

		if strings.HasPrefix(question, "/reset") {
			display.Infof("🔄 Resetting %s context", crewAgent.GetName())
			crewAgent.ResetMessages()
			continue
		}

		display.NewLine()

		result, err := crewAgent.StreamCompletion(question,
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
			display.Errorf("[%s][%v]failed to get completion: %v", crewAgent.GetName(), crewAgent.GetContextSize(), err)
			return
		}

		display.NewLine()
		display.Separator()
		display.KeyValue("Finish reason", result.FinishReason)
		display.KeyValue("Context size", fmt.Sprintf("%d characters", crewAgent.GetContextSize()))
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
