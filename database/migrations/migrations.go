package migrations

import (
	"log/slog"
	"time"

	"github.com/rdruzian/challenge_interview/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(model.Device{})
	if err != nil {
		slog.Error("Error to migrate database")
		return
	}

	// seed inicial: apenas se tabela estiver vazia
	var count int64
	if err := db.Model(&model.Device{}).Count(&count).Error; err != nil {
		slog.Error("Error counting devices", "err", err)
		return
	}
	if count == 0 {
		devices := []model.Device{
			{Name: "Router A", Brand: "Cisco", State: "available", CreationDate: time.Now().Add(-72 * time.Hour)},
			{Name: "Switch B", Brand: "HP", State: "in-use", CreationDate: time.Now().Add(-48 * time.Hour)},
			{Name: "Phone C", Brand: "Apple", State: "inactive", CreationDate: time.Now().Add(-120 * time.Hour)},
			{Name: "Laptop D", Brand: "Dell", State: "available", CreationDate: time.Now().Add(-24 * time.Hour)},
			{Name: "Sensor E", Brand: "Xiaomi", State: "in-use", CreationDate: time.Now().Add(-10 * time.Hour)},
			{Name: "Camera F", Brand: "Logitech", State: "inactive", CreationDate: time.Now().Add(-5 * time.Hour)},
		}
		if err := db.Create(&devices).Error; err != nil {
			slog.Error("Error seeding devices", "err", err)
		} else {
			slog.Info("Seed devices inserted", "count", len(devices))
		}
	}
}
