package main

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-fiber-template/docs"
	"go-fiber-template/helpers"
)

// @title Go Gin Template
// @version 1.0
// @description Go Gin Template
// @license MIT
// @BasePath /
func main() {
	app, err := InitializeApp()

	if err != nil {
		panic(err)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // default

	app.Use(helpers.RecoveryMiddleware())

	// configure routes
	routers, err := InitialGinRouters()

	if err != nil {
		panic(err)
	}

	for _, router := range routers {
		router.ConfigureRoutes(app)
	}

	println("Server is running on port 8080")
	err = app.Run(":8080")

	if err != nil {
		panic(err)
	}

	println("Server is stopped")
}
