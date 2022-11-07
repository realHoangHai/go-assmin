package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	common "github.com/realHoangHai/go-assmin/internal/common/const"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/common/request"
)

func (h *Handler) Authorize(ignorePaths ...IgnoreFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if IgnoreHandler(c, ignorePaths...) {
			c.Next()
			return
		}

		token := request.GetToken(c)
		payload, err := h.tokenProvider.VerifyToken(token)
		if err != nil {
			panic(errors.ErrUnauthorized(err))
		}
		userID := uuid.MustParse(payload.UserID)
		user, err := h.repo.GetUserByID(userID)
		if err != nil {
			panic(errors.ErrUnauthorized(err))
		}
		request.Set(c, common.Requester, user)
		c.Next()
	}
}
