package migrations

import (
	"github.com/RossLaing8417/react-go-mvc/server/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Business{},
		&models.Address{},
	)
}
