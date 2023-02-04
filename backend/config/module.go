package config

import (
	"github.com/IN-45/INT20H-test-task/pkg/db"
	"github.com/IN-45/INT20H-test-task/pkg/jwt"
	"github.com/IN-45/INT20H-test-task/pkg/server"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewConfig,
	func(cfg *Config) db.Config {
		return cfg.DB
	},
	func(cfg *Config) server.Config {
		return cfg.HTTPServer
	},
	func(cfg *Config) jwt.Config {
		return cfg.JWTConfig
	},
)
