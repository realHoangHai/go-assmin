package middleware

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/repo"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
)

type Handler struct {
	ctx           context.Context
	repo          repo.IRepo
	redis         redis.UniversalClient
	tokenProvider tokenprovider.TokenMaker
}

func NewHandler(ctx context.Context, iRepo repo.IRepo, redis redis.UniversalClient, provider tokenprovider.TokenMaker) *Handler {
	return &Handler{
		ctx:           ctx,
		repo:          iRepo,
		redis:         redis,
		tokenProvider: provider,
	}
}

func (h *Handler) Empty() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (h *Handler) Cors() gin.HandlerFunc {
	return cors.Default()
}

func (h *Handler) NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(errors.ErrMethodNotAllowed)
	}
}

func (h *Handler) NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(errors.ErrEntityNotFound("page", fmt.Errorf("not found")))
	}
}
