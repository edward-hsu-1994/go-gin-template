package routes

import (
	"github.com/gofiber/fiber/v2"
	_ "go-fiber-template/domain"
)

type NewsRouter struct {
}

func NewNewsRouter() *NewsRouter {
	return &NewsRouter{}
}

func (r *NewsRouter) ConfigureRoutes(app *fiber.App) {
	routes := app.Group("/api/v1/news")

	routes.Get("/", ListNews)
	routes.Get("/:newsId", GetNewsById)
}

func ListNews(c *fiber.Ctx) error {
	return c.SendString("List news")
}

func GetNewsById(c *fiber.Ctx) error {
	newsId := c.Params("newsId")

	return c.SendString("Get news by id: " + newsId)
}
