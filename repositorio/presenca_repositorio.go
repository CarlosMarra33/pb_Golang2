package repositorio

import (
	"application/database"
	"application/models"
	"fmt"
	"time"
)

type PresencaRepository struct {}



func (pr *PresencaRepository) GetPresencaAula(idAula *uint, idAluno *uint) []models.Presenca {
	db := database.GetDatabase()
	var p []models.Presenca
	pre := models.Presenca{
		AlunoId: *idAluno,
		AulaId:  *idAula,
	}
	err := db.Where("aluno_id = ?", &pre.AlunoId).Where("aula_id = ?", &pre.AulaId).Find(&p).Error
	if err != nil {
		return nil
	}
	fmt.Println(p, "        present")
	return p
}

func (pr *PresencaRepository) ChecaPresenca(idAula *uint, idAluno *uint) bool {
	db := database.GetDatabase()
	var p []models.Presenca
	err := db.Where("aluno_id = ?", &idAluno).Where("aula_id = ?", &idAula).Find(&p).Error
	if err != nil {
		return false
	}
	for _, p := range p {
		fmt.Println("teste de data:   ", p.DataCreate.Day())
		if p.DataCreate.Day() == time.Now().Day() {
			return false
		}
	}
	return true
}

func (pr *PresencaRepository) MarcarPresenca(presenca *models.Presenca) {
	db := database.GetDatabase()
	fmt.Println("teste banco")
	fmt.Println(presenca)
	err := db.Create(presenca).Error
	if err != nil {
		return
	}
}

func (pr *PresencaRepository) AtualizarPresenca(tipo string, presencaOldId uint) error {
	db := database.GetDatabase()
	var presenca models.Presenca
	err := db.First(&presenca, presencaOldId).Error
	if err != nil {
		return err
	}
	presenca.Tipo = tipo
	presenca.DataUpdate = time.Now()
	err = db.Save(&presenca).Error
	if err != nil {
		return err
	}
	return nil
}
