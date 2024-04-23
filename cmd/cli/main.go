package main

import (
	"github.com/QuentinN42/golang-boilerplate/pkg/env"
	"github.com/QuentinN42/golang-boilerplate/pkg/logfmt"
	"github.com/QuentinN42/golang-boilerplate/pkg/telemetry"
)

type Envs struct {
	MANDATORY_STR string
	OPTIONAL_STR  *string

	MANDATORY_INT int
	OPTIONAL_INT  *int

	OPTIONAL_BOOL1 bool // "" => false
	OPTIONAL_BOOL2 *bool
}

func main() {
	telemetry.Init()
	logfmt.Info("Hello, World!")
	envs, err := env.GetAll(&Envs{})
	if err != nil {
		logfmt.Fatal("Failed to get environment variables", err)
	}

	logfmt.Info("MANDATORY_STR: %v", envs.MANDATORY_STR)
	logfmt.Info("OPTIONAL_STR: %v", *envs.OPTIONAL_STR)
	logfmt.Info("MANDATORY_INT: %v", envs.MANDATORY_INT)
	logfmt.Info("OPTIONAL_INT: %v", *envs.OPTIONAL_INT)
	logfmt.Info("OPTIONAL_BOOL1: %v", envs.OPTIONAL_BOOL1)
	logfmt.Info("OPTIONAL_BOOL2: %v", *envs.OPTIONAL_BOOL2)
}
