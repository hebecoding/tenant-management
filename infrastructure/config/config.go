package config

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var Config *Configurations

type Configurations struct {
	Environment string         `mapstructure:"environment"`
	Application Application    `mapstructure:"application"`
	DB          DatabaseConfig `mapstructure:"database"`
}

type Application struct {
	Name    string `mapstructure:"name"`
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type DatabaseConfig struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

const (
	Local = "local"
	Dev   = "dev"
	Test  = "test"
	Stage = "stage"
	Prod  = "prod"
)

func ReadInConfig(logger *utils.Logger) error {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("infrastructure/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath(".././../config")

	logger.Info("Reading in config file")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err)
		return errors.Wrap(err, "failed to read in config file")
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.Error(err)
		return errors.Wrap(err, "failed to unmarshal config file")
	}

	logger.Info("Successfully read in config file")

	return nil
}
