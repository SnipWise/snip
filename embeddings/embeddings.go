package embeddings

import (
	"context"
	"snip/helpers"
	"snip/rag"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/openai/openai-go/option"
)

func Generate(ctx context.Context, llmURL string, embeddingsModel string) (ai.Embedder, rag.MemoryVectorStore, error) {

	store := rag.MemoryVectorStore{
		Records: make(map[string]rag.VectorRecord),
	}

	oaiPlugin := &openai.OpenAI{
		APIKey: "tada",
		Opts: []option.RequestOption{
			option.WithBaseURL(llmURL),
		},
	}
	g := genkit.Init(ctx, genkit.WithPlugins(oaiPlugin))

	embedder := oaiPlugin.DefineEmbedder(embeddingsModel, nil)

	snippetsFolder := helpers.GetEnvOrDefault("SNIPPETS_FOLDER", "./snippets")

	_, err := helpers.ForEachFile(snippetsFolder, ".md", func(path string) error {
		jsonFilePath := strings.TrimSuffix(path, ".md") + ".json"
		fmt.Println("🚧 Path:", path)

		// Test if JSON file already exists
		if _, err := os.Stat(jsonFilePath); err == nil {
			fmt.Println("⚠️  JSON file already exists, skipping:", jsonFilePath)

			// [IMPORTANT] read the JSON file and load the vector record in the vector store
			data, err := os.ReadFile(jsonFilePath)
			if err != nil {
				fmt.Println("😡 Error reading JSON file:", err)
				return err
			}
			var record rag.VectorRecord
			if err := json.Unmarshal(data, &record); err != nil {
				fmt.Println("😡 Error unmarshaling JSON file:", err)
				return err
			}
			// Add the record to the vector store
			if _, err := store.Save(record); err != nil {
				fmt.Println("😡 Error saving vector record to store:", err)
				return err
			}

			//fmt.Println("✅ Loaded record from JSON file:", record.Prompt, record.Id)
			fmt.Println("✅ Loaded record from JSON file:", record.Id)

			return nil // Skip processing this file
		} else if !os.IsNotExist(err) {
			// Some other error
			fmt.Println("😡 Error checking JSON file:", err)
			return err
		} else {
			fmt.Println("🆕 JSON file does not exist, processing:", jsonFilePath)

			// Reading file content
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("😡 Error reading file:", err)
				return err
			}
			content := string(data)
			resp, err := genkit.Embed(ctx, g,
				ai.WithEmbedder(embedder),
				ai.WithTextDocs(content),
			)

			if err != nil {
				fmt.Println("😡 Error generating embedding:", err)
				return err
			}
			fmt.Println(resp.Embeddings)

			for i, emb := range resp.Embeddings {
				//fmt.Printf("Chunk %d (%s) embedding: %v\n", i, chunks[i], emb)

				// Store the embedding in the vector store
				record, errSave := store.Save(rag.VectorRecord{
					Prompt:    content,
					Embedding: emb.Embedding,
				})
				if errSave != nil {
					fmt.Println("😡 Error saving vector record:", errSave)
					return errSave
				}
				fmt.Println("💾", i, "Saved record:", record.Prompt, record.Id)

				// Save the vector record to a JSON file
				vectorRecordJSON, err := json.MarshalIndent(record, "", "  ")
				if err != nil {
					fmt.Println("😡 Error marshaling vector record to JSON:", err)
					return err
				}

				// Write the JSON to a file
				err = os.WriteFile(jsonFilePath, vectorRecordJSON, 0644)
				if err != nil {
					fmt.Println("😡 Error writing vector record to file:", err)
					return err
				}
				return nil

			}

			return nil

		}

	})

	return embedder, store, err

}

// GenerateForFile generates embeddings for a single file and saves/loads the vector record
func GenerateForFile(ctx context.Context, filePath string, embedder ai.Embedder, store *rag.MemoryVectorStore, genkitInstance any) error {
	jsonFilePath := strings.TrimSuffix(filePath, ".md") + ".json"
	fmt.Println("🚧 Processing file:", filePath)

	// Test if JSON file already exists
	if _, err := os.Stat(jsonFilePath); err == nil {
		fmt.Println("⚠️  JSON file already exists, loading:", jsonFilePath)

		// Read the JSON file and load the vector record in the vector store
		data, err := os.ReadFile(jsonFilePath)
		if err != nil {
			fmt.Println("😡 Error reading JSON file:", err)
			return err
		}
		var record rag.VectorRecord
		if err := json.Unmarshal(data, &record); err != nil {
			fmt.Println("😡 Error unmarshaling JSON file:", err)
			return err
		}
		// Add the record to the vector store
		if _, err := store.Save(record); err != nil {
			fmt.Println("😡 Error saving vector record to store:", err)
			return err
		}

		fmt.Println("✅ Loaded record from JSON file:", record.Id)
		return nil
	} else if !os.IsNotExist(err) {
		// Some other error
		fmt.Println("😡 Error checking JSON file:", err)
		return err
	}

	// JSON file does not exist, generate embeddings
	fmt.Println("🆕 JSON file does not exist, generating embeddings for:", filePath)

	// Reading file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("😡 Error reading file:", err)
		return err
	}
	content := string(data)

	g, ok := genkitInstance.(*genkit.Genkit)
	if !ok {
		return fmt.Errorf("invalid genkit instance type")
	}

	resp, err := genkit.Embed(ctx, g,
		ai.WithEmbedder(embedder),
		ai.WithTextDocs(content),
	)
	if err != nil {
		fmt.Println("😡 Error generating embedding:", err)
		return err
	}

	if len(resp.Embeddings) == 0 {
		return fmt.Errorf("no embeddings generated for file: %s", filePath)
	}

	// Store the embedding in the vector store (use first embedding)
	record, errSave := store.Save(rag.VectorRecord{
		Prompt:    content,
		Embedding: resp.Embeddings[0].Embedding,
	})
	if errSave != nil {
		fmt.Println("😡 Error saving vector record:", errSave)
		return errSave
	}
	fmt.Println("💾 Saved record:", record.Id)

	// Save the vector record to a JSON file
	vectorRecordJSON, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		fmt.Println("😡 Error marshaling vector record to JSON:", err)
		return err
	}

	// Write the JSON to a file
	err = os.WriteFile(jsonFilePath, vectorRecordJSON, 0644)
	if err != nil {
		fmt.Println("😡 Error writing vector record to file:", err)
		return err
	}

	fmt.Println("✅ Embedding generated and saved for:", filePath)
	return nil
}

