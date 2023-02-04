package user

import (
	"github.com/IN-45/INT20H-test-task/modules/user/internal/handler"
	"github.com/IN-45/INT20H-test-task/modules/user/internal/repository"
	"github.com/IN-45/INT20H-test-task/modules/user/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		handler.NewAuthHandler,
		service.NewAuthService,
		repository.NewUserRepository,
	),
	fx.Invoke(handler.RegisterAuthHandler),
)
