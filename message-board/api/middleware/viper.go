package middleware

import (
	"fmt"
	"github.com/spf13/viper"
)

func ViperSetup() {
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Viper set up failed,err:%v\n", err)
		panic(err)
	}
	viper.WatchConfig()
}
