package request

type AdminCreateReq struct {
	UserName     string `json:"user_name" binding:"required"`
	UserPhone    string `json:"user_phone" binding:"required"`
	Password     string `json:"password" binding:"required"`
	DepartmentId int    `json:"department_id" binding:"required"`
	Status       string `json:"status" binding:"required"`
}
