package middleware

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"

	logger "go-kratos/pkg/logger"
)

func RequestMiddleware(logger *logger.Logger) middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			var kind string
			var operation string
			var requestHeader transport.Header
			var replyHeader transport.Header

			start := time.Now()
			resp, err := handler(ctx, req)
			cost := time.Since(start)

			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
				requestHeader = info.RequestHeader()
				replyHeader = info.ReplyHeader()
			}

			logger.Info("grpc",
				"since", cost,
				"kind", kind,
				"operation", operation,
				"request", req,
				"request-head", requestHeader,
				"response-head", replyHeader,
				"response", resp,
				"trace_id", tracing.TraceID()(ctx),
			)

			return

		}

	}
}
