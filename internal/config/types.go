package config

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
}
