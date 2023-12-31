package config

import (
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// NewJaegerTracer for current service
func NewJaegerTracer() (tracer opentracing.Tracer, closer io.Closer, err error) {
	jaegerEndpoint, err := GetValue(JaegerEndpoint)
	if err != nil {
		return nil, nil, err
	}

	serviceName, err := GetValue(ServiceName)
	if err != nil {
		return nil, nil, err
	}

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint:   jaegerEndpoint,
		},
		ServiceName: serviceName,
	}

	tracer, closer, err = cfg.NewTracer()
	if err != nil {
		return
	}

	opentracing.SetGlobalTracer(tracer)
	return
}
