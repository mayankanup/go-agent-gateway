package provider

import (
	"context"

	"github.com/mayankanup/go-agent-gateway/internal/agent"
	"github.com/mayankanup/go-agent-gateway/internal/models"
)

type MockProvider struct {
	agent *agent.Agent
}

func NewMockProvider(
	agent *agent.Agent,
) *MockProvider {

	return &MockProvider{
		agent: agent,
	}
}

func (m *MockProvider) Chat(
	ctx context.Context,
	messages []models.Message,
) (string, error) {

	return m.agent.Process(
		ctx,
		messages,
	)
}
