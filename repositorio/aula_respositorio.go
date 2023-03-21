package repositorio

import (
	"application/database"
	"application/models"
)

func Create(aula models.Aula) (string, error) {
	db := database.GetDatabase()
	err := db.Create(&aula).Error
	return "", err
}
