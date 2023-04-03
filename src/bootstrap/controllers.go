package bootstrap

import (
	"base/src/present/httpui/controllers"
	"base/src/present/httpui/validator"
	"go.uber.org/fx"
)

func BuildControllerModule() fx.Option {
	return fx.Options(
		fx.Provide(controllers.NewBaseController),
	)
}

func BuildValidator() fx.Option {
	return fx.Options(
		fx.Provide(validator.NewValidator),
		fx.Invoke(validator.RegisterDecimalTypeFunc),
		fx.Invoke(validator.RegisterValidations),
	)
}
