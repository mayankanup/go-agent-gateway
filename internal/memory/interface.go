package memory

import "github.com/mayankanup/go-agent-gateway/internal/models"

type Repository interface {
	GetConversation(
		conversationID string,
	) []models.Message

	AppendMessage(
		conversationID string,
		message models.Message,
	)
}
