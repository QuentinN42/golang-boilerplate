package main

import (
	"github.com/QuentinN42/golang-boilerplate/pkg/env"
	"github.com/QuentinN42/golang-boilerplate/pkg/logfmt"
	"github.com/QuentinN42/golang-boilerplate/pkg/telemetry"
)

type Envs struct {
	OTEL_SERVICE_NAME string
}

func main() {
	telemetry.Init()
	envs, err := env.GetAll(&Envs{})
	if err != nil {
		logfmt.Fatal("Unable to set up environment", err)
	}

	logfmt.Trace("Env: %v", envs)
}
