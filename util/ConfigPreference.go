package util

import (
	"github.com/spf13/viper"
)

var viperConfig viper.Viper

func init() {

	viperConfig=viper.New()
	viperConfig.SetConfigFile("application")
	viperConfig.AddConfigPath("config/")
	err := viperConfig.ReadInConfig()
	if err != nil {
		panic("an error occurred when trying to read config file, error message :" + err)
	}

}

func GetConfig() viper.Viper {

	return viperConfig

}

