package service

import (
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/common/request"
	"github.com/realHoangHai/go-assmin/internal/common/response"
	"github.com/realHoangHai/go-assmin/internal/ent"
	"github.com/realHoangHai/go-assmin/internal/model"
	"github.com/realHoangHai/go-assmin/pkg/util"
)

// CreateUser godoc
// @Summary Create user.
// @Description Create user.
// @Tags users
// @Accept json
// @Produce json
// @Param users body model.CreateUserRequest true "Create user request"
// @Success 200 {object} response.Success
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/register [post]
func (s *Service) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateUserRequest
		// parse json
		request.ParseJson(c, &req)
		// validate
		if err := req.Validate(); err != nil {
			panic(errors.ErrInvalidRequest(err))
		}
		hashedPass, err := util.HashPassword(req.Password)
		if err != nil {
			panic(errors.ErrInternal(err))
		}
		// save to db
		data := &ent.User{
			Name:     req.Name,
			Password: hashedPass,
			Phone:    req.Phone,
			Email:    req.Email,
			Status:   model.UserActive,
		}
		if err := s.repo.CreateUser(data); err != nil {
			panic(err)
		}
		c.JSON(200, response.SimpleSuccess(true))
	}
}
