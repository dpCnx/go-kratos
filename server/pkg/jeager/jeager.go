package jeager

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"go-kratos/conf"
	logger "go-kratos/pkg/logger"
)

type TracerDownFunc func()

func TracerProvider(logger *logger.Logger, config *conf.Config) TracerDownFunc {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.Jaeger.Address)))
	if err != nil {
		panic(err)
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.Service.Name),
		)),
	)
	logger.Debug("jaeger start successful")
	otel.SetTracerProvider(tp)

	return func() {
		_ = tp.Shutdown(context.Background())
	}
}
