package main

import (
	"log"
	"rest-app-pos/src/api"
	"rest-app-pos/src/config"
	"rest-app-pos/src/containers"
)

func main() {
	// Initialize Dependency Injection
	app := containers.InitAppDependencies()

	// Setup Router
	r := api.SetupRouter(app)
	port := config.AppConfig.Server.Port

	log.Println("Server running on port", port)

	r.SetTrustedProxies(nil)
	r.Run(":" + port)
}