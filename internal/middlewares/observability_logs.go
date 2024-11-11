package middlewares

import (
	"context"
	"time"

	"github.com/FabricioCosati/onfly-test/internal/utils"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sc "go.opentelemetry.io/otel/semconv/v1.17.0"
	otrace "go.opentelemetry.io/otel/trace"
)

func InitTracerMetrics(fileName string) (*trace.TracerProvider, error) {
	utils.CreateFolderIfNotExists()
	file := utils.GetFileToSave(fileName)

	exp, err := stdouttrace.New(stdouttrace.WithWriter(file))
	if err != nil {
		return nil, err
	}

	r, err := resource.New(context.Background(),
		resource.WithAttributes(
			sc.ServiceNameKey.String("onfly-orders"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
func OtelMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		initTimer := time.Now()
		ctx.Next()
		endTimer := time.Since(initTimer)

		span := otrace.SpanFromContext(ctx.Request.Context())
		span.SetAttributes(
			attribute.Int64("http.response_time_ms", endTimer.Milliseconds()),
		)
	}
}
