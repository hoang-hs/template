package constant

import "base/src/common/configs"

const (
	AppEnvDev  = "dev"
	AppEnvProd = "prod"

	TraceIdName = "trace_id"
)

func IsProdEnv() bool {
	return configs.Get().Mode == AppEnvProd
}
