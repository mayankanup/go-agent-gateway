package api

import (
	"encoding/json"
	"net/http"

	"github.com/mayankanup/go-agent-gateway/internal/metrics"
	"github.com/mayankanup/go-agent-gateway/internal/models"
	"github.com/mayankanup/go-agent-gateway/internal/provider"
	"github.com/mayankanup/go-agent-gateway/internal/repository"
)

type ChatHandler struct {
	provider provider.LLMProvider
	repo     repository.ConversationRepository
	metrics  *metrics.Metrics
}

func NewChatHandler(
	p provider.LLMProvider,
	repo repository.ConversationRepository,
	metrics *metrics.Metrics,
) *ChatHandler {

	return &ChatHandler{
		provider: p,
		repo:     repo,
		metrics:  metrics,
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

	h.metrics.IncrementRequests()

	var req models.ChatRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if req.ConversationID == "" {
		http.Error(
			w,
			"conversationId is required",
			http.StatusBadRequest,
		)
		return
	}

	if req.Message == "" {
		http.Error(
			w,
			"message is required",
			http.StatusBadRequest,
		)
		return
	}

	userMessage := models.Message{
		Role:    "user",
		Content: req.Message,
	}

	err = h.repo.AppendMessage(
		req.ConversationID,
		userMessage,
	)

	if err != nil {
		http.Error(
			w,
			"failed to save message",
			http.StatusInternalServerError,
		)
		return
	}

	history, err := h.repo.GetConversation(
		req.ConversationID,
	)

	if err != nil {
		http.Error(
			w,
			"failed to load history",
			http.StatusInternalServerError,
		)
		return
	}

	response, err := h.provider.Chat(
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

	err = h.repo.AppendMessage(
		req.ConversationID,
		assistantMessage,
	)

	if err != nil {
		http.Error(
			w,
			"failed to save response",
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

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}
