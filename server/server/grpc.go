package server

import (
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"go-kratos/api/v1"
	"go-kratos/conf"
	"go-kratos/internal/service"
	"go-kratos/middleware"
	logger "go-kratos/pkg/log"
)

func NewGrpcServer(config *conf.Config, log *logger.Logger, demo *service.Demo, demo2 *service.Demo2) *grpc.Server {

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
	// s api.RegisterHelloServer2Server(server, demo2)

	return server
}
