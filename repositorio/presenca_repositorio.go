package repositorio

import (
	"application/database"
	"application/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PresencaRepository struct {
	db *gorm.DB
}

func NewPresencaRepository() *PresencaRepository {
	return &PresencaRepository{
		db: database.GetDatabase(),
	}
}

func (pr *PresencaRepository) GetPresencaAula(idAula *uint, idAluno *uint) []models.Presenca {
	var p []models.Presenca
	pre := models.Presenca{
		AlunoId: *idAluno,
		AulaId:  *idAula,
	}
	err := pr.db.Where("aluno_id = ?", &pre.AlunoId).Where("aula_id = ?", &pre.AulaId).Find(&p).Error
	if err != nil {
		return nil
	}
	fmt.Println(p, "        present")
	return p
}

func (pr *PresencaRepository) ChecaPresenca(idAula *uint, idAluno *uint) bool {
	var p []models.Presenca
	err := pr.db.Where("aluno_id = ?", &idAluno).Where("aula_id = ?", &idAula).Find(&p).Error
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
	fmt.Println("teste banco")
	fmt.Println(presenca)
	err := pr.db.Create(presenca).Error
	if err != nil {
		return
	}
}

func (pr *PresencaRepository) AtualizarPresenca(tipo string, presencaOldId uint) error {
	var presenca models.Presenca
	err := pr.db.First(&presenca, presencaOldId).Error
	if err != nil {
		return err
	}
	presenca.Tipo = tipo
	presenca.DataUpdate = time.Now()
	err = pr.db.Save(&presenca).Error
	if err != nil {
		return err
	}
	return nil
}
