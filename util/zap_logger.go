package util

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger using zap and make outputs to a file.
//
// documentation of zap:
// https://pkg.go.dev/go.uber.org/zap
type zapLogger struct {
	level   Level
	logger  *zap.Logger
	host    string
	service string
}

func NewZapLogger(
	path string,
	host string,
	service string,
	level Level,
) (*zapLogger, func(), error) {

	cfg := zap.NewProductionConfig()

	// Output to files.
	cfg.OutputPaths = []string{
		path,
	}

	lv := zap.NewAtomicLevel()
	switch level {
	case Degub:
		lv.SetLevel(zapcore.DebugLevel)
	case Info:
		lv.SetLevel(zapcore.InfoLevel)
	case Warn:
		lv.SetLevel(zapcore.WarnLevel)
	case Error:
		lv.SetLevel(zapcore.ErrorLevel)
	case Critical:
		lv.SetLevel(zapcore.FatalLevel)
	}
	cfg.Level = lv

	// Define output format
	cfg.EncoderConfig = zapcore.EncoderConfig{
		LevelKey:   "",
		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,
		MessageKey: "message",
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, nil, err
	}

	logger := &zapLogger{
		level:   level,
		logger:  l,
		host:    host,
		service: service,
	}
	return logger, func() {
		// Nothing?
	}, nil
}

func (z *zapLogger) Critical(ctx context.Context, msg string) {
	z.Print(ctx, Critical, msg)
}

func (z *zapLogger) Error(ctx context.Context, msg string) {
	z.Print(ctx, Error, msg)
}
func (z *zapLogger) Warn(ctx context.Context, msg string) {
	z.Print(ctx, Warn, msg)
}

func (z *zapLogger) Info(ctx context.Context, msg string) {
	z.Print(ctx, Info, msg)
}

func (z *zapLogger) Debug(ctx context.Context, msg string) {
	z.Print(ctx, Degub, msg)
}

func (z *zapLogger) Criticalf(ctx context.Context, msg string, a ...interface{}) {
	z.Print(ctx, Critical, fmt.Sprintf(msg, a...))
}

func (z *zapLogger) Errorf(ctx context.Context, msg string, a ...interface{}) {
	z.Print(ctx, Error, fmt.Sprintf(msg, a...))
}

func (z *zapLogger) Warnf(ctx context.Context, msg string, a ...interface{}) {
	z.Print(ctx, Warn, fmt.Sprintf(msg, a...))
}

func (z *zapLogger) Infof(ctx context.Context, msg string, a ...interface{}) {
	z.Print(ctx, Info, fmt.Sprintf(msg, a...))
}

func (z *zapLogger) Debugf(ctx context.Context, msg string, a ...interface{}) {
	z.Print(ctx, Degub, fmt.Sprintf(msg, a...))
}

func (z *zapLogger) Print(ctx context.Context, lv Level, msg string) {

	defer z.logger.Sugar()

	switch lv {
	case Degub:
		z.logger.Debug(
			msg,
			zap.String("hostname", z.host),
			zap.String("service", z.service),
			zap.String("status", lv.String()),
		)
	case Info:
		z.logger.Info(
			msg,
			zap.String("hostname", z.host),
			zap.String("service", z.service),
			zap.String("status", lv.String()),
		)
	case Warn:
		z.logger.Warn(
			msg,
			zap.String("hostname", z.host),
			zap.String("service", z.service),
			zap.String("status", lv.String()),
		)
	case Error:
		z.logger.Error(
			msg,
			zap.String("hostname", z.host),
			zap.String("service", z.service),
			zap.String("status", lv.String()),
		)
	case Critical:
		z.logger.Fatal(
			msg,
			zap.String("hostname", z.host),
			zap.String("service", z.service),
			zap.String("status", lv.String()),
		)
	}
}
