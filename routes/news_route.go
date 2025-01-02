package routes

import (
	"github.com/gin-gonic/gin"
	_ "go-gin-template/domain"
)

type NewsRouter struct {
}

func NewNewsRouter() *NewsRouter {
	return &NewsRouter{}
}

func (r *NewsRouter) ConfigureRoutes(app *gin.Engine) {
	routes := app.Group("/api/v1/news")

	routes.GET("/", ListNews)
	routes.GET("/:newsId", GetNewsById)
}

func ListNews(c *gin.Context) {
	c.String(200, "List news")
}

func GetNewsById(c *gin.Context) {
	newsId := c.Param("newsId")

	c.String(200, "Get news by id: "+newsId)
}
