package route

import (
	"net/http"

	GPT "chat/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Configure CORS
	// Use the cors middleware to set CORS headers and allow all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.POST("/gpt_image_gen", GPT.ImageGenContr)
	router.POST("/gpt_to_text", GPT.ToTextContr)

	router.OPTIONS("/gpt_image_gen", HandleOptions)
	router.OPTIONS("/gpt_to_text", HandleOptions)

	return router
}

func HandleOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")

	// Respond with 200 OK
	c.Status(http.StatusOK)
}
