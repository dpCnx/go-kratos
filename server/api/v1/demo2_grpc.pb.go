// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: api/v1/demo2.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloServer2Client is the client API for HelloServer2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServer2Client interface {
	SayHelloServer2(ctx context.Context, in *Hello2Request, opts ...grpc.CallOption) (*Hello2Response, error)
}

type helloServer2Client struct {
	cc grpc.ClientConnInterface
}

func NewHelloServer2Client(cc grpc.ClientConnInterface) HelloServer2Client {
	return &helloServer2Client{cc}
}

func (c *helloServer2Client) SayHelloServer2(ctx context.Context, in *Hello2Request, opts ...grpc.CallOption) (*Hello2Response, error) {
	out := new(Hello2Response)
	err := c.cc.Invoke(ctx, "/api.v1.HelloServer2/SayHelloServer2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer2Server is the server API for HelloServer2 service.
// All implementations must embed UnimplementedHelloServer2Server
// for forward compatibility
type HelloServer2Server interface {
	SayHelloServer2(context.Context, *Hello2Request) (*Hello2Response, error)
	mustEmbedUnimplementedHelloServer2Server()
}

// UnimplementedHelloServer2Server must be embedded to have forward compatible implementations.
type UnimplementedHelloServer2Server struct {
}

func (UnimplementedHelloServer2Server) SayHelloServer2(context.Context, *Hello2Request) (*Hello2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloServer2 not implemented")
}
func (UnimplementedHelloServer2Server) mustEmbedUnimplementedHelloServer2Server() {}

// UnsafeHelloServer2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer2Server will
// result in compilation errors.
type UnsafeHelloServer2Server interface {
	mustEmbedUnimplementedHelloServer2Server()
}

func RegisterHelloServer2Server(s grpc.ServiceRegistrar, srv HelloServer2Server) {
	s.RegisterService(&HelloServer2_ServiceDesc, srv)
}

func _HelloServer2_SayHelloServer2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer2Server).SayHelloServer2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.HelloServer2/SayHelloServer2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer2Server).SayHelloServer2(ctx, req.(*Hello2Request))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloServer2_ServiceDesc is the grpc.ServiceDesc for HelloServer2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloServer2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.HelloServer2",
	HandlerType: (*HelloServer2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloServer2",
			Handler:    _HelloServer2_SayHelloServer2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/demo2.proto",
}
