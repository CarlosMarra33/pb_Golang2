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

		present := main.Group("presenca")
		{
			present.POST("/presente", controllers.MarcarPresença)
			present.POST("/falta", controllers.MarcarFalta)
			present.PUT("atualza", controllers.AtualizarPresenca)
		}
	}
	return router
}
