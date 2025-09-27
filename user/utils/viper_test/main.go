package main

import (
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return
	}
}
