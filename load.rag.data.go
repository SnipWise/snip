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
				if err != nil {
					display.Errorf("❌ Error extracting keywords from snippet %d: %v", idx, err)
					// Continue with embedding even if keyword extraction fails
				} else {
					display.Infof("🏷️  Keywords: %v", metadata.Keywords)
					display.Infof("📌 Topic: %s | Category: %s",
						metadata.MainTopic, metadata.Category)

					// Enrich the chunk with metadata
					enrichedSnippet := fmt.Sprintf("[METADATA]\nKeywords: %v\nTopic: %s\nCategory: %s\n\nContent:\n%s",
						metadata.Keywords, metadata.MainTopic, metadata.Category, snippet,
					)

					snippet = enrichedSnippet
				}

				// === SAVE EMBEDDING FOR SECTION ===
				err = ragAgent.SaveEmbedding(snippet)
				if err != nil {
					display.Errorf("❌ Error embedding snippet %d: %v", idx, err)
				} else {
					display.Infof("✅ Indexed snippet %d/%d", idx+1, len(snippets))

					//fmt.Println(snippet)
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
