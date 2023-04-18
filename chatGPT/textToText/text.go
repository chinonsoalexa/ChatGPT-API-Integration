package textToText

import (
	"log"
	"os"
	"context"
	"fmt"
	"strings"

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

	responseText :=  resp.Choices[0].Message.Content
	modifiedResponse := strings.ReplaceAll(responseText, "ChatGPT", "Patric")
	modifiedResponse1 := strings.ReplaceAll(modifiedResponse, "OpenAI", "Madonna University Computer Science Students")
	modifiedResponse2 := strings.ReplaceAll(modifiedResponse1, "an AI language model", "Patric")

	return modifiedResponse2
}
