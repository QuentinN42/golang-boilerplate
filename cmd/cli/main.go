package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/QuentinN42/golang-boilerplate/pkg/sdk/env"
	"github.com/QuentinN42/golang-boilerplate/pkg/sdk/functool"
	"github.com/QuentinN42/golang-boilerplate/pkg/sdk/logfmt"
	"github.com/QuentinN42/golang-boilerplate/pkg/sdk/telemetry"
)

type Envs struct {
}

func main() {
	telemetry.Init()
	envs, err := env.GetAll(&Envs{})
	if err != nil {
		logfmt.Fatal("Unable to set up environment", err)
	}

	logfmt.Trace("Env: %v", envs)

	lst := functool.Collect(
		functool.Filter(
			functool.Apply(
				functool.Apply(
					functool.Iter(context.Background(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
					func(elm int) int { return 3 * elm },
				),
				func(elm int) string { return fmt.Sprint(elm) },
			),
			func(elm string) bool { return strings.Contains(elm, "1") },
		),
	)
	logfmt.Trace("List: %v", lst)
}
