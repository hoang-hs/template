package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mode string `mapstructure:"mode"`

	// services
	Server struct {
		Name string `mapstructure:"name"`
		Http struct {
			Address string `mapstructure:"address"`
			Prefix  string `mapstructure:"prefix"`
		} `mapstructure:"httpui"`
		Grpc struct {
			Address string `mapstructure:"address"`
		} `mapstructure:"grpc"`
	} `mapstructure:"server"`

	Kafka struct {
		Host     string `mapstructure:"host"`
		Producer struct {
		} `mapstructure:"producer"`
	} `mapstructure:"kafka"`

	//swagger
	Swagger struct {
		Enabled bool `mapstructure:"enabled"`
	} `mapstructure:"swagger"`

	//tracer
	Tracer struct {
		Enabled bool `mapstructure:"enabled"`
		Jaeger  struct {
			Endpoint string `mapstructure:"endpoint"`
			Active   bool   `mapstructure:"active"`
		} `mapstructure:"jaeger"`
	} `mapstructure:"tracer"`

	GrpcServer struct {
	} `mapstructure:"grpc_server"`

	ApiKey map[string]string `mapstructure:"api_key"`
}

var common *Config

func Get() *Config {
	return common
}

func LoadConfig(pathConfig string) error {
	viper.SetConfigFile(pathConfig)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&common)

	return nil
}
