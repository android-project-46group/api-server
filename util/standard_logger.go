package util

import (
	"context"
	"encoding/json"
	"fmt"
)

// standardLogger prints log to standard output stream.
type standardLogger struct {
	level   Level
	host    string
	service string
}

func NewStandardLogger(
	host string,
	service string,
) (Logger, func(), error) {

	logger := &standardLogger{
		level:   Info,
		host:    host,
		service: service,
	}
	return logger, func() {}, nil
}

func (s *standardLogger) Critical(ctx context.Context, msg string) {
	s.Print(ctx, Critical, msg)
}

func (s *standardLogger) Error(ctx context.Context, msg string) {
	s.Print(ctx, Error, msg)
}
func (s *standardLogger) Warn(ctx context.Context, msg string) {
	s.Print(ctx, Warn, msg)
}

func (s *standardLogger) Info(ctx context.Context, msg string) {
	s.Print(ctx, Info, msg)
}

func (s *standardLogger) Debug(ctx context.Context, msg string) {
	s.Print(ctx, Degub, msg)
}

func (s *standardLogger) Criticalf(ctx context.Context, msg string, a ...interface{}) {
	s.Print(ctx, Critical, fmt.Sprintf(msg, a...))
}

func (s *standardLogger) Errorf(ctx context.Context, msg string, a ...interface{}) {
	s.Print(ctx, Error, fmt.Sprintf(msg, a...))
}

func (s *standardLogger) Warnf(ctx context.Context, msg string, a ...interface{}) {
	s.Print(ctx, Warn, fmt.Sprintf(msg, a...))
}

func (s *standardLogger) Infof(ctx context.Context, msg string, a ...interface{}) {
	s.Print(ctx, Info, fmt.Sprintf(msg, a...))
}

func (s *standardLogger) Debugf(ctx context.Context, msg string, a ...interface{}) {
	s.Print(ctx, Degub, fmt.Sprintf(msg, a...))
}

func (s *standardLogger) Print(ctx context.Context, lv Level, msg string) {
	if !shouldPrint(s.level, lv) {
		return
	}

	m := logMessage{
		Host:    s.host,
		Service: s.service,
		Message: msg,
		Status:  lv.String(),
	}
	sm, err := json.Marshal(m)
	if err != nil {
		fmt.Println("{\"Error\": \"Failed to Marshal Struct to Json\"}")
	}
	sm = append(sm, []byte("\n")...)

	fmt.Println(string(sm))
}

// Set Level after struct is initialized.
// The default log level is set to Info.
func (s *standardLogger) SetLevel(lv Level) {
	s.level = lv
}
