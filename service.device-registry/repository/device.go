package repository

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

	"github.com/jinzhu/copier"

	"github.com/jakewright/home-automation/service.device-registry/domain"
)

// DeviceRepository provides access to the underlying storage layer
type DeviceRepository struct {
	// ConfigFilename is the path to the device config file
	ConfigFilename string

	// ReloadInterval is the amount of time to wait before reading from disk again
	ReloadInterval time.Duration

	devices  []*domain.Device
	reloaded time.Time
	lock     sync.RWMutex
}

// FindAll returns all devices
func (r *DeviceRepository) FindAll() ([]*domain.Device, error) {
	if err := r.reload(); err != nil {
		return nil, err
	}

	r.lock.RLock()
	defer r.lock.RUnlock()

	var devices []*domain.Device
	for _, device := range r.devices {
		out := &domain.Device{}
		if err := copier.Copy(&out, device); err != nil {
			return nil, err
		}
		devices = append(devices, out)
	}

	return devices, nil
}

// Find returns a device by ID
func (r *DeviceRepository) Find(id string) (*domain.Device, error) {
	if err := r.reload(); err != nil {
		return nil, err
	}

	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, device := range r.devices {
		if device.ID == id {
			out := &domain.Device{}
			if err := copier.Copy(out, device); err != nil {
				return nil, err
			}

			return out, nil
		}
	}

	return nil, nil
}

// FindByController returns all devices with the given controller name
func (r *DeviceRepository) FindByController(controllerName string) ([]*domain.Device, error) {
	// Skip if we've recently reloaded
	if err := r.reload(); err != nil {
		return nil, err
	}

	r.lock.RLock()
	defer r.lock.RUnlock()

	var devices []*domain.Device
	for _, device := range r.devices {
		if device.ControllerName == controllerName {
			out := &domain.Device{}
			if err := copier.Copy(out, device); err != nil {
				return nil, err
			}
			devices = append(devices, out)
		}
	}

	return devices, nil
}

// FindByRoom returns all devices for the given room
func (r *DeviceRepository) FindByRoom(roomID string) ([]*domain.Device, error) {
	if err := r.reload(); err != nil {
		return nil, err
	}

	r.lock.RLock()
	defer r.lock.RUnlock()

	var devices []*domain.Device
	for _, device := range r.devices {
		if device.RoomID == roomID {
			out := &domain.Device{}
			if err := copier.Copy(out, device); err != nil {
				return nil, err
			}
			devices = append(devices, out)
		}
	}

	return devices, nil
}

// reload reads the config and applies changes
func (r *DeviceRepository) reload() error {
	// Skip if we've recently reloaded
	if r.reloaded.Add(r.ReloadInterval).After(time.Now()) {
		return nil
	}

	data, err := ioutil.ReadFile(r.ConfigFilename)
	if err != nil {
		return err
	}

	var cfg struct {
		Devices []*domain.Device `json:"devices"`
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return err
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	r.devices = cfg.Devices

	r.reloaded = time.Now()
	return nil
}
