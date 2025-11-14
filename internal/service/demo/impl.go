package demo

import (
	"web/internal/dao"
	"web/internal/dao/demo"
)

type Service struct {
	demoDao dao.Demo
}

func NewService() *Service {
	return &Service{
		demoDao: demo.NewDao(),
	}
}
