package migrations

import (
	"TestInterview/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.Device{})
}
