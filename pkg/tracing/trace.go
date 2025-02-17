package tracing

import (
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// InitTracer menginisialisasi Jaeger Tracer
func InitTracer(service string) (opentracing.Tracer, func()) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("❌ Gagal menginisialisasi Jaeger: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, func() { _ = closer.Close() }
}
