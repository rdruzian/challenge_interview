package inbound

import "TestInterview/model"

type DeviceInterface interface {
	CreateDevice(device model.Device) error
	UpdateDevice(device model.Device) (model.Device, error)
	GetDevice(id int) (model.Device, error)
	GetAllDevice() ([]model.Device, error)
	GetDeviceByBrand(brand string) ([]model.Device, error)
	GetDeviceByState(state int) ([]model.Device, error)
	DeleteDevice(id int) error
}
