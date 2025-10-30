package service

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/rdruzian/challenge_interview/inbound"
	"github.com/rdruzian/challenge_interview/model"
	"github.com/rdruzian/challenge_interview/outbound"
)

func NewDeviceService(deviceRepository outbound.DeviceInterface) inbound.DeviceInterface {
	return &deviceService{
		deviceRepository,
	}
}

type deviceService struct {
	deviceRepository outbound.DeviceInterface
}

var allowedStates = map[string]struct{}{
	"available": {},
	"in-use":    {},
	"inactive":  {},
}

func normalizeState(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func validateState(s string) error {
	state := normalizeState(s)
	if _, ok := allowedStates[state]; !ok {
		return errors.New("invalid state; allowed: available, in-use, inactive")
	}
	return nil
}

func (d deviceService) CreateDevice(device model.Device) error {
	device.State = normalizeState(device.State)
	slog.Info("request_device", "device", device)
	if err := validateState(device.State); err != nil {
		return err
	}
	return d.deviceRepository.CreateDevice(device)
}

func (d deviceService) UpdateDevice(device model.Device) (model.Device, error) {
	if err := validateState(device.State); err != nil {
		return model.Device{}, err
	}
	return d.deviceRepository.UpdateDevice(device)
}

func (d deviceService) GetDevice(id int) (model.Device, error) {
	return d.deviceRepository.GetDevice(id)
}

func (d deviceService) GetAllDevice() ([]model.Device, error) {
	return d.deviceRepository.GetAllDevice()
}

func (d deviceService) GetDeviceByBrand(brand string) ([]model.Device, error) {
	return d.deviceRepository.GetDeviceByBrand(brand)
}

func (d deviceService) GetDeviceByState(state string) ([]model.Device, error) {
	if err := validateState(state); err != nil {
		return nil, err
	}
	return d.deviceRepository.GetDeviceByState(normalizeState(state))
}

func (d deviceService) DeleteDevice(id int) error {
	return d.deviceRepository.DeleteDevice(id)
}
