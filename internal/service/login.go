package service

import (
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/common/request"
	"github.com/realHoangHai/go-assmin/internal/common/response"
	"github.com/realHoangHai/go-assmin/internal/model"
)

var sample = &model.LoginRequest{
	Email:    "a@gmail.com",
	Password: "12345678",
}

func (s *Service) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *model.LoginRequest
		if err := request.ParseJson(c, &req); err != nil {
			panic(err)
		}
		if err := req.Validate(); err != nil {
			panic(errors.ErrInvalidRequest(err, "invalid request"))
		}
		if req.Email != sample.Email || req.Password == sample.Password {
			panic(errors.ErrInvalidRequest(nil, "email or password incorrect"))
		}
		c.JSON(200, response.SimpleSuccess(true))
	}
}
