package routes

import (

	GPT "chat/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.POST("/gpt_streaming", GPT.ChatGPT3CompletionContr)
    router.POST("/gpt_streaming", GPT.GTP3StreamingCompletionContr)
    router.POST("/gpt_streaming", GPT.GPTstreamingContr)
    router.POST("/gpt_streaming", GPT.ImageGenContr)
    router.POST("/gpt_streaming", GPT.SpeechToTextContr)
    router.POST("/gpt_streaming", GPT.ToTextContr)

    return router
}