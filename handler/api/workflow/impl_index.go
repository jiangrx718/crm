package workflow

import (
	"io"
	"web/gopkg/utils/httputil"

	"github.com/gin-gonic/gin"
)

type ChatReq struct {
	Question string `json:"question" binding:"required"`
}

func (h *Handler) ChatStream(ctx *gin.Context) {
	var reqBody ChatReq
	if err := ctx.Bind(&reqBody); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	ch, err := h.workflowService.ChatStream(ctx, reqBody.Question)
	if err != nil {
		httputil.BadRequest(ctx, err)
		return
	}
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-ch; ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})

	return
}
