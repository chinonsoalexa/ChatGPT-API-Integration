package controllers

import (
	"net/http"

	GPT3 "chat/chatGPT/chatGPT3Completion"

	"github.com/gin-gonic/gin"
)

type Prompt struct {
    Text string `json:"text"`
    // Field2 int `json:"field2"`
}

func ChatGPT3CompletionContr(ctx *gin.Context) {
	    // Parse the request body into the struct
		var data Prompt
		if err := ctx.BindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	prompt := Prompt{Text: data.Text}
	result := GPT3.GPT3Completion(prompt.Text)

  // Send a response with the completion result
  ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})
}
