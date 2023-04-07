package bootstrap

import (
	"base/src/infra/external"
	"go.uber.org/fx"
)

func BuildExtServicesModules() fx.Option {
	return fx.Options(
		fx.Provide(external.NewBaseClient),
	)
}
