package repositorio

import (
	"application/database"
	"application/models"
	"fmt"
	"time"
)

func GetPresencaAula(idAula *uint, idAluno *uint) []models.Presenca {
	db := database.GetDatabase()
	var p []models.Presenca

	err := db.Where("aluno_id = ?", &idAluno).Where("aula_id = ?", &idAula).Find(&p).Error
	fmt.Println("erro do banco de dados      : ",err.Error)
	if err != nil {
		return nil
	}
	fmt.Println(p, "        present")
	return p
}

func ChecaPresenca(idAula uint, idAluno uint) bool {
	db := database.GetDatabase()
	var p []models.Presenca
	err := db.Where("AlunoId = ?", idAluno).Where("AulaId = ?", idAula).Find(&p).Error
	if err != nil {
		return false
	}

	for _, p := range p {
		if p.DataCreate == time.Now().Day() {
			return false
		}
	}

	return true
}

func MarcarPresenca(presenca *models.Presenca) {
	db := database.GetDatabase()

	fmt.Println("teste banco")

	fmt.Println(presenca)

	err := db.Create(presenca).Error
	if err != nil {
		return
	}

}

func AtualizarPresenca(tipo string, presencaOldId uint) error {
	db := database.GetDatabase()
	var presenca models.Presenca

	err := db.First(&presenca, presencaOldId).Error
	if err != nil {
		return err
	}
	presenca.Tipo = tipo
	presenca.DataUpdate = time.Now().Day()
	err = db.Save(&presenca).Error
	if err != nil {
		return err
	}
	return nil

}
