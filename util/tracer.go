package util

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

const (
	service   = "saka-api"
	agentHost = "127.0.0.1:5775"
)

func NewJaegerTracer() (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
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
