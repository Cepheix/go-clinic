package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments"
	"github.com/go-clinic/common"
	"github.com/spf13/viper"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	config := ReadConfiguration()

	RegisterModules(router, config)

	router.Run()
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

func RegisterModules(router gin.IRouter, configuration common.Configuration) {
	v1 := router.Group("/api")

	appointments.RegisterModule(v1.Group("/appointments"))
}
