package middleware

import (
	"github.com/gin-gonic/gin"
	errors "github.com/realHoangHai/go-assmin/internal/common/errors"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(errors.ErrMethodNotAllowed)
	}
}

func EmptyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
