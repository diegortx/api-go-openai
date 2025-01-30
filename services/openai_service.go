package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/config"
	"go-api/models"
	"net/http"
)

type OpenAIService interface {
	GenerateResponse(userID string, message string, context models.ChatRequest, conversationService ConversationService) (string, error)
}

type OpenAIServiceImpl struct {
	config *config.Config
	client *http.Client
}

func NewOpenAIService(config *config.Config) OpenAIService {
	return &OpenAIServiceImpl{
		config: config,
		client: &http.Client{},
	}
}

func (s *OpenAIServiceImpl) GenerateResponse(userID string, message string, context models.ChatRequest, conversationService ConversationService) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// Get previous conversation context
	previousMessages := conversationService.GetConversation(userID)

	// Add system message if provided
	var allMessages []models.Message
	if context.Context.SystemPrompt != "" {
		allMessages = append(allMessages, models.Message{
			Role:    "system",
			Content: context.Context.SystemPrompt,
		})
	}

	// Add previous messages
	allMessages = append(allMessages, previousMessages...)

	// Add new user message
	userMessage := models.Message{
		Role:    "user",
		Content: message,
	}
	allMessages = append(allMessages, userMessage)

	// Set temperature from context or use default
	temperature := float32(0.5)
	if context.Context.Temperature != nil {
		temperature = *context.Context.Temperature
	}

	reqBody := models.OpenAIRequest{
		Model:       "gpt-4o-mini",
		Messages:    allMessages,
		Temperature: temperature,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.config.OpenAIAPIKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var openAIResp models.OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	// Store both user message and assistant response in conversation
	conversationService.AddMessage(userID, userMessage)
	conversationService.AddMessage(userID, openAIResp.Choices[0].Message)

	return openAIResp.Choices[0].Message.Content, nil
}
