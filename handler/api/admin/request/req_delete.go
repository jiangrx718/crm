package request

type AdminDeleteReq struct {
	AdminId string `json:"admin_id" binding:"required"`
}
