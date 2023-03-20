package controllers

import (
	"application/controllers/dtos"
	"application/services"

	"github.com/gin-gonic/gin"
)

func CreateAula(c *gin.Context) {
	var request dtos.ProfessoAulaDto

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	resp, err := services.CreateNewAula(request.ProfessorID, request.Materia, request.Alunos)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi possivel criar: " + err.Error(),
		})
		return
	}
	resp = "criado com sucesso"
	c.JSON(200, resp)

}
