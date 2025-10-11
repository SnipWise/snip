package helpers

import "github.com/firebase/genkit/go/ai"

// CalculateTokenCount calculates the total number of tokens from a list of messages
// Uses a rough approximation: 1 token ≈ 4 characters
func CalculateTokenCount(messages []*ai.Message) int {
	totalTokens := 0
	for _, msg := range messages {
		// Count tokens in the message content
		for _, part := range msg.Content {
			if part.IsText() {
				// Rough approximation: 1 token ≈ 4 characters
				totalTokens += len(part.Text) / 4
			}
		}
	}
	return totalTokens
}
