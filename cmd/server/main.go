package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mayankanup/go-agent-gateway/internal/agent"
	"github.com/mayankanup/go-agent-gateway/internal/api"
	"github.com/mayankanup/go-agent-gateway/internal/database"
	"github.com/mayankanup/go-agent-gateway/internal/metrics"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
	"github.com/mayankanup/go-agent-gateway/internal/repository"
	"github.com/mayankanup/go-agent-gateway/internal/tools"
)

func main() {

	db, err := sql.Open(
		"sqlite3",
		"gateway.db",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = database.Migrate(db)

	if err != nil {
		log.Fatal(err)
	}

	repo :=
		repository.NewSQLiteRepository(
			db,
		)

	metricsSvc := &metrics.Metrics{}

	toolRegistry :=
		tools.NewRegistry()

	toolRegistry.Register(
		tools.NewDateTool(),
	)

	toolRegistry.Register(
		tools.NewCalculatorTool(),
	)

	agentSvc :=
		agent.NewAgent(
			toolRegistry,
		)

	mockProvider :=
		provider.NewMockProvider(
			agentSvc,
		)

	chatHandler :=
		api.NewChatHandler(
			mockProvider,
			repo,
			metricsSvc,
		)

	healthHandler :=
		api.NewHealthHandler()

	metricsHandler :=
		api.NewMetricsHandler(
			metricsSvc,
		)

	conversationHandler :=
		api.NewConversationHandler(
			repo,
		)

	mux := http.NewServeMux()

	mux.Handle(
		"/chat",
		chatHandler,
	)

	mux.Handle(
		"/health",
		healthHandler,
	)

	mux.Handle(
		"/metrics",
		metricsHandler,
	)

	mux.Handle(
		"/conversations/",
		conversationHandler,
	)

	log.Println(
		"=================================",
	)
	log.Println(
		"Agent Gateway Started",
	)
	log.Println(
		"Listening on :8080",
	)
	log.Println(
		"=================================",
	)

	err = http.ListenAndServe(
		":8080",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}
