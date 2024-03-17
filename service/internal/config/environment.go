package config

import (
	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     int    `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
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
