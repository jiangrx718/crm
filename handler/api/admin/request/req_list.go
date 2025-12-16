package request

import "crm/gopkg/utils/httputil"

type ListQuery struct {
	httputil.Pagination
	Status    string `json:"status" form:"status"`
	UserPhone string `json:"user_phone" form:"user_phone"`
}

const (
	MaxLimit = 100
)
