package routes

import "github.com/gin-gonic/gin"

type GinRouter interface {
	ConfigureRoutes(app *gin.Engine)
}
