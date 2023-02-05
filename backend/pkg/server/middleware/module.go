package middleware

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewTokenValidatorMiddleware,
	),
	fx.Invoke(func(
		server *fiber.App,
		validatorMiddleware *tokenValidatorMiddleware,
	) {
		server.Use(
			cors.New(),
			logger.New(),
			skip.New(validatorMiddleware.validateToken, func(ctx *fiber.Ctx) bool {
				path := ctx.Path()
				isSwaggerPath, _ := regexp.MatchString("/swagger/([a-z]*)", path)
				isAuthPath, _ := regexp.MatchString("/sign-([a-z]*)", path)

				return isAuthPath || isSwaggerPath
			}),
		)
	}),
)
