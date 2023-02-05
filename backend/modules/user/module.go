package user

import (
	user_handler "github.com/IN-45/INT20H-test-task/modules/user/internal/handler"
	user_repository "github.com/IN-45/INT20H-test-task/modules/user/internal/repository"
	user_service "github.com/IN-45/INT20H-test-task/modules/user/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		user_handler.NewAuthHandler,
		user_service.NewAuthService,
		user_repository.NewUserRepository,
	),
	fx.Invoke(user_handler.RegisterAuthHandler),
)
