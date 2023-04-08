package main

import (
	// "context"
	// "fmt"
	// openai "github.com/sashabaranov/go-openai"
 	voice "chat/voicetext"
)

func main() {
	// for i :=1; true; i++ {
	// 	fmt.Println("Please enter your question here: ")
	// 	var first string
  
    // // Taking input from user
    // fmt.Scanln(&first)
	// 	client := openai.NewClient("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")
	// resp, err := client.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: first,
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Printf("ChatCompletion error: %v\n", err)
	// 	return
	// }

	// fmt.Println(resp.Choices[0].Message.Content)
	// }
	voice.Voice()
}
