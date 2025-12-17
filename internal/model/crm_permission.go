package model

import (
	"time"
)

type CRMPermission struct {
	Id             int       `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:主键;" json:"id"`
	PermissionId   string    `gorm:"column:permission_id;type:char(36);unique;comment:业务主键;" json:"permission_id"`
	PermissionName string    `gorm:"column:permission_name;type:varchar(128);default:'';comment:权限名称" json:"permission_name"`
	PermissionURL  string    `gorm:"column:permission_url;type:varchar(128);default:'';comment:权限URL" json:"permission_url"`
	PermissionType int       `gorm:"column:permission_type;type:int;default:0;comment:权限类型,1菜单,2按钮,3接口" json:"permission_type"`
	ParentId       string    `gorm:"column:parent_id;type:char(36);unique;comment:父级id;" json:"parent_id"`
	Status         string    `gorm:"column:status;type:varchar(32);default:'on';comment:状态,off禁用,on启用" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at;type:time;autoCreateTime;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:time;autoUpdateTime;index:idx_updated_at;comment:更新时间" json:"updated_at"`
}

func (r *CRMPermission) TableName() string {
	return "crm_permission"
}
