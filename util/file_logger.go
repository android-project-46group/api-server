package util

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type fileLogger struct {
	level   Level
	logfile *os.File
	host    string
	service string
}

type logMessage struct {
	Host    string `json:"hostname"`
	Service string `json:"service"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewFileLogger(
	path string,
	host string,
	service string,
) (Logger, func(), error) {

	logfile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, err
	}

	logger := &fileLogger{
		level:   Info,
		logfile: logfile,
		host:    host,
		service: service,
	}
	return logger, func() {
		logfile.Close()
	}, nil
}

func (f *fileLogger) Critical(ctx context.Context, msg string) {
	f.Print(ctx, Critical, msg)
}

func (f *fileLogger) Error(ctx context.Context, msg string) {
	f.Print(ctx, Error, msg)
}
func (f *fileLogger) Warn(ctx context.Context, msg string) {
	f.Print(ctx, Warn, msg)
}

func (f *fileLogger) Info(ctx context.Context, msg string) {
	f.Print(ctx, Info, msg)
}

func (f *fileLogger) Debug(ctx context.Context, msg string) {
	f.Print(ctx, Degub, msg)
}

func (f *fileLogger) Criticalf(ctx context.Context, msg string, a ...interface{}) {
	f.Print(ctx, Critical, fmt.Sprintf(msg, a...))
}

func (f *fileLogger) Errorf(ctx context.Context, msg string, a ...interface{}) {
	f.Print(ctx, Error, fmt.Sprintf(msg, a...))
}

func (f *fileLogger) Warnf(ctx context.Context, msg string, a ...interface{}) {
	f.Print(ctx, Warn, fmt.Sprintf(msg, a...))
}

func (f *fileLogger) Infof(ctx context.Context, msg string, a ...interface{}) {
	f.Print(ctx, Info, fmt.Sprintf(msg, a...))
}

func (f *fileLogger) Debugf(ctx context.Context, msg string, a ...interface{}) {
	f.Print(ctx, Degub, fmt.Sprintf(msg, a...))
}

func (f *fileLogger) Print(ctx context.Context, lv Level, msg string) {
	if !shouldPrint(f.level, lv) {
		return
	}

	m := logMessage{
		Host:    f.host,
		Service: f.service,
		Message: msg,
		Status:  lv.String(),
	}
	s, err := json.Marshal(m)
	if err != nil {
		f.logfile.WriteString("{\"Error\": \"Failed to Marshal Struct to Json\"}")
	}
	s = append(s, []byte("\n")...)

	f.logfile.Write(s)
}

// Set Level after struct is initialized.
// The default log level is set to Info.
func (f *fileLogger) SetLevel(lv Level) {
	f.level = lv
}
