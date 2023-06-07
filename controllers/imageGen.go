package controllers

import (
	"net/http"

	GPT3i "chat/chatGPT/imageGen"
	"github.com/gin-gonic/gin"
)

// ImageDef defines the structure for the JSON request body
type ImageDef struct {
    Text string `json:"text"` // Text field in the JSON request body
}

// ImageGenContr is a Gin handler function that handles the image generation request
func ImageGenContr(ctx *gin.Context) {
	// Parse the request body into the struct
	var data ImageDef
	if err := ctx.BindJSON(&data); err != nil { // Bind the JSON data to the ImageDef struct
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // If there's an error, return a 400 Bad Request status code with the error message
		return
	}

	prompt := ImageDef{Text: data.Text} // Extract the text field from the ImageDef struct
	result := GPT3i.ImageRequest(prompt.Text) // Call the ImageRequest function from the imageGen package to generate the image

	// Send a response with the completion result
	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})
}
