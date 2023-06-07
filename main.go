package main

import (
	"chat/route"
	// "github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// Set up the router
	router := route.SetupRouter()
	// local host path :8080
	router.Run(":8080")

}
