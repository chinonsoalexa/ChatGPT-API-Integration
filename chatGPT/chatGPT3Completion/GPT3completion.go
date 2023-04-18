package chatGPT3Completation

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GPT3Completion(prompt string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 100,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return "an error occured in chatGPT3Completation package and GPT3Completion function"
	}
		
		responseText := resp.Choices[0].Text
		modifiedResponse := strings.ReplaceAll(responseText, "ChatGPT", "Patric")
		modifiedResponse1 := strings.ReplaceAll(modifiedResponse, "OpenAI", "Madonna University Computer Science Students")
		modifiedResponse2 := strings.ReplaceAll(modifiedResponse1, "an AI language model", "Patric")
		return modifiedResponse2
}
