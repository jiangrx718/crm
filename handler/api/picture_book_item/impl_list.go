package picture_book_item

import (
	"crm/gopkg/gins"
	"crm/handler/api/picture_book_item/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ItemList(c *gin.Context) {
	var req request.ItemListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		gins.BadRequest(c, err)
		return
	}
	res, err := h.service.ItemList(c, req.Offset, req.Limit, req.BookId)
	if err != nil {
		gins.ServerError(c, err)
		return
	}
	c.JSON(200, res)
}
