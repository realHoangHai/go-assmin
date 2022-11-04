package service

import (
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/internal/middleware"
	"github.com/realHoangHai/go-assmin/internal/repo"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
)

type Service struct {
	repo          repo.IRepo
	handler       *middleware.Handler
	tokenprovider tokenprovider.TokenMaker
}

func NewService(repo repo.IRepo, handler *middleware.Handler, provider tokenprovider.TokenMaker) *Service {
	return &Service{
		repo:          repo,
		handler:       handler,
		tokenprovider: provider,
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
