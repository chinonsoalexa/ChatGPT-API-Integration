package controllers

import (
	"net/http"

	GPTs "chat/chatGPT/chatGPTStreamingCompletion"

	"github.com/gin-gonic/gin"
)

type Prompt3 struct {
    Text string `json:"text"`
    // Field2 int `json:"field2"`
}

func GPTstreamingContr(ctx *gin.Context) {
	    // Parse the request body into the struct
		var data Prompt3
		if err := ctx.BindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	prompt := Prompt3{Text: data.Text}
	result := GPTs.GPTstreaming(prompt.Text)

  // Send a response with the completion result
  ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})
}
