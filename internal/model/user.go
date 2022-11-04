package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/pkg/util"
	"strings"
	"time"
)

const (
	UserDisable = iota
	UserActive
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	r.Password = strings.TrimSpace(r.Password)
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 0)),
	)
}

type LoginResponse struct {
	SessionID              uuid.UUID `json:"session_id"`
	AccessToken            string    `json:"access_token"`
	AccessTokenExpireTime  time.Time `json:"access_token_expire_time"`
	RefreshToken           string    `json:"refresh_token"`
	RefreshTokenExpireTime time.Time `json:"refresh_token_expire_time"`
}

type RenewTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewTokenResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpireTime time.Time `json:"access_token_expire_time"`
}

func (r *RenewTokenRequest) Validate() error {
	r.RefreshToken = strings.TrimSpace(r.RefreshToken)
	return validation.ValidateStruct(r,
		validation.Field(&r.RefreshToken, validation.Required),
	)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.Password = strings.TrimSpace(r.Password)
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Phone, validation.Required, validation.By(util.ValidatePhone)),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 0)),
	)
}
