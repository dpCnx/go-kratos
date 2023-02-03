//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"go-kratos/conf"
	"go-kratos/internal/repo"
	"go-kratos/internal/server"
	"go-kratos/internal/service"
	"go-kratos/pkg/etcd"
	"go-kratos/pkg/jeager"
	logger "go-kratos/pkg/logger"
)

var providerSet = wire.NewSet(

	conf.LoadConfig,
	logger.NewLogger,
	jeager.TracerProvider,
	etcd.NewEtcd,

	repo.NewMysql,
	repo.NewRedis,
	repo.NewData,

	service.NewDemo,

	server.NewHTTPServer,
	server.NewGrpcServer,
	server.NewServer,
)

func App() *server.Server {
	panic(wire.Build(providerSet))
}
