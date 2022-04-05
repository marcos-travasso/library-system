package controllers

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitializeControllers() {
	router = gin.Default()

	initializeUserController()
	initializeBookController()
	initializeLendingController()
}

func RunRouter() {
	InitializeControllers()

	router.Run()
}
