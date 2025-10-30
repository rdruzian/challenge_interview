package migrations

import (
	"github.com/rdruzian/challenge_interview/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.Device{})
}
