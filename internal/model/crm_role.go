package model

import "time"

// CRMRole 角色表
type CRMRole struct {
	Id        int       `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	RoleId    string    `gorm:"column:role_id;type:char(36);unique;comment:业务主键;" json:"role_id"`
	RoleName  string    `gorm:"column:role_name;type:varchar(128);not null;default:'';comment:角色名称" json:"role_name"`
	Status    string    `gorm:"column:status;type:varchar(32);not null;default:'on';comment:状态,off禁用,on启用" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:time;autoCreateTime;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:time;autoUpdateTime;index:idx_updated_at;comment:更新时间" json:"updated_at"`
}

func (r *CRMRole) TableName() string {
	return "crm_role"
}
