package provider

import "context"

type LLMProvider interface {
	Chat(
		ctx context.Context,
		message string,
	) (string, error)
}
