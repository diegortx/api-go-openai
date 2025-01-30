package services

import (
	"go-api/models"
	"sync"
	"time"
)

type ConversationService interface {
	GetConversation(userID string) []models.Message
	AddMessage(userID string, message models.Message) error
	ClearConversation(userID string) error
}

type ConversationServiceImpl struct {
	conversations map[string]models.Conversation
	mutex         sync.RWMutex
	maxMessages   int
}

func NewConversationService() ConversationService {
	return &ConversationServiceImpl{
		conversations: make(map[string]models.Conversation),
		maxMessages:   10, // Maximum messages to keep in context
	}
}

func (s *ConversationServiceImpl) GetConversation(userID string) []models.Message {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if conv, exists := s.conversations[userID]; exists {
		return conv.Messages
	}
	return []models.Message{}
}

func (s *ConversationServiceImpl) AddMessage(userID string, message models.Message) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	conv, exists := s.conversations[userID]
	if !exists {
		conv = models.Conversation{
			UserID:   userID,
			Messages: []models.Message{},
		}
	}

	// Add new message
	conv.Messages = append(conv.Messages, message)

	// Keep only the last N messages
	if len(conv.Messages) > s.maxMessages {
		conv.Messages = conv.Messages[len(conv.Messages)-s.maxMessages:]
	}

	conv.LastUpdated = time.Now().Unix()
	s.conversations[userID] = conv

	return nil
}

func (s *ConversationServiceImpl) ClearConversation(userID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.conversations, userID)
	return nil
}
