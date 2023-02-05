package server

import (
	"github.com/IN-45/INT20H-test-task/pkg/server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHTTPServer),
	middleware.Module,
)
