package repositorio

import (
	"application/database"
	"application/models"
	"time"
)

func ChecaPresenca(idAula uint, idAluno uint) bool {
	db := database.GetDatabase()
	var p []models.Presenca
	err := db.Where("AlunoId = ?", idAluno).Where("AulaId = ?", idAula).Find(&p).Error
	if err != nil {
		return false
	}

	for _, p := range p {
		if p.DataCreate.Day() == time.Now().Day() {
			return false
		}
	}

	return true
}

func MarcarPresenca(presenca models.Presenca) error {
	db := database.GetDatabase()

	err := db.Create(presenca).Error
	if err != nil {
		return err
	}
	return nil

}

func AtualizarPresenca(tipo string, presencaOldId uint) error {
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
