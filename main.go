package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "go-fiber-template/docs"
	"go-fiber-template/helpers"
	"strings"
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

	// Setting hostname for swagger
	alreadySettingSwaggerHostname := false
	app.Use(func(c *gin.Context) {
		if alreadySettingSwaggerHostname {
			c.Next()
			return
		}

		if strings.HasPrefix(c.FullPath(), "/swagger/") == false {
			c.Next()
			return
		}

		alreadySettingSwaggerHostname = true

		hostname := c.GetHeader("X-Forwarded-Host")

		if hostname == "" {
			hostname = c.Request.Host
		}

		docs.SwaggerInfo.Host = hostname

		c.Next()
	})
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
