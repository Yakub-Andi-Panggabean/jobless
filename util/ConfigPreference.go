package util

import (
	"github.com/spf13/viper"
)

var viperConfig *viper.Viper

func init() {

	viperConfig=viper.New()
	viperConfig.SetConfigType("toml")
	viperConfig.SetConfigFile("./config/application.toml")
	//viperConfig.AddConfigPath("./config")
	err := viperConfig.ReadInConfig()
	if err != nil {
		panic("an error occurred when trying to read config file, error message :" + err.Error())
	}

}

func GetConfig() *viper.Viper {

	return viperConfig

}

