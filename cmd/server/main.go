package main

import (
	"log"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/api"
	"github.com/mayankanup/go-agent-gateway/internal/memory"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
)

func main() {

	mockProvider :=
		provider.NewMockProvider()

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
