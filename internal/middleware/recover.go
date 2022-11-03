package middleware

import (
	"github.com/gin-gonic/gin"
	errors "github.com/realHoangHai/go-assmin/internal/common/errors"
)

func (h *Handler) Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*errors.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}
				appErr := errors.ErrInternal(err.(error), "something went wrong in the server")
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
