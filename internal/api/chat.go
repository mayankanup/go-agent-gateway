package api

import (
	"encoding/json"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/memory"
	"github.com/mayankanup/go-agent-gateway/internal/models"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
)

type ChatHandler struct {
	provider provider.LLMProvider
	memory   memory.Repository
}

func NewChatHandler(
	p provider.LLMProvider,
	m memory.Repository,
) *ChatHandler {

	return &ChatHandler{
		provider: p,
		memory:   m,
	}
}

func (h *ChatHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req models.ChatRequest

	err :=
		json.NewDecoder(
			r.Body,
		).Decode(&req)

	if err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	userMessage := models.Message{
		Role:    "user",
		Content: req.Message,
	}

	h.memory.AppendMessage(
		req.ConversationID,
		userMessage,
	)

	history :=
		h.memory.GetConversation(
			req.ConversationID,
		)

	response, err :=
		h.provider.Chat(
			r.Context(),
			history,
		)

	if err != nil {
		http.Error(
			w,
			"provider error",
			http.StatusInternalServerError,
		)
		return
	}

	assistantMessage := models.Message{
		Role:    "assistant",
		Content: response,
	}

	h.memory.AppendMessage(
		req.ConversationID,
		assistantMessage,
	)

	resp := models.ChatResponse{
		Response: response,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}
