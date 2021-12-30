package server

import (
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go-kratos/api/v1"
	"go-kratos/conf"
	"go-kratos/internal/service"
	logger "go-kratos/log"
	"go-kratos/middleware"
)

func NewGrpcServer(config *conf.Config, log *logger.Logger, demo *service.Demo) *grpc.Server {

	server := grpc.NewServer(
		grpc.Address(config.Grpc.Address),
		grpc.Middleware(
			tracing.Server(),
			mmd.Server(),
			middleware.RequestMiddleware(log),
			middleware.Recover(log),
		),
		// grpc.Logger(logger.NewLogger(config)),
	)

	api.RegisterHelloServerServer(server, demo)

	return server
}
