package client

import (
	"context"
	"google.golang.org/genai"
)

func NewClient(ctx context.Context, apiKey string) *genai.Client {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		panic(err)
	}
	return client
}
