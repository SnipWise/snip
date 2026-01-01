package main

import (
	"context"

	"github.com/joho/godotenv"

	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/logger"
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

	storePath := env.GetEnvOrDefault("STORE_PATH", "./")
	storeFile := env.GetEnvOrDefault("STORE_FILE", "python.snippets.store.json")
	dataPath := env.GetEnvOrDefault("DATA_PATH", "./python")

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

}
