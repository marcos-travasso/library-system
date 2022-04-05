package controllers

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitializeControllers() {
	router = gin.Default()

	initializeUserController()
}
