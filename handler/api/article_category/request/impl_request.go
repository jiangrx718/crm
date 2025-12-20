package request

import "crm/gopkg/utils/httputil"

// ArticleCategoryCreateReq 创建参数
type ArticleCategoryCreateReq struct {
	ParentId      string `json:"parent_id"`
	CategoryImage string `json:"category_image"`
	Position      int    `json:"position"`
	CategoryName  string `json:"category_name" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

// ArticleCategoryUpdateReq 修改参数
type ArticleCategoryUpdateReq struct {
	CategoryId    string `json:"category_id" binding:"required"`
	ParentId      string `json:"parent_id"`
	CategoryImage string `json:"category_image"`
	Position      int    `json:"position"`
	CategoryName  string `json:"category_name" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

// ArticleCategoryDeleteReq 删除参数
type ArticleCategoryDeleteReq struct {
	CategoryId string `json:"category_id" binding:"required"`
}

// ListQuery 列表参数
type ListQuery struct {
	httputil.Pagination
}

const (
	MaxLimit = 100
)
