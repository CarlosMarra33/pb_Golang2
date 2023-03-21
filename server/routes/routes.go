package routes

import (
	// "os/user"

	"application/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")
	{
		aula := main.Group("/aula")
		{
			aula.POST("/criar", controllers.CreateAula)
			aula.POST("/presente", controllers.MarcarPresença)
		}

		present := main.Group("presenca")
		{
			// present.
			present.POST("/falta", controllers.MarcarFalta)
			present.PUT("/atualza", controllers.AtualizarPresenca)
			present.GET("/getPresenca/:idAluno/:idAula", controllers.GetPresencaByAluno)
		}
	}
	return router
}
