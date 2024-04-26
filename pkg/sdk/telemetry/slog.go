package telemetry

import (
	"context"
	"os"

	"github.com/sagikazarmark/slog-shim"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelDebug   = slog.LevelDebug
	LevelInfo    = slog.LevelInfo
	LevelWarning = slog.LevelWarn
	LevelError   = slog.LevelError
	LevelFatal   = slog.Level(12)
)

type handler struct {
	slog.Handler
}

// Automatically add trace_id to log messages
func (h handler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(traceIDKey{}).(string); ok {
		r.Add("trace_id", slog.StringValue(traceID))
	}

	return h.Handler.Handle(ctx, r)
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key != slog.LevelKey {
		return a
	}

	level := a.Value.Any().(slog.Level)
	switch {
	case level < LevelDebug:
		a.Value = slog.StringValue("trace")
	case level < LevelInfo:
		a.Value = slog.StringValue("debug")
	case level < LevelWarning:
		a.Value = slog.StringValue("info")
	case level < LevelError:
		a.Value = slog.StringValue("warning")
	case level < LevelFatal:
		a.Value = slog.StringValue("error")
	default:
		a.Value = slog.StringValue("fatal")
	}
	return a
}

func newLogger() *slog.Logger {
	return slog.New(
		handler{
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level:       LevelTrace,
					ReplaceAttr: replaceAttr,
				},
			),
		},
	)
}
