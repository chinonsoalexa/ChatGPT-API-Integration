package chatGTP3StreamingCompletion
//  still under construct

import (
	"errors"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GTP3StreamingCompletion(prompt2 string) openai.CompletionResponse {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT4,
		MaxTokens: 100,
		Prompt:    prompt2,
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("CompletionStream error: %v\n", err)
		// return "an error occured in chatGTP3StreamingCompletion package and GTP3StreamingCompletion function"
	}
	defer stream.Close()

		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			// return	
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			// return 
		}


		fmt.Printf("Stream response: %v\n", response)
		return response 
}