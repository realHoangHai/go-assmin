package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/common/request"
)

func (h *Handler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := request.GetToken(c)
		payload, err := h.tokenProvider.VerifyToken(token)
		if err != nil {
			panic(errors.ErrUnauthoried)
		}
		userID := uuid.MustParse(payload.UserID)
		user, err := h.repo.GetUserByID(userID)
		if err != nil {
			panic(errors.ErrUnauthoried)
		}
		fmt.Println(user)
	}
}
