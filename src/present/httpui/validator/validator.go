package validator

import (
	"base/src/common/log"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"reflect"
)

func NewValidator() *validator.Validate {
	return validator.New()

}

func RegisterDecimalTypeFunc(validator *validator.Validate) {
	validator.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		if valuer, ok := field.Interface().(decimal.Decimal); ok {
			return valuer.String()
		}
		return nil
	}, decimal.Decimal{})
}

func registerValidation(validator *validator.Validate, tag string, fn validator.Func) {
	if err := validator.RegisterValidation(tag, fn); err != nil {
		log.GetLogger().GetZap().Fatalf("Register custom validation %s failed with error: %s", tag, err.Error())
	}
	return
}

func RegisterValidations(validator *validator.Validate) {

}
