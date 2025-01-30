package routes

import (
	"fmt"
	"go-api/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, chatController *controllers.ChatController) {
	// Add middleware
	router.Use(corsMiddleware())
	router.Use(requestLogger())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		v1.POST("/chat", chatController.HandleChat)
		v1.DELETE("/chat/:userID", chatController.ClearChat)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		duration := time.Since(start)
		logMessage := fmt.Sprintf("[%s] %s %s %d %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
		gin.DefaultWriter.Write([]byte(logMessage))
	}
}
