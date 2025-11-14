package service

import (
	"context"
	"web/gopkg/services"
)

type Markdown interface {
	MarkdownExtractFile(ctx context.Context, markdown string) (services.Result, error)
	MarkdownExtractTitle(ctx context.Context, markdown string) (services.Result, error)
	MarkdownExtractSection(ctx context.Context, markdown string) (services.Result, error)
	MarkdownExtractTitleSection(ctx context.Context, markdown string) (services.Result, error)
}
