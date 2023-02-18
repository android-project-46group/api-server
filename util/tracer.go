package util

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

const (
	service   = "saka-api"
	agentHost = "127.0.0.1:5775"
)

func NewJaegerTracer() (opentracing.Tracer, io.Closer, error) {
	cfg := jaegerConfig.Configuration{
		ServiceName: service,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: time.Second,
			LocalAgentHostPort:  agentHost,
		},
	}

	return cfg.NewTracer()
}
