package main

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments"
	"github.com/go-clinic/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	config := ReadConfiguration()

	logger := CreateLogger(config)

	router := CreateGinServerInstance(logger)

	RegisterModules(router, config)

	router.Run(config.Server.Address)
}

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

func CreateGinServerInstance(logger *zap.Logger) *gin.Engine {
	router := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))

	return router
}

func RegisterModules(router gin.IRouter, configuration common.Configuration) {
	v1 := router.Group("/api")

	appointments.RegisterModule(v1.Group("/appointments"))
}
