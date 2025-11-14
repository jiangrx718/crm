package markdown

import (
	"context"
	"web/gopkg/services"
	"web/gopkg/utils/md"
)

func (s *Service) MarkdownExtractTitleSection(ctx context.Context, markdown string) (services.Result, error) {

	type SectionInfo struct {
		Count     int      `json:"count"`
		TextListO []string `json:"text_list_o"`
		TextListC []string `json:"text_list_c"`
	}
	//textList := make([]string, 0)
	section, err := md.ExtractSections(markdown)
	if err != nil {
		return nil, err
	}
	////paragraphTexts := md.ExtractParagraphsByMarkerAsText(markdown)
	//textList = append(textList, paragraphTexts...)
	//
	//textListClean := make([]string, 0)
	//for _, text := range textList {
	//	textTemp := md.CleanString(text)
	//	textListClean = append(textListClean, textTemp)
	//}
	//
	//var sectionInfo SectionInfo
	//sectionInfo.TextListO = textList
	//sectionInfo.TextListC = textListClean
	//sectionInfo.Count = len(textList)

	return services.Success(ctx, section)
}
