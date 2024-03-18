package config

import (
	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	DatabaseDSN string `mapstructure:"DATABASE_DSN"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

func init() {
	// Bind the environment variables to the configuration
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
}

func LoadEnvironment() (config EnvironmentConfig, err error) {
	// Read the command line flags
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshal the configuration into the struct
	err = viper.Unmarshal(&config)
	return
}
