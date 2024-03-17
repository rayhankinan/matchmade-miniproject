package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type FlagConfig struct {
	Port int `mapstructure:"port"`
}

func init() {
	// Define the flags for the application
	pflag.Int("port", 8080, "Port to run the application on")

	// It is important to bind immediately after defining each flag
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()
}

func LoadFlags() (config FlagConfig, err error) {
	// Unmarshal the configuration into the struct
	err = viper.Unmarshal(&config)
	return
}
