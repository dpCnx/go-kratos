package etcd

import (
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"

	"go-kratos/conf"
	logger "go-kratos/pkg/logger"

	etcdc "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

func NewEtcd(log *logger.Logger, config *conf.Config) *etcd.Registry {

	client, err := etcdc.New(etcdc.Config{
		Endpoints:            []string{config.Etcd.Address},
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    3 * time.Second, // 每3秒ping一次服务器
		DialKeepAliveTimeout: time.Second,     // 1秒没有返回则代表故障
		DialOptions:          []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		panic(err)
	}

	log.Debug("etcd start successful")

	return etcd.New(client)

}
