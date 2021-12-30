package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go-kratos/conf"
)

func NewKratosServer(rpc *grpc.Server, r *etcd.Registry, config *conf.Config) *kratos.App {

	return kratos.New(
		kratos.Server(
			rpc,
		),
		kratos.Registrar(r),
		kratos.Name(config.Service.Name),
		// kratos.Logger(logger.NewLogger(config)),
	)
}
