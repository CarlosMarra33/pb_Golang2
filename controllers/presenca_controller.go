package controllers

import (
	"application/controllers/dtos"
	"application/services"

	"github.com/gin-gonic/gin"
)

func MarcarPresença(c *gin.Context) {
	var request dtos.PresencaDto

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = services.MarcarPresenca(request.IdAula, request.IdAluno)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi possível marcar presença " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func AtualizarPresenca(c *gin.Context){
	var presenca dtos.PresencaDto

	err := c.ShouldBindJSON(presenca)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = services.UpdatePresenca(presenca.IdAluno, presenca.Tipo)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi atualizar presença " + err.Error(),
		})
		return
	}


	c.Status(200)
}
