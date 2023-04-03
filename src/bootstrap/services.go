package bootstrap

import (
	"base/src/core/services"
	"go.uber.org/fx"
)

func BuildServiceModule() fx.Option {
	return fx.Options(
		fx.Provide(services.NewBaseService),
	)
}
