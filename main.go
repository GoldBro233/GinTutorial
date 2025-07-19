package main

import (
	"GinTutorial/config"
	"GinTutorial/router"
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
