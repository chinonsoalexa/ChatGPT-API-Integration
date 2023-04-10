package configProxy

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func GPTConfig() {
	// Initialize the OpenAI client
	config := openai.DefaultConfig("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")
	client := openai.NewClientWithConfig(config)

	// Initialize the Gin router
	router := gin.Default()

	// Define a route to handle completion requests
	router.POST("/complete", func(c *gin.Context) {
		// Parse the request JSON body
		var req struct {
			Prompt    string `json:"prompt"`
			// MaxLength int    `json:"max_length"`
			Model     string `json:"model"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Set up the completion request
		completionReq := &openai.CompletionRequest{
			Prompt:    req.Prompt,
			// MaxLength: req.MaxLength,
			Model:     req.Model,
		}

		// Send the completion request
		res, err := client.CreateCompletion(context.Background(), *completionReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the response
		c.JSON(http.StatusOK, gin.H{"text": res.Choices[0].Text})
	})

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
