package main

import (
	"log"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/agent"
	"github.com/mayankanup/go-agent-gateway/internal/api"
	"github.com/mayankanup/go-agent-gateway/internal/memory"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
	"github.com/mayankanup/go-agent-gateway/internal/tools"
)

func main() {

	registry :=
		tools.NewRegistry()

	registry.Register(
		tools.NewDateTool(),
	)

	registry.Register(
		tools.NewCalculatorTool(),
	)

	agent :=
		agent.NewAgent(
			registry,
		)

	mockProvider :=
		provider.NewMockProvider(
			agent,
		)

	repo :=
		memory.NewInMemoryRepository()

	chatHandler :=
		api.NewChatHandler(
			mockProvider,
			repo,
		)

	mux := http.NewServeMux()

	mux.Handle(
		"/chat",
		chatHandler,
	)

	log.Println(
		"Agent Gateway running on :8080",
	)

	log.Fatal(
		http.ListenAndServe(
			":8080",
			mux,
		),
	)
}
