package voicetext

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func Voice() {
	c := openai.NewClient("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")
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "/Users/Nonso/downloads/Davido-Timeless-(ThinkNews/Davido-Feel_(thinkNews.com.ng).mp3",
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	fmt.Println(resp.Text)
}