package demo

import (
	"crm/internal/dao"
	"crm/internal/dao/demo"
)

type Service struct {
	demoDao dao.Demo
}

func NewService() *Service {
	return &Service{
		demoDao: demo.NewDao(),
	}
}
