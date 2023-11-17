// services/openai_service.go
package services

import (
	"context"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func ChatWithOpenAI(prompt string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second) // Increased to 120 seconds
	defer cancel()

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
