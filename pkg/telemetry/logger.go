package telemetry

import (
	"context"
	"fmt"

	"github.com/sagikazarmark/slog-shim"
)

type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warning(msg string, err error)
	Error(msg string, err error)
	Fatal(msg string, err error)
	With(ctx context.Context) Logger
}

func fmtMsgErr(msg string, err error) string {
	if err == nil {
		return msg
	}
	return msg + ": '" + err.Error() + "'"
}

type ctxLogger struct {
	Ctx    context.Context
	Logger *slog.Logger
}

func (l *ctxLogger) log(level slog.Level, msg string, args ...interface{}) {
	l.Logger.Log(l.Ctx, level, fmt.Sprintf(msg, args...))
}

func (l *ctxLogger) logErr(level slog.Level, msg string, err error) {
	l.Logger.Log(l.Ctx, level, fmtMsgErr(msg, err))
}

func (l *ctxLogger) Trace(msg string, args ...interface{}) {
	l.log(LevelTrace, msg, args...)
}
func (l *ctxLogger) Debug(msg string, args ...interface{}) {
	l.log(LevelDebug, msg, args...)
}
func (l *ctxLogger) Info(msg string, args ...interface{}) {
	l.log(LevelInfo, msg, args...)
}

func (l *ctxLogger) Warning(msg string, err error) {
	l.logErr(LevelWarning, msg, err)
}
func (l *ctxLogger) Error(msg string, err error) {
	l.logErr(LevelError, msg, err)
}
func (l *ctxLogger) Fatal(msg string, err error) {
	l.logErr(LevelFatal, msg, err)
	panic(err)
}

func (l *ctxLogger) With(ctx context.Context) Logger {
	return &ctxLogger{
		Ctx:    ctx,
		Logger: l.Logger,
	}
}

func NewLogger() Logger {
	return &ctxLogger{
		Ctx:    context.Background(),
		Logger: newLogger(),
	}
}
