package service

import (
	"context"
	"web/internal/service/workflow"
)

type Workflow interface {
	ChatStream(ctx context.Context, question string) (chan workflow.ChatStream, error)
}
