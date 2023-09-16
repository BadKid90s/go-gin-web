package config

import (
	"github.com/spf13/viper"
	"log"
)

func NewConfig() *viper.Viper {
	log.Println("loading conf file")
	return getConfig("config/application.yml")
}
func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
