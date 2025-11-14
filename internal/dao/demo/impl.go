package demo

import (
	"web/gopkg/gorms"
)

type Dao struct {
	*gorms.BaseDao
}

func NewDao() *Dao {
	return &Dao{
		BaseDao: gorms.NewBaseDao(),
	}
}
