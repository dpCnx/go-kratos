package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"go-kratos/api/v1"
	"go-kratos/conf"
	"go-kratos/internal/component/response"
	"go-kratos/internal/component/swagger"
	"go-kratos/internal/service"
	logger "go-kratos/pkg/logger"
)

type HTTPServer struct {
	config *conf.Config
	log    *logger.Logger
	demo   *service.Demo
}

func NewHTTPServer(config *conf.Config, log *logger.Logger, demo *service.Demo) *HTTPServer {

	return &HTTPServer{
		config: config,
		log:    log,
		demo:   demo,
	}
}

func (h *HTTPServer) router() http.Handler {

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(gin.Recovery())

	engine.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// 单独service执行的中间件
	var midis []gin.HandlerFunc

	api.RegisterHelloServerHTTPServer(engine, h.demo, &response.Response{}, midis...)

	// swagger handler
	swagger.Router(engine, engine.BasePath())

	return engine
}

func (h *HTTPServer) run() {

	h.log.Debug(fmt.Sprintf("server start :%s", h.config.Http.Address))

	srv := &http.Server{
		Addr:           h.config.Http.Address,
		Handler:        h.router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			h.log.Debug(fmt.Sprintf("server err %v", err))
			return
		}
	}()

	quit := make(chan os.Signal, 1)

	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	h.log.Debug("server stop")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		h.log.Debug(fmt.Sprintf("server shutdown err %v", err))
		return
	}

	h.log.Debug("server quit")
}
