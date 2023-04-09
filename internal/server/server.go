package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"

	"go-kratos/conf"
	"go-kratos/pkg/jeager"
)

type Server struct {
	rpc            *GrpcServer
	http           *HTTPServer
	r              *etcd.Registry
	config         *conf.Config
	tracerDownFunc jeager.TracerDownFunc
}

func NewServer(rpc *GrpcServer, http *HTTPServer, r *etcd.Registry, config *conf.Config, tracerDownFunc jeager.TracerDownFunc) *Server {
	return &Server{
		rpc:            rpc,
		http:           http,
		r:              r,
		config:         config,
		tracerDownFunc: tracerDownFunc,
	}
}

func (s *Server) getKServer() *kratos.App {

	return kratos.New(
		kratos.Server(
			s.rpc.getGrpcServer(),
		),
		kratos.Registrar(s.r),
		kratos.Name(s.config.Service.Name),
	)
}

func (s *Server) Run() {

	defer func() {
		s.tracerDownFunc()
	}()

	go func() {
		_ = s.getKServer().Run()
	}()
	s.http.run()

}
