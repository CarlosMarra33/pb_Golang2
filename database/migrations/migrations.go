package migrations

import (
	"application/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Aula{})
	db.AutoMigrate(models.Presenca{})
}
