package repository

import (
	"errors"
	"strings"
	"time"

	"github.com/rdruzian/challenge_interview/model"
	"github.com/rdruzian/challenge_interview/outbound"
	"gorm.io/gorm"
)

func NewDeviceRepository(
	db *gorm.DB,
) outbound.DeviceInterface {
	return &deviceRepository{
		databaseConnection: db,
	}
}

type deviceRepository struct {
	databaseConnection *gorm.DB
}

func (d deviceRepository) CreateDevice(device model.Device) error {
	if device.CreationDate.IsZero() {
		device.CreationDate = time.Now().UTC()
	}
	result := d.databaseConnection.Create(&device)
	return result.Error
}

func (d deviceRepository) UpdateDevice(device model.Device) (model.Device, error) {
	if device.ID == 0 {
		return model.Device{}, errors.New("missing device ID")
	}

	var current model.Device
	if err := d.databaseConnection.First(&current, device.ID).Error; err != nil {
		return model.Device{}, err
	}
	if strings.EqualFold(current.State, "in-use") {
		return model.Device{}, errors.New("cannot update device in-use")
	}

	result := d.databaseConnection.Model(&model.Device{}).
		Where("id = ?", device.ID).
		Updates(map[string]interface{}{
			"name":  device.Name,
			"brand": device.Brand,
			"state": device.State,
		})
	if result.Error != nil {
		return model.Device{}, result.Error
	}
	var updated model.Device
	if err := d.databaseConnection.First(&updated, device.ID).Error; err != nil {
		return model.Device{}, err
	}
	return updated, nil
}

func (d deviceRepository) GetDevice(id int) (model.Device, error) {
	var device model.Device
	result := d.databaseConnection.First(&device, id)
	return device, result.Error
}

func (d deviceRepository) GetAllDevice() ([]model.Device, error) {
	var devices []model.Device
	result := d.databaseConnection.Find(&devices)
	return devices, result.Error
}

func (d deviceRepository) GetDeviceByBrand(brand string) ([]model.Device, error) {
	var devices []model.Device
	result := d.databaseConnection.Where("brand = ?", brand).Find(&devices)
	return devices, result.Error
}

func (d deviceRepository) GetDeviceByState(state string) ([]model.Device, error) {
	var devices []model.Device
	result := d.databaseConnection.Where("state = ?", state).Find(&devices)
	return devices, result.Error
}

func (d deviceRepository) DeleteDevice(id int) error {
	var current model.Device
	if err := d.databaseConnection.First(&current, id).Error; err != nil {
		return err
	}
	if strings.EqualFold(current.State, "in-use") {
		return errors.New("cannot delete device in-use")
	}
	result := d.databaseConnection.Delete(&model.Device{}, id)
	return result.Error
}
