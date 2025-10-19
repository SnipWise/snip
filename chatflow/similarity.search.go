package chatflow

import (
	"context"
	"snip/embeddings"

	"github.com/firebase/genkit/go/ai"
)

// performSimilaritySearch retrieves similar documents and updates the message history
func performSimilaritySearch(
	ctx context.Context,
	userMessage string,
	config StreamingChatFlowConfig,
) error {
	// Retrieve relevant context from the vector store
	similarDocuments, details, err := embeddings.RetrieveSimilarDocuments(ctx, userMessage, config.MemoryRetriever)
	if err != nil {
		return err
	}

	// Update global similarities data via callback
	if config.UpdateSimilarities != nil {
		config.UpdateSimilarities(userMessage, details)
	}

	if similarDocuments != "" {
		// Add Similarities to Messages
		*config.Messages = append(*config.Messages, ai.NewSystemTextMessage("Relevant context:\n"+similarDocuments))
	}

	return nil
}
