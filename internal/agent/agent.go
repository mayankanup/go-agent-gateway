package agent

import (
	"context"
	"strings"

	"github.com/mayankanup/go-agent-gateway/internal/models"
	"github.com/mayankanup/go-agent-gateway/internal/tools"
)

type Agent struct {
	registry *tools.Registry
}

func NewAgent(
	registry *tools.Registry,
) *Agent {

	return &Agent{
		registry: registry,
	}
}

func (a *Agent) Process(
	ctx context.Context,
	messages []models.Message,
) (string, error) {

	if len(messages) == 0 {
		return "No messages.", nil
	}

	last :=
		strings.ToLower(
			messages[len(messages)-1].Content,
		)

	if strings.Contains(
		last,
		"date",
	) {

		tool, _ :=
			a.registry.Get("date")

		return tool.Execute(
			ctx,
			last,
		)
	}

	if strings.Contains(
		last,
		"calculate",
	) {

		tool, _ :=
			a.registry.Get(
				"calculator",
			)

		return tool.Execute(
			ctx,
			last,
		)
	}

	return "Mock agent response.", nil
}
