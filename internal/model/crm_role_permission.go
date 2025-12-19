package model

import "time"

// CRMRolePermission 角色权限表
type CRMRolePermission struct {
	Id           int       `gorm:"column:id;type:int;not null;primaryKey;autoIncrement;comment:自增主键;" json:"id"`
	RoleId       string    `gorm:"column:role_id;type:char(36);not null;default:'';comment:角色ID;" json:"role_id"`
	PermissionId string    `gorm:"column:permission_id;type:char(36);not null;default:'';comment:权限ID" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (r *CRMRolePermission) TableName() string {
	return "crm_role_permission"
}
