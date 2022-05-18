package migrations

import (
	"github.com/G-Basilio/goApiRest/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
