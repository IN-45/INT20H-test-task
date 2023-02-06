package config

import (
	"github.com/IN-45/INT20H-test-task/pkg/cloudstorage"
	"github.com/IN-45/INT20H-test-task/pkg/db"
	"github.com/IN-45/INT20H-test-task/pkg/jwt"
	"github.com/IN-45/INT20H-test-task/pkg/server"
	"github.com/caarlos0/env/v7"
)

type Config struct {
	DB                 db.Config
	HTTPServer         server.Config
	JWTConfig          jwt.Config
	CloudStorageConfig cloudstorage.Config
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
