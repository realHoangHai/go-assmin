package token

import (
	"github.com/google/wire"
	"github.com/realHoangHai/go-assmin/config"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider/paseto"
)

var (
	ProviderTokenSet = wire.NewSet(NewTokenProvider)
)

func NewTokenProvider() tokenprovider.TokenMaker {
	p := paseto.NewPasetoProvider(config.C.Token.SecretKey)
	//p := jwt.NewJWTProvider(config.C.Token.SecretKey)
	return p
}
