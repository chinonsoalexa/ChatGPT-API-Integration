package controllers

import (
	"net/http"

	GPT3x "chat/chatGPT/textToText"

	"github.com/gin-gonic/gin"
)

type InputText struct {
	Text string `json:"text"`
	// Field2 int `json:"field2"`
}

func ToTextContr(ctx *gin.Context) {
	// Parse the request body into the struct
	var data InputText
	if err := ctx.BindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prompt := InputText{Text: data.Text}
	result := GPT3x.ToText(prompt.Text)

	// Send a response with the completion result
	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result})
}
