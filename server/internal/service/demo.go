package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/metadata"

	"go-kratos/api/v1"
	"go-kratos/internal/repo"
	logger "go-kratos/pkg/log"
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

	if "d" == req.Name {
		return nil, api.ErrorInvalidParameter("参数错误")
	}

	if md, ok := metadata.FromServerContext(ctx); ok {
		extra := md.Get("x-md-global-jwt")
		fmt.Println(extra)
	}

	if err := d.data.Insert(ctx); err != nil {
		return nil, err
	}

	if _, err := d.data.SetUser(ctx); err != nil {
		return nil, err
	}

	return &api.HelloResponse{
		Msg: "hello:" + req.Name,
	}, nil

}

func (d *Demo) SayHello2(ctx context.Context, request *api.HelloRequest) (*api.HelloResponse, error) {

	return &api.HelloResponse{
		Msg:     request.Name,
		Message: request.Name,
	}, nil
}
