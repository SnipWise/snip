package embeddings

import (
	"context"
	"snip/helpers"
	"snip/rag"
	"fmt"

	"github.com/firebase/genkit/go/ai"
)

// SimilarityDetail represents detailed information about a similarity search result
type SimilarityDetail struct {
	ID         string
	Similarity float64
	Content    string
}

// RetrieveSimilarDocuments performs similarity search and returns relevant context with details
func RetrieveSimilarDocuments(ctx context.Context, query string, retriever ai.Retriever) (string, []SimilarityDetail, error) {
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
		return "", nil, err
	}

	similarDocuments := ""
	details := make([]SimilarityDetail, 0, len(retrieveResponse.Documents))

	fmt.Printf("\n📘 Found %d similar documents:\n", len(retrieveResponse.Documents))
	for i, doc := range retrieveResponse.Documents {
		similarity := doc.Metadata["cosine_similarity"].(float64)
		id := doc.Metadata["id"].(string)
		content := doc.Content[0].Text

		fmt.Printf("%d. ID: %s, Similarity: %.4f\n", i+1, id, similarity)
		fmt.Printf("   Content: %s\n\n", content)

		similarDocuments += content
		details = append(details, SimilarityDetail{
			ID:         id,
			Similarity: similarity,
			Content:    content,
		})
	}

	return similarDocuments, details, nil
}
