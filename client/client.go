package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go-kratos/api/v1"
	"go-kratos/conf"
	"go-kratos/pkg/jeager"
	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

func main() {

	// http://127.0.0.1:16686
	jeager.TracerProvider(nil, &conf.Config{
		Service: struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{
			Name: "go-kratos",
		},
		Jaeger: struct {
			Address string `json:"address"`
		}{
			"http://localhost:14268/api/traces",
		},
	})

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"10.64.146.48:49329"},
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    3 * time.Second, // 每3秒ping一次服务器
		DialKeepAliveTimeout: time.Second,     // 1秒没有返回则代表故障
		DialOptions:          []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithDiscovery(r),
		grpc.WithEndpoint("discovery:///go-kratos"),
		grpc.WithMiddleware(
			tracing.Server(),
			mmd.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewHelloServerClient(conn)

	ctx := context.Background()
	ctx = metadata.AppendToClientContext(ctx, "x-md-global-jwt", "jwt-test")
	reply, err := client.SayHello(ctx, &api.HelloRequest{Name: "p"})
	if err != nil {
		e := errors.FromError(err)
		fmt.Println(e)
		return
	}
	fmt.Println(reply.Msg)
	fmt.Println(reply.CreatedAt.AsTime().Local())
}
