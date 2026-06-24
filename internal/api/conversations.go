package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mayankanup/go-agent-gateway/internal/repository"
)

type ConversationHandler struct {
	repo repository.ConversationRepository
}

func NewConversationHandler(
	repo repository.ConversationRepository,
) *ConversationHandler {

	return &ConversationHandler{
		repo: repo,
	}
}

func (h *ConversationHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

	id :=
		strings.TrimPrefix(
			r.URL.Path,
			"/conversations/",
		)

	messages, err :=
		h.repo.GetConversation(id)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		messages,
	)
}
