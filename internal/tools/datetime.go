package tools

import (
	"context"
	"time"
)

type DateTool struct{}

func NewDateTool() *DateTool {
	return &DateTool{}
}

func (d *DateTool) Name() string {
	return "date"
}

func (d *DateTool) Execute(
	ctx context.Context,
	input string,
) (string, error) {

	return time.Now().
		Format(time.RFC1123), nil
}
