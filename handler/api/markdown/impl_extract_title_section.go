package markdown

import (
	"web/gopkg/gins"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MarkdownExtractTitleSection(ctx *gin.Context) {
	markdownContent := `# Abstract

Academic research papers are typically published in a two-column PDF format that is difficult to read on mobile devices. Scholar uses a state-of-the-art algorithm to transform complex PDFs into a seamless, eye-friendly reading experience on your mobile and PC to solve the pain point.

# I. INTRODUCTION

Over time, the format of research papers has evolved to meet the changing needs of readers and publishers. Early papers were often handwritten or printed in small quantities, making them difficult to distribute widely. As printing technology improved, papers were printed in larger quantities and distributed more widely, making it easier for researchers to share their findings with colleagues around the world.

![](s3://sgcc-1255000019/duplicate-checker/contrast-file/2025/0814/480/国网重庆信通公司-生产检修业务-2022年适配同源维护集成改造-集成实施项目-成本性_attachment/caa944997baf280613db83195943ab89.png)

# II. MAIN FEATURES

Scholar offers a range of features that set it apart from other similar tools. Its main features include: Reflow Mode, Interlinear Translation, Integration with Scholar, Citation Popups, Clickable Content Navigation, Customizable Text Size, Reading Lists, Annotation, Dark Mode, Cross-platform, and Browser Plugin to Save References.

## A. Reflow mode

One of the most significant features is the "Reflow Mode". It automatically transforms PDFs from a hard-to-read, two-column PDF format to a mobile-friendly, one-page mobile reading format. This allows readers to smoothly zoom in and out of figures, tables, formulas, and text without squinting.

## B. Interlinear Translation

Scholar now facilitates the understanding of foreign research papers through the Interlinear Translation feature. This feature seamlessly integrates advanced AI translation engines, enabling readers to compare the original text with translated content in real time.

## C. Integration with Zoter

Scholar seamlessly integrates with Zoter, enabling one-click synchronization of academic papers from Zoter to the Scholar platform. This integration ensures convenient access to the Scholar library within Scholar.

### C-1. Citation Popups 

Scholar also makes it easy to view citations and bibliographic information without jumping to the bottom of the page. With just a click, you can view the citation and the corresponding bibliographic information simultaneously as shown in Figure 3.

# IV. CONCLUSION

Scholar is an innovative mobile application that transforms the way people read and engage with research papers.

`
	res, err := h.markdownService.MarkdownExtractTitleSection(ctx, markdownContent)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
