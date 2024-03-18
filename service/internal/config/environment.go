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
	viper.BindEnv("DATABASE_HOST")
	viper.BindEnv("DATABASE_PORT")
	viper.BindEnv("DATABASE_USER")
	viper.BindEnv("DATABASE_PASSWORD")
	viper.BindEnv("DATABASE_NAME")
}

func LoadEnvironment() (cfg EnvironmentConfig, err error) {
	// Unmarshal the configuration into the struct
	err = viper.Unmarshal(&cfg)
	return
}
