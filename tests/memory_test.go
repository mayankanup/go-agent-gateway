package tests

import (
	"testing"

	"github.com/mayankanup/go-agent-gateway/internal/memory"
	"github.com/mayankanup/go-agent-gateway/internal/models"
)

func TestConversationMemory(
	t *testing.T,
) {

	repo :=
		memory.NewInMemoryRepository()

	repo.AppendMessage(
		"conv1",
		models.Message{
			Role:    "user",
			Content: "hello",
		},
	)

	messages :=
		repo.GetConversation(
			"conv1",
		)

	if len(messages) != 1 {
		t.Fatalf(
			"expected 1 message got %d",
			len(messages),
		)
	}
}
