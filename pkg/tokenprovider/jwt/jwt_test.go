package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	errors "github.com/realHoangHai/go-assmin/internal/common/errors"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
	"github.com/realHoangHai/go-assmin/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	provider := NewJWTProvider(util.RandomString(32))

	userId := util.RandomUUID()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := provider.CreateToken(userId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = provider.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, userId, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	provider := NewJWTProvider(util.RandomString(32))

	token, payload, err := provider.CreateToken(util.RandomUUID(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = provider.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errors.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := tokenprovider.NewPayload(util.RandomUUID(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	provider := NewJWTProvider(util.RandomString(32))
	require.NoError(t, err)

	payload, err = provider.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errors.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
