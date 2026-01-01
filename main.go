package main

import (
	"context"

	"github.com/joho/godotenv"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/server"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/toolbox/logger"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func main() {

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

	httpPort := env.GetEnvOrDefault("HTTP_PORT", "8080")

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:12434/engines/llama.cpp/v1")
	ragModelId := env.GetEnvOrDefault("RAG_EMBEDDING_MODEL_ID", "ai/embeddinggemma:latest")
	metadataModelId := env.GetEnvOrDefault("METADATA_MODEL_ID", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	ragAgent, err := CreateRagAgent(ctx, engineURL, ragModelId)
	if err != nil {
		return
	}

	metadataExtractorAgent, err := CreateMetadataExtractorAgent(ctx, engineURL, metadataModelId)
	if err != nil {
		return
	}

	err = LoadSnippetData(dataPath, storePath, storeFile, ragAgent, metadataExtractorAgent)
	if err != nil {
		return
	}

	compressorAgent, err := CreateCompressorAgent(ctx, engineURL)
	if err != nil {
		return
	}

	// ------------------------------------------------
	// Create the server agent
	// ------------------------------------------------
	temperature := env.GetEnvFloatOrDefault("CODER_TEMPERATURE", 0.8)
	topP := env.GetEnvFloatOrDefault("CODER_TOP_P", 0.9)
	coderAgentModelID := env.GetEnvOrDefault("CODER_MODEL_ID", "hf.co/quantfactory/deepseek-coder-7b-instruct-v1.5-gguf:q4_k_m")
	directivesPathFile := env.GetEnvOrDefault("CODER_DIRECTIVES_PATH_FILE", "./directives/coder.agent.system.instructions.md")
	coderAgentSystemInstructionsContent, err := files.ReadTextFile(directivesPathFile)
	if err != nil {
		return
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
		return
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

	// Start the server agent
	if err := serverAgent.StartServer(); err != nil {
		display.Errorf("❌ Error starting crew server agent: %v", err)
		return
	}
}
