package routes

import (
	// "os/user"

	"application/controllers"
	"application/repositorio"
	"application/services"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	presencaController := controllers.NewPresencaController(*services.NewPresencaService(*repositorio.NewPresencaRepository()))
	aulaController := controllers.NewAulaController(*services.NewAulaService(*repositorio.NewAulaRepository()))
	main := router.Group("/api")
	{
		aula := main.Group("/aula")
		{
			aula.POST("/criar", aulaController.CreateAula)

		}

		present := main.Group("presenca")
		{
			present.POST("/presente", presencaController.MarcarPresença)
			present.POST("/falta", presencaController.MarcarFalta)
			present.PUT("/atualizar", presencaController.AtualizarPresenca)
			present.GET("/getPresenca/:idAluno/:idAula", presencaController.GetPresencaByAluno)
		}
	}
	return router
}
