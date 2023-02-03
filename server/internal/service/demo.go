package service

import (
	"context"

	"go-kratos/api/v1"
	"go-kratos/internal/repo"
	logger "go-kratos/pkg/logger"
)

var _ api.HelloServerServer = (*Demo)(nil)

type Demo struct {
	api.UnimplementedHelloServerServer

	data *repo.Data
	log  *logger.Logger
}

func NewDemo(data *repo.Data, log *logger.Logger) *Demo {
	return &Demo{data: data, log: log}
}

func (d *Demo) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {

	return nil, api.ErrorInvalidParameter("参数错误")

}

func (d *Demo) SayHello2(ctx context.Context, request *api.HelloRequest) (*api.HelloResponse, error) {

	return &api.HelloResponse{
		Msg:     request.Name,
		Message: request.Name,
	}, nil
}
