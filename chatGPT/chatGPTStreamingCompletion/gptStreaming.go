package chatGPTStreamingCompletion

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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
		MaxTokens: 20,
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

		return response.Choices[0].Delta.Content
}