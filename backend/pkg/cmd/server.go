package cmd

import (
	"github.com/IN-45/INT20H-test-task/pkg/server"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

func NewServerStartCommand(fxOptions []fx.Option) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "start HTTP server",
		Action: func(c *cli.Context) error {
			fxOptions = append(
				fxOptions,
				fx.Invoke(
					server.RunServer,
				),
			)
			app := fx.New(fxOptions...)

			app.Run()

			return nil
		},
	}
}
