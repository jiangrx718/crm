package markdown

import (
	"context"
	"strings"
	"web/gopkg/services"
	"web/gopkg/utils/md"
)

func (s *Service) MarkdownExtractTitle(ctx context.Context, markdown string) (services.Result, error) {

	// mdTitleNode, err := md.ExtractTOC(markdown, 5)
	mdTitleNode, err := md.ExtractTOCReader(strings.NewReader(markdown), 5)
	if err != nil {
		return nil, err
	}
	return services.Success(ctx, mdTitleNode)
}
