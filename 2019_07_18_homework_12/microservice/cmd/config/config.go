package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     int    `yaml:"port"`
	LogLevel uint32 `yaml:"logLevel"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func init() {
	viper.SetDefault("MICRO_CONFIG", "config.yaml")
	loadConfig()
}

func loadConfig() {
	cfgPath := viper.GetString("MICRO_CONFIG")
	viper.SetConfigFile(cfgPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
}
