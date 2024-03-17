package configs

import (
	"github.com/spf13/viper"
)

type EnvironmentConfig struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

func LoadEnvironment() (config EnvironmentConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
