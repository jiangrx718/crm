package model

import (
	"time"
	"web/gopkg/gorms"

	"gorm.io/datatypes"
)

type Demo struct {
	gorms.Entity
	DemoId      string                           `gorm:"column:demo_id;type:varchar(32);index:uni_demo_id,unique;comment:案例ID" json:"demo_id"`
	Name        string                           `gorm:"column:name;type:varchar(1024);default:'';comment:案例标题" json:"name"`
	FileType    int                              `gorm:"column:file_type;type:integer;default:0;comment:文档类型,1可研报告,2需求报告" json:"file_type"`
	ProjectType int                              `gorm:"column:project_type;type:integer;default:0;comment:项目类型,1产品购置类,2开发实施类,3数据工程类,4咨询、运维类" json:"project_type"`
	Content     string                           `gorm:"column:content;type:varchar(1024);default:'';comment:文档内容" json:"content"`
	Metadata    datatypes.JSONType[DemoMetadata] `gorm:"column:metadata;type:json;comment:元数据" json:"metadata"`
	CreatedAt   time.Time                        `gorm:"column:created_at;type:time;autoCreateTime;index:idx_demo_created_at;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time                        `gorm:"column:updated_at;type:time;autoUpdateTime;index:idx_demo_updated_at;comment:更新时间" json:"updated_at"`
}

func (r *Demo) TableName() string {
	return "demo"
}

type DemoMetadata struct {
	ProjectType string     `json:"project_type"` // 项目类型
	Databases   []Database `json:"databases"`    // 数据集列表
}

type Database struct {
	DatabaseId   string `json:"database_id"`
	DatabaseType string `json:"database_type"`
	Repository   string `json:"repository"`
}
