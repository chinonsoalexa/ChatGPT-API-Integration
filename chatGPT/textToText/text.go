package textToText

import (
    "context"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/joho/godotenv"
    openai "github.com/sashabaranov/go-openai"
)

var chatHistory []openai.ChatCompletionMessage

func ToText(inputText string) string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    apiKey := os.Getenv("API_KEY")
    client := openai.NewClient(apiKey)
    chatHistory = append(chatHistory, openai.ChatCompletionMessage{
        Role:    openai.ChatMessageRoleUser,
        Content: inputText,
    })
    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model:    openai.GPT3Dot5Turbo,
            Messages: chatHistory,
        },
    )

    if err != nil {
        fmt.Printf("ChatCompletion error: %v\n", err)
        return "Back-End server not connected to the internet. Please make sure that you're connected to the internet."
    }

    responseText := resp.Choices[0].Message.Content
    modifiedResponse := strings.ReplaceAll(responseText, "ChatGPT", "Patric")
    modifiedResponse1 := strings.ReplaceAll(modifiedResponse, "OpenAI", "Madonna University Computer Science Students")
    modifiedResponse2 := strings.ReplaceAll(modifiedResponse1, "an AI language model", "Patric")

    chatHistory = append(chatHistory, resp.Choices[0].Message)

    return modifiedResponse2
}
