package markdown

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"web/gopkg/log"
	"web/gopkg/services"
	"web/gopkg/utils/md"
)

// MarkdownExtractFile 提取md中的文件地址路径
func (s *Service) MarkdownExtractFile(ctx context.Context, markdown string) (services.Result, error) {

	hlog.Infof("开始清洗知识库配置，共 %d 条记录，环境: %s", 10, "dev~")

	log.Sugar().Infof("开始清洗知识库配置，共 %d 条记录，环境: %s", 10, "dev")
	// onlyImages := md.ExtractFileLinksByType(markdown, false)

	onlyImages := md.ExtractFileLinksByTypeReader(strings.NewReader(markdown), true)
	return services.Success(ctx, onlyImages)
}
