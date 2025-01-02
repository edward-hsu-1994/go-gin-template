package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	docs "go-fiber-template/docs"
	"go-fiber-template/helpers"
	"strings"
)

// @title Go Fiber Template
// @version 1.0
// @description Go Fiber Template
// @license MIT
// @BasePath /
func main() {
	app, err := InitializeApp()

	if err != nil {
		panic(err)
	}

	// Setting hostname for swagger
	alreadySettingSwaggerHostname := false
	app.Use(func(c *fiber.Ctx) error {
		if alreadySettingSwaggerHostname {
			return c.Next()
		}

		if strings.HasPrefix(c.Path(), "/swagger/") == false {
			return c.Next()
		}

		alreadySettingSwaggerHostname = true

		docs.SwaggerInfo.Host = c.GetRespHeader("X-Forwarded-For", c.Hostname())

		return c.Next()
	})
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return helpers.ErrorResponse(c, err)
		}
		return err
	})

	// configure routes
	routers, err := InitialFiberRouters()

	if err != nil {
		panic(err)
	}

	for _, router := range routers {
		router.ConfigureRoutes(app)
	}

	println("Server is running on port 3000")
	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}

	println("Server is stopped")
}
