//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/realHoangHai/go-assmin/cmd/server"
	"github.com/realHoangHai/go-assmin/internal/core"
	"github.com/realHoangHai/go-assmin/internal/middleware"
	"github.com/realHoangHai/go-assmin/internal/repo"
	"github.com/realHoangHai/go-assmin/internal/service"
	"github.com/realHoangHai/go-assmin/pkg/redis"
)

func initializeServer(ctx context.Context) (*server.Server, error) {
	wire.Build(
		repo.ProviderRepoSet,
		redis.ProviderRedisSet,
		core.InitTokenMaker,
		middleware.NewHandler,
		service.NewService,
		server.NewServer,
	)

	return new(server.Server), nil
}
