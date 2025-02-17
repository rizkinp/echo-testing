package utils

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer() *trace.TracerProvider {
	ctx := context.Background() //+
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"),
	)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil { //+
		log.Fatalf("failed to create exporter: %v", err) //+
	}
	tp := trace.NewTracerProvider( //+
		trace.WithBatcher(exporter), //+
		trace.WithResource(resource.NewWithAttributes( //+
			semconv.SchemaURL,                         //+
			semconv.ServiceNameKey.String("echo-app"), //+
		)), //+
	) //+
	otel.SetTracerProvider(tp) //+
	return tp                  //+
}
