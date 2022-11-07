package core

import (
	"github.com/realHoangHai/go-assmin/config"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider/paseto"
)

func InitTokenMaker() tokenprovider.TokenMaker {
	p := paseto.NewPasetoProvider(config.C.Token.SecretKey)
	//p := jwt.NewJWTProvider(config.C.Token.SecretKey)
	return p
}
