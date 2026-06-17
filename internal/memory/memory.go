package memory

import (
	"sync"

	"github.com/mayankanup/go-agent-gateway/internal/models"
)

type InMemoryRepository struct {
	mu            sync.RWMutex
	conversations map[string][]models.Message
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		conversations: make(
			map[string][]models.Message,
		),
	}
}

func (r *InMemoryRepository) GetConversation(
	conversationID string,
) []models.Message {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.conversations[conversationID]
}

func (r *InMemoryRepository) AppendMessage(
	conversationID string,
	message models.Message,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.conversations[conversationID] =
		append(
			r.conversations[conversationID],
			message,
		)
}
