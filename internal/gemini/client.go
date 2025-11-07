package gemini

import (
	"context"

	"google.golang.org/genai"
)

func GetGenaiClient() (*genai.Client, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GetApiKey(),
		Backend: genai.BackendGeminiAPI,
	})

	return client, err
}
