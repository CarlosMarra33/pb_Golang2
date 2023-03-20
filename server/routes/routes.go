package routes

import (
	// "os/user"

	"application/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")
	{
		aula := main.Group("aula")
		{
			aula.POST("/criar", controllers.CreateAula)
		}
	}
	return router
}
