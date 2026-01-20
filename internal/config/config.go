package config

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

func RegisterTerminalFlags() {
	// command line parameters
	flag.String(FLAG_ENV, C_ENV_DEV, "Options: dev | test | prod")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func LoadConfig() (*ApiConfig, error) {

	// extract mode
	paramFlagEnv := viper.GetString(FLAG_ENV)

	// TOML config
	switch strings.ToLower(paramFlagEnv) {

	case C_ENV_DEV:
		viper.SetConfigName(fmt.Sprintf("%s_%s", C_CONFIG_BASE_NAME, C_ENV_DEV))

	case C_ENV_TEST:
		viper.SetConfigName(fmt.Sprintf("%s_%s", C_CONFIG_BASE_NAME, C_ENV_TEST))

	case C_ENV_PROD:
		viper.SetConfigName(C_CONFIG_BASE_NAME)

	default:
		viper.SetConfigName(fmt.Sprintf("%s_%s", C_CONFIG_BASE_NAME, C_ENV_DEV))

	}

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

func GetEnv() string {
	return viper.GetString(FLAG_ENV)
}
