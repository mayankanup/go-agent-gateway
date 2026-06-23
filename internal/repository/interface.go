package repository

import "github.com/mayankanup/go-agent-gateway/internal/models"

type ConversationRepository interface {
	AppendMessage(
		conversationID string,
		message models.Message,
	) error

	GetConversation(
		conversationID string,
	) ([]models.Message, error)
}
