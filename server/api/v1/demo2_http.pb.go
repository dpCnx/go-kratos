// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.21.10
// source: api/v1/demo2.proto

package api

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationHelloServer2SayHelloServer2 = "/api.v1.HelloServer2/SayHelloServer2"

type HelloServer2HTTPServer interface {
	SayHelloServer2(context.Context, *Hello2Request) (*Hello2Response, error)
}

func RegisterHelloServer2HTTPServer(s *http.Server, srv HelloServer2HTTPServer) {
	r := s.Route("/")
	r.POST("/say-hello2-post", _HelloServer2_SayHelloServer20_HTTP_Handler(srv))
}

func _HelloServer2_SayHelloServer20_HTTP_Handler(srv HelloServer2HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Hello2Request
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHelloServer2SayHelloServer2)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHelloServer2(ctx, req.(*Hello2Request))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Hello2Response)
		return ctx.Result(200, reply)
	}
}

type HelloServer2HTTPClient interface {
	SayHelloServer2(ctx context.Context, req *Hello2Request, opts ...http.CallOption) (rsp *Hello2Response, err error)
}

type HelloServer2HTTPClientImpl struct {
	cc *http.Client
}

func NewHelloServer2HTTPClient(client *http.Client) HelloServer2HTTPClient {
	return &HelloServer2HTTPClientImpl{client}
}

func (c *HelloServer2HTTPClientImpl) SayHelloServer2(ctx context.Context, in *Hello2Request, opts ...http.CallOption) (*Hello2Response, error) {
	var out Hello2Response
	pattern := "/say-hello2-post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationHelloServer2SayHelloServer2))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
