package repositorio

import (
	"application/database"
	"application/models"
)

type AulaRepository struct {}

func (aularepo *AulaRepository) Create(aula models.Aula) (string, error) {
	db := database.GetDatabase()
	err := db.Create(&aula).Error
	return "", err
}
