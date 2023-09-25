package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewConfig() *viper.Viper {
	log.Println("loading conf file")
	appEnv := os.Getenv("APP_ENV")
	configFile := "application-prod.yml"
	if appEnv == "" || appEnv == "dev" {
		configFile = "application-dev.yml"
	}
	return getConfig(fmt.Sprintf("config/%s", configFile))
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
