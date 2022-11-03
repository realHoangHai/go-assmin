//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/realHoangHai/go-assmin/cmd/server"
	"github.com/realHoangHai/go-assmin/internal/middleware"
	"github.com/realHoangHai/go-assmin/internal/repo"
	"github.com/realHoangHai/go-assmin/internal/service"
	"github.com/realHoangHai/go-assmin/pkg/redis"
	"github.com/realHoangHai/go-assmin/pkg/tokenprovider/token"
)

func initializeServer(ctx context.Context) (s *server.Server, err error) {
	wire.Build(
		server.NewServer,
		service.NewService,
		middleware.NewHandler,
		token.ProviderTokenSet,
		redis.ProviderRedisSet,
		repo.ProviderRepoSet,
	)

	return s, err
}
