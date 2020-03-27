package main

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments"
	"github.com/go-clinic/common"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

func CreateGinServerInstance(logger *zap.Logger) *gin.Engine {
	router := gin.New()

	// Add a ginzap middleware, which Logs all requests
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	router.Use(ginzap.RecoveryWithZap(logger, true))

	return router
}

func RegisterModules(router gin.IRouter, configuration common.Configuration) {
	AddSwaggerMiddleware(router, configuration)

	api := router.Group("/api")

	appointments.RegisterModule(api.Group("/appointments"))
}

func AddSwaggerMiddleware(router gin.IRouter, configuration common.Configuration) {
	url := ginSwagger.URL(configuration.Server.SwaggerDocumentAddress()) // The url pointing to API definition
	swaggerMatch := configuration.Server.SwaggerPrefix + "/*any"
	router.GET(swaggerMatch, ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
