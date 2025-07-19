package main

import (
	"GinWithGormTutorial/config"
	"GinWithGormTutorial/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}

	_ = r.Run(port)
}
