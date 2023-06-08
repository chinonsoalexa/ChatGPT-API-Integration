package imageGen

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// ImageRequest takes in a string as an argument and returns a string as a response.
// It creates an image using OpenAI's GPT-3 API, saves the image to a file,
// and returns a success message.
func ImageRequest(imageDef string) string {
	API_KEY := "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"

	// Create a new client for the OpenAI API using the API key
	client := openai.NewClient(API_KEY)

	// Create a context for the API requests
	ctx := context.Background()

	// Create an image by link using OpenAI's GPT-3 API
	reqUrl := openai.ImageRequest{
		Prompt:         imageDef,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	// Make a request to the OpenAI API to create the image
	respUrl, err := client.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
	}

	// Print the URL of the created image
	fmt.Println(respUrl.Data[0].URL)

	// Create an example image as base64 using OpenAI's GPT-3 API
	reqBase64 := openai.ImageRequest{
		Prompt:         imageDef,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	// Make a request to the OpenAI API to create the image
	respBase64, err := client.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
	}

	// Decode the base64 image data into bytes
	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
	}

	// Create a new reader from the image bytes
	r := bytes.NewReader(imgBytes)

	// Decode the PNG image data from the reader
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
	}

	// Create a new file to save the image to
	file, err := os.Create("example.png")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
	}

	// Close the file when the function returns
	defer file.Close()

	// Encode the PNG image data and save it to the file
	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
	}

	// Return a success message
	return "Your request to create " + imageDef + " was created successfully."
}
