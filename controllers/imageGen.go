package controllers

import (
	"net/http"

	GPT3i "chat/chatGPT/imageGen"

	"github.com/gin-gonic/gin"
)

type ImageDef struct {
    Text string `json:"text"`
    // Field2 int `json:"field2"`
}

func ImageGenContr(ctx *gin.Context) {
	    // Parse the request body into the struct
		var data ImageDef
		if err := ctx.BindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	prompt := ImageDef{Text: data.Text}
	result := GPT3i.ImageRequest(prompt.Text)

  // Send a response with the completion result
  ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})
}
