package main

import (
	"github.com/go-clinic/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ReadConfiguration() common.Configuration {
	var config common.Configuration

	viper := viper.New()
	viper.AddConfigPath(".")
	viper.SetConfigFile("settings.json")
	viper.ReadInConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func CreateLogger(configuration common.Configuration) *zap.Logger {
	var loggingConfiguration zap.Config

	if configuration.Logging.Development {
		loggingConfiguration = zap.NewDevelopmentConfig()
	} else {
		loggingConfiguration = zap.NewProductionConfig()
	}

	logger, err := loggingConfiguration.Build()

	if err != nil {
		panic(err.Error())
	}

	return logger
}
