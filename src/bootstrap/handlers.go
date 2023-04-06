package bootstrap

import (
	"base/src/present/grpcui/handlers"
	"go.uber.org/fx"
)

func BuildHandlersModules() fx.Option {
	return fx.Options(
		fx.Provide(handlers.NewBaseHandler),
	)
}
