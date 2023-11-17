package services

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// ChatWithOpenAI sends a prompt to ChatGPT and returns the response
func ChatWithOpenAI(prompt string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // Choose the model you want to use
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
