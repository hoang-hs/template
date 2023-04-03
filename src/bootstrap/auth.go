package bootstrap

import (
	"base/src/present/httpui/middlewares"
	"go.uber.org/fx"
)

func BuildAuthModules() fx.Option {
	return fx.Options(
		fx.Provide(middlewares.NewAuthService),
	)
}
