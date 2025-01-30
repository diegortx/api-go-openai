package controllers

import (
	"fmt"
	"go-api/models"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ChatController struct {
	openAIService       services.OpenAIService
	conversationService services.ConversationService
}

func NewChatController(openAIService services.OpenAIService, conversationService services.ConversationService) *ChatController {
	return &ChatController{
		openAIService:       openAIService,
		conversationService: conversationService,
	}
}

func (c *ChatController) HandleChat(ctx *gin.Context) {
	var request models.ChatRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		var errorMessages string
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "UserID":
					errorMessages = "User ID is required"
				case "Message":
					switch e.Tag() {
					case "required":
						errorMessages = "Message is required"
					case "min":
						errorMessages = "Message is too short"
					case "max":
						errorMessages = "Message is too long (max 1000 characters)"
					}
				}
			}
		} else {
			errorMessages = "Invalid request format"
		}

		ctx.JSON(http.StatusBadRequest, models.ChatResponse{
			Success: false,
			Error:   errorMessages,
		})
		return
	}

	// Log incoming request with context information
	logMessage := fmt.Sprintf(
		"Incoming chat request from user %s: %s, System Prompt: %s\n",
		request.UserID,
		request.Message,
		request.Context.SystemPrompt,
	)
	gin.DefaultWriter.Write([]byte(logMessage))

	response, err := c.openAIService.GenerateResponse(
		request.UserID,
		request.Message,
		request,
		c.conversationService,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ChatResponse{
			Success: false,
			Error:   "Failed to generate response: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ChatResponse{
		Success:  true,
		Response: response,
	})
}

func (c *ChatController) ClearChat(ctx *gin.Context) {
	userID := ctx.Param("userID")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, models.ChatResponse{
			Success: false,
			Error:   "User ID is required",
		})
		return
	}

	if err := c.conversationService.ClearConversation(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ChatResponse{
			Success: false,
			Error:   "Failed to clear conversation: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ChatResponse{
		Success:  true,
		Response: "Conversation cleared successfully",
	})
}
