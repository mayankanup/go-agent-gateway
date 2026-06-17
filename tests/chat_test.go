package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mayankanup/go-agent-gateway/internal/api"
	"github.com/mayankanup/go-agent-gateway/internal/memory"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
)

func TestChatEndpoint(t *testing.T) {

	mockProvider := provider.NewMockProvider()

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
