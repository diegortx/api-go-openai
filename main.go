package main

import (
	"go-api/config"
	"go-api/controllers"
	"go-api/routes"
	"go-api/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize router
	router := gin.Default()

	// Initialize services and controllers
	openAIService := services.NewOpenAIService(cfg)
	conversationService := services.NewConversationService()
	chatController := controllers.NewChatController(openAIService, conversationService)

	// Initialize routes
	routes.InitializeRoutes(router, chatController)

	// Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
