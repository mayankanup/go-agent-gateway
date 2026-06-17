package models

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	ConversationID string `json:"conversationId"`
	Message        string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}
