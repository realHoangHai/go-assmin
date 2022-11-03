package service

import (
	"github.com/gin-gonic/gin"
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

// Welcome godoc
// @Summary Homepage Go Assmin Backend Side.
// @Description Welcome.
// @Tags public
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (s *Service) Welcome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"success": true})
	}
}
