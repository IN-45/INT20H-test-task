package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewHTTPServer(cfg Config) *fiber.App {
	app := fiber.New(FiberConfig(cfg))

	return app
}

func RunServer(lc fx.Lifecycle, cfg Config, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Printf("Starting HTTP server at %s\n", cfg.Addr)
			go app.Listen(cfg.Addr) //nolint:errcheck

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
