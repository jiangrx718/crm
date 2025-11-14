package markdown

import (
	"web/gopkg/gins"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MarkdownExtractTitle(ctx *gin.Context) {
	markdownContent := `# Abstract -- 一级标题

1--11Academic research papers are typically published in a two-column PDF format that is difficult to read on mobile devices. Scholar uses a state-of-the-art algorithm to transform complex PDFs into a seamless, eye-friendly reading experience on your mobile and PC to solve the pain point.

# I. INTRODUCTION -- 一级标题

Over time, the format of research papers has evolved to meet the changing needs of readers and publishers. Early papers were often handwritten or printed in small quantities, making them difficult to distribute widely. As printing technology improved, papers were printed in larger quantities and distributed more widely, making it easier for researchers to share their findings with colleagues around the world.

# II. MAIN FEATURES -- 一级标题

Scholar offers a range of features that set it apart from other similar tools. Its main features include: Reflow Mode, Interlinear Translation, Integration with Zoter, Citation Popups, Clickable Content Navigation, Customizable Text Size, Reading Lists, Annotation, Dark Mode, Cross-platform, and Browser Plugin to Save References.

## A. Reflow mode -- 二级标题

One of the most significant features is the "Reflow Mode". It automatically transforms PDFs from a hard-to-read, two-column PDF format to a mobile-friendly, one-page mobile reading format. This allows readers to smoothly zoom in and out of figures, tables, formulas, and text without squinting.

## B. Interlinear Translation -- 二级标题

Scholar now facilitates the understanding of foreign research papers through the Interlinear Translation feature. This feature seamlessly integrates advanced AI translation engines, enabling readers to compare the original text with translated content in real time.

## C. Integration with Zoter -- 二级标题

Scholar seamlessly integrates with Zoter, enabling one-click synchronization of academic papers from Zoter to the Scholar platform. This integration ensures convenient access to the Zoter library within Scholar.

### D. Citation Popups -- 三级标题

Scholar also makes it easy to view citations and bibliographic information without jumping to the bottom of the page. With just a click, you can view the citation and the corresponding bibliographic information simultaneously as shown in Figure 3.

#### E. Citation Popups -- 四级标题

Scholar

##### F. Citation Popups -- 五级标题

Scholar

##### G. This Is Fifth Title -- 五级标题

Scholar

# III. BENEFITS -- 一级标题

Scholar provides several benefits to users. First, it eliminates the frustration of zooming in and out to read tiny text, making the reading experience more enjoyable.

# IV. CONCLUSION -- 一级标题

Scholar is an innovative mobile application that transforms the way people read and engage with research papers.


<!--内部使用勿删 --><div style="font-size: 10px; text-align: center; margin: 0.1em 0; padding: 0.1em; color: #999;">16</div>

`
	res, err := h.markdownService.MarkdownExtractTitle(ctx, markdownContent)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
