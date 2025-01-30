package main

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, chatController *controllers.ChatController) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		v1.POST("/chat", chatController.HandleChat)
	}
}
