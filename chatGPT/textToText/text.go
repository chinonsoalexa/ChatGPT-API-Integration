package textToText

import (
	"log"
	"os"
	"context"
	"fmt"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func ToText(inputText string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: inputText,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "an error occured in textToText package and ToText function"
	}

	return resp.Choices[0].Message.Content
}
