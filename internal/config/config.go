package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() (*ApiConfig, error) {
	viper.SetConfigName("gouser_api_config")
	viper.SetConfigType("toml")

	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Config file error: %s", err)
	}

	var cfg ApiConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Error decoding config file: %s", err)
	}
	return &cfg, nil
}
