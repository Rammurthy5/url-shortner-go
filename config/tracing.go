package config

import (
	"context"
	"fmt"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
)

var (
	onceTrace      sync.Once
	tracerProvider *sdktrace.TracerProvider
)

func initTracer() (*sdktrace.TracerProvider, error) {
	var err error
	var exp *otlptrace.Exporter

	onceTrace.Do(func() {
		client := otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("tempo:4318"),
			otlptracehttp.WithInsecure(),
		)
		exp, err = otlptrace.New(context.Background(), client)
		if err == nil {
			resource := resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName("url-shortener"),
				semconv.ServiceVersion("1.0.0"),
			)

			tracerProvider = sdktrace.NewTracerProvider(
				sdktrace.WithBatcher(exp),
				sdktrace.WithResource(resource),
			)
			otel.SetTracerProvider(tracerProvider)
		}
	})

	if err != nil {
		return nil, fmt.Errorf("creating OTLP exporter: %w", err)
	}

	return tracerProvider, nil
}

func ShutdownTracer() {
	if tracerProvider != nil {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			_logger.Error("Error shutting down tracer provider", zap.Error(err))
		}
	}
}
