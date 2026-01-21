package config

import ()

const (
	C_ENV_DEBUG string = "debug"
	C_ENV_DEV   string = "dev"
	C_ENV_TEST  string = "test"
	C_ENV_PROD  string = "prod"
)

const (
	FLAG_ENV = "env"
)

const (
	C_CONFIG_BASE_NAME = "gouser_api_config"
)

type ApiConfig struct {
	Logger struct {
		LogLevel string `mapstructure:"log_level"`
	} `mapstructure:"logger"`

	GouserApi struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"gouser-api"`

	ExtRandomuser struct {
		Host  string `mapstructure:"host"`
		Route string `mapstructure:"route"`
	} `mapstructure:"ext-randomuser"`

	Profiler struct {
		EnablePprof bool `mapstructure:"enable_pprof"`
	} `mapstructure:"profiler"`
}
