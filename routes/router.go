package routes

import "github.com/gofiber/fiber/v2"

type FiberRouter interface {
	ConfigureRoutes(app *fiber.App)
}
