package main

import (
	"github.com/marcos-travasso/library-system/controllers"
	"github.com/marcos-travasso/library-system/services"
)

func main() {
	services.InitializeServices()
	controllers.RunRouter()
}
