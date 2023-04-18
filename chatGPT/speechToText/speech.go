package voicetext

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func SpeechToText(speechfilepath string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: speechfilepath,
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
	}

	responseText :=  resp.Text
	modifiedResponse := strings.ReplaceAll(responseText, "ChatGPT", "Patric")
	modifiedResponse1 := strings.ReplaceAll(modifiedResponse, "OpenAI", "Madonna University Computer Science Students")
	modifiedResponse2 := strings.ReplaceAll(modifiedResponse1, "an AI language model", "Patric")
	return modifiedResponse2
}
