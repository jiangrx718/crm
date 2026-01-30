package request

type BookCreateRequest struct {
	CategoryId   string `json:"category_id" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Icon         string `json:"icon" binding:"required"`
	Status       string `json:"status" binding:"required"`
	Position     int    `json:"position"`
	CategoryType int    `json:"category_type" binding:"required"`
}

type BookUpdateRequest struct {
	BookId       string `json:"book_id" binding:"required"`
	CategoryId   string `json:"category_id" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Icon         string `json:"icon" binding:"required"`
	Status       string `json:"status" binding:"required"`
	Position     int    `json:"position"`
	CategoryType int    `json:"category_type" binding:"required"`
}

type BookDeleteRequest struct {
	BookId string `json:"book_id" binding:"required"`
}

type BookStatusRequest struct {
	BookId string `json:"book_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type BookDetailRequest struct {
	BookId string `form:"book_id" binding:"required"`
}

type BookListRequest struct {
	Offset       int64  `form:"offset"`
	Limit        int64  `form:"limit"`
	Title        string `form:"title"`
	CategoryType int    `form:"category_type"`
	CategoryId   string `form:"category_id"`
}
