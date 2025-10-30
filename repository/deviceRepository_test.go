package repository

import (
	"testing"
	"time"

	"github.com/rdruzian/challenge_interview/database"
	"github.com/rdruzian/challenge_interview/model"
)

func TestRepository_CRUD_WithRealDB(t *testing.T) {
	// inicia DB local (espera docker-compose rodando)
	database.StartDB()
	db := database.GetDatabase()
	repo := NewDeviceRepository(db)

	// cria um device de teste
	name := "TestDevice-" + time.Now().Format("150405")
	dev := model.Device{Name: name, Brand: "TestBrandX", State: "available", CreationDate: time.Now().UTC()}
	if err := repo.CreateDevice(dev); err != nil {
		t.Fatalf("create error: %v", err)
	}

	// busca por brand para encontrar ID inserido
	list, err := repo.GetDeviceByBrand("TestBrandX")
	if err != nil {
		t.Fatalf("get by brand error: %v", err)
	}
	var created model.Device
	for _, d := range list {
		if d.Name == name {
			created = d
			break
		}
	}
	if created.ID == 0 {
		t.Fatalf("inserted device not found by brand")
	}

	// update state
	created.State = "inactive"
	updated, err := repo.UpdateDevice(created)
	if err != nil {
		t.Fatalf("update error: %v", err)
	}
	if updated.State != "inactive" {
		t.Fatalf("expected state inactive, got %s", updated.State)
	}

	// get by id
	byID, err := repo.GetDevice(int(created.ID))
	if err != nil {
		t.Fatalf("get by id error: %v", err)
	}
	if byID.Name != name {
		t.Fatalf("unexpected name by id: %s", byID.Name)
	}

	// list by state seeded
	if devices, err := repo.GetDeviceByState("available"); err != nil || len(devices) == 0 {
		t.Fatalf("expected available devices from seed")
	}

	// cleanup
	if err := repo.DeleteDevice(int(created.ID)); err != nil {
		t.Fatalf("cleanup delete error: %v", err)
	}
}