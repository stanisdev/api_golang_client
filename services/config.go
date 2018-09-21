package services

import (
	"github.com/spf13/viper"
	"fmt"
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./src/app")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}	
}