package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

func (h *Handler) EmptyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
