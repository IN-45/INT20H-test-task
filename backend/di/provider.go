package di

import (
	"github.com/IN-45/INT20H-test-task/config"
	"github.com/IN-45/INT20H-test-task/docs"
	"github.com/IN-45/INT20H-test-task/modules/storage"
	"github.com/IN-45/INT20H-test-task/modules/user"
	"github.com/IN-45/INT20H-test-task/pkg/db"
	"github.com/IN-45/INT20H-test-task/pkg/jwt"
	"github.com/IN-45/INT20H-test-task/pkg/server"
	"github.com/IN-45/INT20H-test-task/pkg/validator"
	"go.uber.org/fx"
)

func ProvideModules() []fx.Option {
	modules := []fx.Option{
		server.Module,
		config.Module,
		db.Module,
		validator.Module,
		jwt.Module,
		user.Module,
		docs.Module,
		storage.Module,
	}

	return modules
}
