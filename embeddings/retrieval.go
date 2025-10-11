package embeddings

import (
	"context"
	"dmr-genkit-stream-completion/helpers"
	"dmr-genkit-stream-completion/rag"
	"fmt"

	"github.com/firebase/genkit/go/ai"
)

// RetrieveSimilarDocuments performs similarity search and returns relevant context
func RetrieveSimilarDocuments(ctx context.Context, query string, retriever ai.Retriever) (string, error) {
	// Create a query document from the user question
	queryDoc := ai.DocumentFromText(query, nil)

	similarityThreshold := helpers.StringToFloat(helpers.GetEnvOrDefault("SIMILARITY_THRESHOLD", "0.5"))
	similarityMaxResults := helpers.StringToInt(helpers.GetEnvOrDefault("SIMILARITY_MAX_RESULTS", "3"))

	// Create a retriever request with custom options
	request := &ai.RetrieverRequest{
		Query: queryDoc,
		Options: rag.MemoryVectorRetrieverOptions{
			Limit:      similarityThreshold,  // Lower similarity threshold to get more results
			MaxResults: similarityMaxResults, // Return top N results
		},
	}

	// Use the memory vector retriever to find similar documents
	retrieveResponse, err := retriever.Retrieve(ctx, request)
	if err != nil {
		return "", err
	}

	similarDocuments := ""

	fmt.Printf("\nFound %d similar documents:\n", len(retrieveResponse.Documents))
	for i, doc := range retrieveResponse.Documents {
		similarity := doc.Metadata["cosine_similarity"]
		id := doc.Metadata["id"]
		fmt.Printf("%d. ID: %s, Similarity: %.4f\n", i+1, id, similarity)
		fmt.Printf("   Content: %s\n\n", doc.Content[0].Text)
		similarDocuments += doc.Content[0].Text
	}

	return similarDocuments, nil
}
