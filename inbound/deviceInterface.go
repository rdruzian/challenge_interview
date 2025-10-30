package inbound

import "github.com/rdruzian/challenge_interview/model"

type DeviceInterface interface {
	CreateDevice(device model.Device) error
	UpdateDevice(device model.Device) (model.Device, error)
	GetDevice(id int) (model.Device, error)
	GetAllDevice() ([]model.Device, error)
	GetDeviceByBrand(brand string) ([]model.Device, error)
	GetDeviceByState(state string) ([]model.Device, error)
	DeleteDevice(id int) error
}
