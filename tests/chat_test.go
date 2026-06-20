package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mayankanup/go-agent-gateway/internal/agent"
	"github.com/mayankanup/go-agent-gateway/internal/api"
	"github.com/mayankanup/go-agent-gateway/internal/memory"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
	"github.com/mayankanup/go-agent-gateway/internal/tools"
)

func TestChatEndpoint(t *testing.T) {

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

	handler := api.NewChatHandler(
		mockProvider,
		repo,
	)

	body := []byte(
		`{"message":"hello"}`,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/chat",
		bytes.NewBuffer(body),
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			recorder.Code,
		)
	}
}
