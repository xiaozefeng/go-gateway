package configs

import (
	"github.com/spf13/viper"
)

func InitializeConfig(cfg string) error {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("configs")
	}
	return viper.ReadInConfig()
}
