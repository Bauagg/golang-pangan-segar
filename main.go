package main

import (
	"log"
	"pangan-segar/config"
	"pangan-segar/databases"
	"pangan-segar/middleware"
	"pangan-segar/migration"
	"pangan-segar/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.IntConfigEnv()

	// Connection databases
	databases.Connect()
	migration.Migration()

	app := gin.Default()

	// midelware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.ErrorMiddleware())

	// Setup static file serving for images
	app.Static("/profile-user", "./public/profile-user")

	// Router
	router.RouterGlobal(app)
	router.RouterKonsumen(app)

	err := app.Run(config.APP_PORT)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
