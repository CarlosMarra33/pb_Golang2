package repositorio

import (
	"application/database"
	"application/models"

	"gorm.io/gorm"
)

type AulaRepository struct {
	db *gorm.DB
}

func NewAulaRepository() *AulaRepository {
	return &AulaRepository{
		db: database.GetDatabase(),
	}
}

func (ar *AulaRepository) Create(aula models.Aula) (string, error) {
	err := ar.db.Create(&aula).Error
	return "", err
}

func (ar *AulaRepository) ValidarAula(id uint) bool {

	var aula models.Aula

	err := ar.db.Where(id).First(&aula).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}
