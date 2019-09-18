package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Logger Logger `yaml:"logger"`
	}

	Logger struct {
		Level  string `yaml:"level"`
		Output string `yaml:"output"`
	}
)

func LoadConfig() (*Config, error) {
	var config *Config
	viper.SetDefault("CALENDAR_CONFIG", "config.yaml")
	viper.SetConfigFile(viper.GetString("CALENDAR_CONFIG"))
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read config with filepath")
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config to struct")
	}

	return config, nil
}
