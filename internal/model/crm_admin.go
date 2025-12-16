package model

import (
	"time"
)

type CRMAdmin struct {
	Id           int       `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	UserName     string    `gorm:"column:user_name;type:varchar(256);default:'';comment:用户名" json:"user_name"`
	UserPhone    string    `gorm:"column:user_phone;type:char(11);default:'';comment:手机号" json:"user_phone"`
	Password     string    `gorm:"column:password;type:char(32);default:'';comment:密码" json:"password"`
	DepartmentId int       `gorm:"column:department_id;type:int;default:0;comment:所属部门ID" json:"department_id"`
	Status       int       `gorm:"column:status;type:int;default:0;comment:状态,0禁用,1启用" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;type:time;autoCreateTime;index:idx_demo_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:time;autoUpdateTime;index:idx_demo_updated_at;comment:更新时间" json:"updated_at"`
}

func (r *CRMAdmin) TableName() string {
	return "crm_admin"
}

const (
	StatusOff = iota
	StatusOn
)
