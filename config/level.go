package config

import (
	"github.com/spf13/viper"
	"log"
)

var conf *viper.Viper


//Load load inital configuration
func Load(env string) (*viper.Viper, error) {
	conf = viper.New()
	conf.SetConfigType("yaml")
	conf.SetConfigName(env)
	conf.AddConfigPath("config/")

	if err := conf.ReadInConfig(); err != nil {
		log.Println("error on parsing config file", err)
		return nil, err
	}

	return conf, nil
}