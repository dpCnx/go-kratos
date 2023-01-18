package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"go-kratos/conf"
	"go-kratos/pkg/jeager"
)

func NewKratosServer(rpc *grpc.Server, http *http.Server, r *etcd.Registry,
	config *conf.Config, tracerDownFunc jeager.TracerDownFunc) *kratos.App {

	return kratos.New(
		kratos.Server(
			rpc,
			http,
		),
		kratos.Registrar(r),
		kratos.Name(config.Service.Name),
		// kratos.Logger(logger.NewLogger(config)),
	)
}
