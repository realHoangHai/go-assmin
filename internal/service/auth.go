package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/internal/common/request"
	"github.com/realHoangHai/go-assmin/internal/common/response"
	"github.com/realHoangHai/go-assmin/internal/ent"
	"github.com/realHoangHai/go-assmin/internal/model"
	"github.com/realHoangHai/go-assmin/pkg/util"
	"time"
)

// Login godoc
// @Summary Login.
// @Description Login.
// @Tags users
// @Accept json
// @Produce json
// @Param users body model.LoginRequest true "Login request"
// @Success 200 {object} response.Success
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/login [post]
func (s *Service) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// declare payload
		var req *model.LoginRequest
		// parse
		request.ParseJson(c, &req)
		// validate request
		if err := req.Validate(); err != nil {
			panic(errors.ErrInvalidRequest(err))
		}
		// get user info from database
		user, err := s.repo.GetUserNotDisableByEmail(req.Email)
		if err != nil {
			panic(err)
		}
		// compare password
		if err := util.CheckPassword(req.Password, user.Password); err != nil {
			panic(errors.ErrInvalidRequest(fmt.Errorf("email or password is incorrect")))
		}
		// generate access token
		accessToken, accessPayload, err := s.tokenprovider.CreateToken(user.ID.String(), time.Hour*24) // 1 day
		if err != nil {
			panic(errors.ErrInternal(err))
		}
		// generate refresh token
		refreshToken, refreshPayload, err := s.tokenprovider.CreateToken(user.ID.String(), time.Hour*24*5) // 5 days
		if err != nil {
			panic(errors.ErrInternal(err))
		}
		session := &ent.Session{
			ID:           refreshPayload.ID,
			UserID:       user.ID,
			RefreshToken: refreshToken,
			UserAgent:    c.Request.UserAgent(),
			ClientIP:     c.ClientIP(),
			IsBlocked:    false,
			ExpireTime:   refreshPayload.ExpiredAt,
		}

		data, err := s.repo.CreateSession(session)
		if err != nil {
			panic(err)
		}

		result := &model.LoginResponse{
			SessionID:              data.ID,
			AccessToken:            accessToken,
			AccessTokenExpireTime:  accessPayload.ExpiredAt,
			RefreshToken:           refreshToken,
			RefreshTokenExpireTime: refreshPayload.ExpiredAt,
		}

		c.JSON(200, response.SimpleSuccess(result))
	}
}

// RenewToken godoc
// @Summary Renew access token.
// @Description Renew access token.
// @Tags users
// @Accept json
// @Produce json
// @Param users body model.RenewTokenRequest true "Renew access token request"
// @Success 200 {object} response.Success
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /api/renew-token [post]
func (s *Service) RenewToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bind request
		var req model.RenewTokenRequest
		// parse json
		request.ParseJson(c, &req)
		// validate
		payload, err := s.tokenprovider.VerifyToken(req.RefreshToken)
		if err != nil {
			panic(errors.ErrUnauthorized(err))
		}
		// get session by id
		session, err := s.repo.GetSessionByID(payload.ID)
		if err != nil {
			panic(err)
		}
		// check if session is blocked
		if session.IsBlocked {
			panic(errors.ErrUnauthorized(fmt.Errorf("bloked session")))
		}
		// check session user id
		if (session.UserID).String() != payload.UserID {
			panic(errors.ErrUnauthorized(fmt.Errorf("invalid session user")))
		}
		// check session refresh token vs request
		if session.RefreshToken != req.RefreshToken {
			panic(errors.ErrUnauthorized(fmt.Errorf("mismatched refresh token")))
		}

		if time.Now().After(session.ExpireTime) {
			panic(errors.ErrUnauthorized(fmt.Errorf("expired session")))
		}

		// account is verified
		accessToken, accessPayload, err := s.tokenprovider.CreateToken(payload.UserID, time.Hour*24) // 1 day
		if err != nil {
			panic(errors.ErrInternal(err))
		}

		result := &model.RenewTokenResponse{
			AccessToken:           accessToken,
			AccessTokenExpireTime: accessPayload.ExpiredAt,
		}
		c.JSON(200, response.SimpleSuccess(result))
	}
}
