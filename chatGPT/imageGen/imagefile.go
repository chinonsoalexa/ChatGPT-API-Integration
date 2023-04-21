package imageGen

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"image/png"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func ImageRequest(imageDef string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	// Sample image by link
	reqUrl := openai.ImageRequest{
		Prompt:         imageDef,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
	}
	fmt.Println(respUrl.Data[0].URL)

	// Example image as base64
	reqBase64 := openai.ImageRequest{
		Prompt:         imageDef,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
	}

	file, err := os.Create("example.png")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
	}

	return "Your request to create " + imageDef  + " was created successfully."
}
