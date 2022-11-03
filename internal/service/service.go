package service

import (
	"github.com/realHoangHai/go-assmin/internal/middleware"
	"github.com/realHoangHai/go-assmin/internal/repo"
)

type Service struct {
	repo    repo.IRepo
	handler *middleware.Handler
}

func NewService(repo repo.IRepo, handler *middleware.Handler) *Service {
	return &Service{
		repo:    repo,
		handler: handler,
	}
}
