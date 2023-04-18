package controllers
//  still under construct
import (
	"net/http"

	GPT3s "chat/chatGPT/chatGPT3StreamingCompletion"

	"github.com/gin-gonic/gin"
)

type Prompt2 struct {
    Text string `json:"text"`
    // Field2 int `json:"field2"`
}

func GTP3StreamingCompletionContr(ctx *gin.Context) {
	    // Parse the request body into the struct
		var data Prompt2
		if err := ctx.BindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	prompt := Prompt2{Text: data.Text}
	result := GPT3s.GTP3StreamingCompletion(prompt.Text)

  // Send a response with the completion result
  ctx.JSON(http.StatusOK, gin.H{"message": "Success", "result": result.Choices[0].Text})
}
