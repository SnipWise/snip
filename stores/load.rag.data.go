package main

import (
	"fmt"

	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/rag/chunks"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/messages"
	"github.com/snipwise/nova/nova-sdk/messages/roles"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func LoadSnippetData(dataPath, storePath, storeFile string, ragAgent *rag.Agent, metadataExtractorAgent *structured.Agent[KeywordMetadata]) error {
	storePathFile := storePath + "/" + storeFile
	// === LOAD OR CREATE STORE ===
	if ragAgent.StoreFileExists(storePathFile) {
		// Load existing store
		err := ragAgent.LoadStore(storePathFile)
		if err != nil {
			display.Errorf("❌ Error loading store %s: %v", storePathFile, err)
			return err
		}
		display.Infof("✅ RAG store loaded from %s", storePathFile)
	} else {
		embeddingDimension := ragAgent.GetEmbeddingDimension()

		display.Infof("📝 Store not found. Creating new store and indexing character sheet...")

		// Read character sheet content
		snippetsDocuments, err := files.GetContentFiles(dataPath, ".md")

		// BEGIN: Indexing multiple snippet documents
		for _, document := range snippetsDocuments {
			// === CHUNK AND INDEX CONTENT ===
			//snippets := chunks.SplitTextWithDelimiter(document, "----------")
			snippets := chunks.SplitMarkdownBySections(document)
			display.Infof("📄 Split snippets document into %d sections", len(snippetsDocuments))

			// BEGIN: Index each section (each snippet)
			for idx, snippet := range snippets {
				// === EXTRACT METADATA FOR snippet ===
				// Extract keywords and metadata using structured agent
				extractionPrompt := fmt.Sprintf(`Analyze the following content and extract relevant metadata.
					Content:
					%s

					Extract:
					- Keywords: only 4 keywords, important terms and concepts from the markdown snippet title then from the content
					- Main topic: the primary subject (use the markdown snippet title)
					- Category: type of content
					`,
					snippet,
				)

				metadata, _, err := metadataExtractorAgent.GenerateStructuredData([]messages.Message{
					{Role: roles.User, Content: extractionPrompt},
				})
				enrichedSnippet := ""
				metaDataContent := ""
				if err != nil {
					display.Errorf("❌ Error extracting keywords from snippet %d: %v", idx, err)
					// Continue with embedding even if keyword extraction fails
					enrichedSnippet = snippet
				} else {
					display.Infof("🏷️  Keywords: %v", metadata.Keywords)
					display.Infof("📌 Topic: %s | Category: %s",
						metadata.MainTopic, metadata.Category)

					metaDataContent = fmt.Sprintf("[METADATA]\nKeywords: %v\nTopic: %s\nCategory: %s",
						metadata.Keywords, metadata.MainTopic, metadata.Category,
					)

					// Enrich the chunk with metadata
					enrichedSnippet = fmt.Sprintf("%s\n\nContent:\n%s",
						metaDataContent, snippet,
					)

					//snippet = enrichedSnippet
				}

				if len(enrichedSnippet) > embeddingDimension {
					display.Infof("⚠️  Snippet %d length (%d) exceeds embedding dimension (%d). Consider chunking further.",
						idx, len(enrichedSnippet), embeddingDimension)

					// NBCHUNK = LENSNIPP / (SIZEONECHUNK + 128 - LENMETA)
					// SIZEONECHUNK = LENSNIPP / NBCHUNK + LENMETA - 128

					sizePerChunk := embeddingDimension - len(metaDataContent) - 128

					newChunks := chunks.ChunkText(snippet, sizePerChunk, 128)

					// Index each new chunk
					for cidx, chunk := range newChunks {
						chunkWithMeta := fmt.Sprintf("%s\n\nContent Chunk:\n%s",
							metaDataContent, chunk,
						)

						err = ragAgent.SaveEmbedding(chunkWithMeta)
						if err != nil {
							display.Errorf("❌ Error embedding snippet %d chunk %d: %v", idx, cidx, err)
						} else {
							display.Infof("✅ Indexed snippet %d chunk %d/%d", idx, cidx+1, len(newChunks))
						}
					}

				} else {
					// Proceed with original enriched snippet
					// === SAVE EMBEDDING FOR SECTION ===
					err = ragAgent.SaveEmbedding(enrichedSnippet)
					if err != nil {
						display.Errorf("❌ Error embedding snippet %d: %v", idx, err)
					} else {
						display.Infof("✅ Indexed snippet %d/%d", idx+1, len(snippets))

						//fmt.Println(snippet)
					}
				}

			} // END: Index each section (each snippet)

		} // END: Indexing multiple snippet documents

		// === PERSIST STORE TO DISK ===
		err = ragAgent.PersistStore(storePathFile)
		if err != nil {
			display.Errorf("❌ Error persisting store: %v", err)
			return err
		}
		display.Infof("💾 RAG store saved to %s", storePathFile)

	}
	return nil
}
