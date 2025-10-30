package service

import (
	"errors"
	"testing"

	"github.com/rdruzian/challenge_interview/model"
)

type stubRepo struct{}

func (s stubRepo) CreateDevice(device model.Device) error { return nil }
func (s stubRepo) UpdateDevice(device model.Device) (model.Device, error) { return device, nil }
func (s stubRepo) GetDevice(id int) (model.Device, error) { return model.Device{ID: int64(id)}, nil }
func (s stubRepo) GetAllDevice() ([]model.Device, error) { return []model.Device{{ID: 1}}, nil }
func (s stubRepo) GetDeviceByBrand(brand string) ([]model.Device, error) { return []model.Device{{Brand: brand}}, nil }
func (s stubRepo) GetDeviceByState(state string) ([]model.Device, error) {
	if state == "invalid" {
		return nil, errors.New("invalid state")
	}
	return []model.Device{{State: state}}, nil
}
func (s stubRepo) DeleteDevice(id int) error { return nil }

func TestCreateDevice_InvalidState(t *testing.T) {
	svc := NewDeviceService(stubRepo{})
	if err := svc.CreateDevice(model.Device{Name: "X", Brand: "Y", State: "wrong"}); err == nil {
		t.Fatalf("expected error for invalid state")
	}
}

func TestUpdateDevice_InvalidState(t *testing.T) {
	svc := NewDeviceService(stubRepo{})
	if _, err := svc.UpdateDevice(model.Device{ID: 1, Name: "X", Brand: "Y", State: "bad"}); err == nil {
		t.Fatalf("expected error for invalid state")
	}
}

func TestGetDeviceByState_OK(t *testing.T) {
	svc := NewDeviceService(stubRepo{})
	res, err := svc.GetDeviceByState("available")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(res) == 0 || res[0].State != "available" {
		t.Fatalf("expected available state result")
	}
}