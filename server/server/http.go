package server

import (
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	"go-kratos/api/v1"
	"go-kratos/conf"
	"go-kratos/internal/service"
	"go-kratos/middleware"
	logger "go-kratos/pkg/log"

	"github.com/go-kratos/swagger-api/openapiv2"
)

func NewHTTPServer(config *conf.Config, log *logger.Logger, demo *service.Demo, demo2 *service.Demo2) *http.Server {

	server := http.NewServer(
		http.Address(config.Http.Address),
		http.Middleware(
			tracing.Server(),
			mmd.Server(),
			middleware.RequestMiddleware(log),
			middleware.Recover(log),
			validate.Validator(),
		),
	)

	openApi := openapiv2.NewHandler()
	server.HandlePrefix("/q/", openApi)

	api.RegisterHelloServerHTTPServer(server, demo)
	api.RegisterHelloServer2HTTPServer(server, demo2)
	return server
}
