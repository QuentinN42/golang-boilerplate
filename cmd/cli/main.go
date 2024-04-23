package main

import (
	"github.com/QuentinN42/golang-boilerplate/pkg/logfmt"
	"github.com/QuentinN42/golang-boilerplate/pkg/telemetry"
)

func main() {
	telemetry.Init()
	logfmt.Info("Hello, World!")
}
