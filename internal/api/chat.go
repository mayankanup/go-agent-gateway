package api

import (
	"encoding/json"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/models"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
)

type ChatHandler struct {
	provider provider.LLMProvider
}

func NewChatHandler(
	p provider.LLMProvider,
) *ChatHandler {
	return &ChatHandler{
		provider: p,
	}
}

func (h *ChatHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var req models.ChatRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	response, err := h.provider.Chat(
		r.Context(),
		req.Message,
	)

	if err != nil {
		http.Error(
			w,
			"provider error",
			http.StatusInternalServerError,
		)
		return
	}

	resp := models.ChatResponse{
		Response: response,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}
