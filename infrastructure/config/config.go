package config

import (
	"os"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var Config *Configurations

type Configurations struct {
	DB DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Url      string `mapstructure:"url"`
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

	if os.Getenv("ENVIRONMENT") == "" {
		logger.Info("ENVIRONMENT not set, defaulting to local config")
		_ = os.Setenv("ENVIRONMENT", Local)
	} else {
		logger.Info("ENVIRONMENT set to: " + os.Getenv("ENVIRONMENT"))
	}

	switch os.Getenv("ENVIRONMENT") {
	case Local, Dev:
		viper.SetConfigName("application")
	case Test:
		viper.SetConfigName("application-test")
	case Stage:
		viper.SetConfigName("application-stage")
	case Prod:
		viper.SetConfigName("application-prod")
	}

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
