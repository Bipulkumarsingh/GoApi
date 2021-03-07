package main

import (
	"github.com/gin-gonic/gin"
)

//gin engine
var router *gin.Engine

func main() {
	// Default middleware
	router = gin.Default()

	// routes defined in routes.go
	initializeRoutes()

	// Running server
	router.Run()
}
