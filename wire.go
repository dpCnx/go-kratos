// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"go-kratos/conf"
	"go-kratos/internal/repo"
	"go-kratos/internal/service"
	logger "go-kratos/log"
	"go-kratos/pkg/etcd"
	"go-kratos/pkg/jeager"
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

	server.NewGrpcServer,
	server.NewKratosServer,
)

func App() *kratos.App {
	panic(wire.Build(providerSet))
}
