package controllers

import (
	"application/controllers/dtos"
	"application/database"
	"application/models"
	"application/repositorio"
	"application/services"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func MarcarPresença(c *gin.Context) {
	var request dtos.PresencaDto
	// db := database.GetDatabase()

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	fmt.Println(request)

	var p models.Presenca
	p.AlunoId = uint(request.IdAluno)
	p.AulaId = uint(request.IdAula)
	p.Tipo = "presente"
	p.DataCreate = time.Now().Day()
	p.DataUpdate = time.Now().Day()

	repositorio.MarcarPresenca(&p)

	// err = db.Create(&p).Error
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "cannot bind JSON: " + err.Error(),
	// 	})
	// 	return
	// }

	// services.MarcarPresenca(uint(request.IdAula), uint(request.IdAluno))
	fmt.Println("primeiro teste de retornada")
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "não foi possível marcar presença " + err.Error(),
	// 	})
	// 	return
	// }
	fmt.Println("testers")
	fmt.Println(err)

	c.Status(200)
}

func MarcarFalta(c *gin.Context) {
	var request dtos.PresencaDto

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	// err = services.MarcarFalta(uint(request.IdAula), uint(request.IdAluno))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi possível marcar presença " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func AtualizarPresenca(c *gin.Context) {
	var presenca dtos.PresencaDto

	err := c.ShouldBindJSON(presenca)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = services.UpdatePresenca(uint(presenca.IdAluno), presenca.Tipo)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi atualizar presença " + err.Error(),
		})
		return
	}

	c.Status(200)
}

func GetPresencaByAluno(c *gin.Context) {
	db := database.GetDatabase()
	idAluno, _ := strconv.ParseInt(c.Param("idAluno"), 10, 64)
	idAula, _ := strconv.ParseInt(c.Param("idAula"), 10, 64)
	var p []models.Presenca

	err := db.Where("aluno_id = ?", uint(idAluno)).Where("aula_id = ?", uint(idAula)).Find(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi atualizar presença " + err.Error(),
		})
		return
	}
	// presnsas := services.GetPresencaAula(uint(idAula), uint(idAluno))
	// if presnsas == nil {
	// 	c.JSON(500, gin.H{
	// 		"error": "não foi possível buscar presenças ",
	// 	})
	// }
	// fmt.Println(presnsas)
		fmt.Println(p)
	c.JSON(200, p)
}
