package service

import (
	"context"
	"fmt"

	"go-kratos/api/v1"
)

var _ api.HelloServerServer = (*Demo)(nil)

type Demo struct {
	api.UnimplementedHelloServerServer
}

func (d Demo) SayHello(ctx context.Context, request *api.PostHelloRequest) (*api.PostHelloResponse, error) {
	return &api.PostHelloResponse{Msg: "demo"}, nil
}

func (d Demo) SayHello2(ctx context.Context, request *api.GetHelloRequest) (*api.GetHelloResponse, error) {
	return &api.GetHelloResponse{Msg: fmt.Sprintf("%s %s", request.GetName(), request.GetAge())}, nil
}
