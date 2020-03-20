package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.Run()
}