package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-clinic/appointments"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	RegisterModules(router)

	router.Run()
}

func RegisterModules(router gin.IRouter) {
	v1 := router.Group("/api")

	appointments.RegisterModule(v1.Group("/appointments"))
}
