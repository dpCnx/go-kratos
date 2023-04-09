package service

import (
	"context"

	"go-kratos/api/v1"
	logger "go-kratos/pkg/logger"
)

var _ api.HelloServer2Server = (*Demo2)(nil)

type Demo2 struct {
	api.UnimplementedHelloServer2Server

	log *logger.Logger
}

func NewDem2(log *logger.Logger) *Demo2 {
	return &Demo2{log: log}
}

func (d Demo2) SayHelloServer2(ctx context.Context, request *api.Hello2Request) (*api.Hello2Response, error) {

	return &api.Hello2Response{
		Msg:     request.Name,
		Message: request.Name,
	}, nil
}
