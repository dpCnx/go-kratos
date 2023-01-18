//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"go-kratos/conf"
	"go-kratos/internal/repo"
	"go-kratos/internal/service"
	"go-kratos/pkg/etcd"
	"go-kratos/pkg/jeager"
	logger "go-kratos/pkg/log"
	"go-kratos/server"
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
	service.NewDem2,

	server.NewHTTPServer,
	server.NewGrpcServer,
	server.NewKratosServer,
)

func App() *kratos.App {
	panic(wire.Build(providerSet))
}
