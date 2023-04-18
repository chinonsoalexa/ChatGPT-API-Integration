package chatGPTStreamingCompletion

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GPTstreaming(Prompt3 string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{

			{
				Role:    openai.ChatMessageRoleUser,
				Content: Prompt3,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return "an error occured in chatGPTStreamingCompletion package and GPTstreaming function"
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")

	response, err := stream.Recv()
	if errors.Is(err, io.EOF) {
		fmt.Println("\nStream finished")
		return "an error occured in chatGPTStreamingCompletion package and GPTstreaming function"
	}

	if err != nil {
		fmt.Printf("\nStream error: %v\n", err)
		return "an error occured in chatGPTStreamingCompletion package and GPTstreaming function"
	}
	responseText :=  response.Choices[0].Delta.Content
	modifiedResponse := strings.ReplaceAll(responseText, "ChatGPT", "Patric")
	modifiedResponse1 := strings.ReplaceAll(modifiedResponse, "OpenAI", "Madonna University Computer Science Students")
	modifiedResponse2 := strings.ReplaceAll(modifiedResponse1, "an AI language model", "Patric")
	return modifiedResponse2
}
