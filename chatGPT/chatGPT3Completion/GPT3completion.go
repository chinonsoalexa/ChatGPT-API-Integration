package chatGPT3Completation

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func GPT3Completion() {
	c := openai.NewClient("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")
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}