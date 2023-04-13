package controllers

import (
	"net/http"

	GPT3t "chat/chatGPT/speechToText"

	"github.com/gin-gonic/gin"
)

type Speechfilepath struct {
    Text string `json:"text"`
    // Field2 int `json:"field2"`
}

func SpeechToTextContr(ctx *gin.Context) {
	    // Parse the request body into the struct
		var data Speechfilepath
		if err := ctx.BindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	prompt := Speechfilepath{Text: data.Text}
	result := GPT3t.SpeechToText(prompt.Text)

  // Send a response with the completion result
  ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})
}
