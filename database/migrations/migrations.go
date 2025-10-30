package migrations

import (
	"log/slog"

	"github.com/rdruzian/challenge_interview/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(model.Device{})
	if err != nil {
		slog.Error("Error to migrate database")
		return
	}
}
