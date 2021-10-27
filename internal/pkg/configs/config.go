package configs

import (
	"github.com/spf13/viper"
)

func Init(cfg string) error {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("conf")
	}
	return viper.ReadInConfig()
}
