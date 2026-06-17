package provider

import (
	"context"

	"github.com/mayankanup/go-agent-gateway/internal/models"
)

type LLMProvider interface {
	Chat(
		ctx context.Context,
		messages []models.Message,
	) (string, error)
}
