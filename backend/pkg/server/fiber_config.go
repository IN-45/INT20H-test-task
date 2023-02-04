package server

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Addr        string `env:"SERVER_ADDR" envDefault:"0.0.0.0:8000"`
	ReadTimeout string `env:"SERVER_READ_TIMEOUT" envDefault:"10"`
}

func FiberConfig(cfg Config) fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(cfg.ReadTimeout)

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
