package middleware

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	logger "go-kratos/log"
)

func Recover(logger *logger.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			defer func() {
				if e := recover(); e != nil {

					var (
						kind      string
						operation string
					)

					if info, ok := transport.FromServerContext(ctx); ok {
						kind = info.Kind().String()
						operation = info.Operation()
					}

					logger.Error("recovery",
						"kind", kind,
						"operation", operation,
						"error", fmt.Sprint(err),
						"stack", string(debug.Stack()),
					)

					errors.InternalServer("SystemError", fmt.Sprintf("panic err: %v", err))
				}
			}()

			return handler(ctx, req)
		}
	}
}
