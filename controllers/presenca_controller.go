package controllers

import (
	"application/controllers/dtos"
	"application/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PresencardController struct {
	presencaService services.PresencaService
}

func NewPresencaController(presencaService services.PresencaService) *PresencardController {
	return &PresencardController{presencaService: presencaService}
}

func (prs *PresencardController) MarcarPresença(c *gin.Context) {
	var request dtos.PresencaDto
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	resposta := prs.presencaService.MarcarPresenca(uint(request.IdAula), uint(request.IdAluno))
	if resposta != "" {
		c.JSON(http.StatusFound, gin.H{
			"msg": "resposta ",
		})
		return
	}
	c.Status(200)
}

func (prs *PresencardController) MarcarFalta(c *gin.Context) {
	var request dtos.PresencaDto
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = prs.presencaService.MarcarFalta(uint(request.IdAula), uint(request.IdAluno))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi possível marcar presença " + err.Error(),
		})
		return
	}
	c.Status(204)
}

func (prs *PresencardController) AtualizarPresenca(c *gin.Context) {
	var presenca dtos.PresencaDto

	err := c.ShouldBindJSON(&presenca)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = prs.presencaService.UpdatePresenca(uint(presenca.IdPresenca), presenca.Tipo)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "não foi atualizar presença " + err.Error(),
		})
		fmt.Println("não foi possivel criar: ", err.Error())
		return
	}
	c.Status(200)
}

func (prs *PresencardController) GetPresencaByAluno(c *gin.Context) {
	idAluno, _ := strconv.ParseInt(c.Param("idAluno"), 10, 64)
	idAula, _ := strconv.ParseInt(c.Param("idAula"), 10, 64)
	presenca := prs.presencaService.GetPresencaAula(uint(idAula), uint(idAluno))
	fmt.Println(presenca)
	c.JSON(200, presenca)
}
