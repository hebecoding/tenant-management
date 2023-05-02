package config

import (
	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var Config *Configurations

type Configurations struct {
	DB DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	MongoURL      string `mapstructure:"url"`
	MongoUsername string `mapstructure:"username"`
	MongoPassword string `mapstructure:"password"`
}

func ReadInConfig(logger *utils.Logger) error {
	viper.AutomaticEnv()
	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./infrastructure/config")

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
