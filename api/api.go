package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-travasso/library-system/api/database"
	"log"
	"os"
)

var dbDir = database.Database{Dir: "./database/database.db"}
var router *gin.Engine

func main() {
	f, err := os.OpenFile("logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	log.SetOutput(f)

	dbDir.CreateDatabase()

	setupRoutes()
}

func setupRoutes() {
	router = gin.Default()

	setupUserFunctions()

	_ = router.Run(":8080")
}
