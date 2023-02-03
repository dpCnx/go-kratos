package server

import (
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"go-kratos/api/v1"
	"go-kratos/conf"
	middleware2 "go-kratos/internal/component/middleware"
	"go-kratos/internal/service"
	logger "go-kratos/pkg/logger"
)

type GrpcServer struct {
	config *conf.Config
	log    *logger.Logger
	demo   *service.Demo
}

func NewGrpcServer(config *conf.Config, log *logger.Logger, demo *service.Demo) *GrpcServer {
	return &GrpcServer{
		config: config,
		log:    log,
		demo:   demo,
	}
}

func (g *GrpcServer) getGrpcServer() *grpc.Server {

	server := grpc.NewServer(
		grpc.Address(g.config.Grpc.Address),
		grpc.Middleware(
			tracing.Server(),
			mmd.Server(),
			middleware2.RequestMiddleware(g.log),
			middleware2.Recover(g.log),
		),
	)

	api.RegisterHelloServerServer(server, g.demo)

	return server
}
