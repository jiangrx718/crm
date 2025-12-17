package request

type AdminUpdateReq struct {
	AdminId      string `json:"admin_id" binding:"required"`
	Password     string `json:"password"`
	DepartmentId int    `json:"department_id" binding:"required"`
	Status       string `json:"status" binding:"required"`
}
