package logfmt

import (
	"github.com/QuentinN42/golang-boilerplate/pkg/telemetry"
)

var logger = telemetry.NewLogger()

var (
	Trace   = logger.Trace
	Debug   = logger.Debug
	Info    = logger.Info
	Warning = logger.Warning
	Error   = logger.Error
	Fatal   = logger.Fatal
	With    = logger.With
)
