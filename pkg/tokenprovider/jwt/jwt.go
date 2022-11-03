package jwt

import (
	bie "errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	errors "github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
	"time"
)

const minSecretKeySize = 32

// jwtProvider is a JSON Web TokenMaker maker
type jwtProvider struct {
	secretKey string
}

// NewJWTProvider creates a new jwt provider
func NewJWTProvider(secretKey string) tokenprovider.TokenMaker {
	if len(secretKey) < minSecretKeySize {
		panic(fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize))
	}
	return &jwtProvider{secretKey}
}

func (p *jwtProvider) CreateToken(userId string, duration time.Duration) (string, *tokenprovider.Payload, error) {
	payload, err := tokenprovider.NewPayload(userId, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(p.secretKey))
	return token, payload, err
}

func (p *jwtProvider) VerifyToken(token string) (*tokenprovider.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.ErrInvalidToken
		}
		return []byte(p.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &tokenprovider.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && bie.Is(verr.Inner, errors.ErrExpiredToken) {
			return nil, errors.ErrExpiredToken
		}
		return nil, errors.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*tokenprovider.Payload)
	if !ok {
		return nil, errors.ErrInvalidToken
	}

	return payload, nil
}
