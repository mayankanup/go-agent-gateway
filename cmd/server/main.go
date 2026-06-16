package main

import (
	"log"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/api"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
)

func main() {

	mockProvider := provider.NewMockProvider()

	chatHandler := api.NewChatHandler(
		mockProvider,
	)

	mux := http.NewServeMux()

	mux.Handle(
		"/chat",
		chatHandler,
	)

	log.Println(
		"Agent Gateway running on :8080",
	)

	err := http.ListenAndServe(
		":8080",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}
