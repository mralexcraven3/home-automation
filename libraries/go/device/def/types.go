// Code generated by jrpc. DO NOT EDIT.

package devicedef

import (
	errors "github.com/jakewright/home-automation/libraries/go/errors"
)

// Device is defined in the .def file
type Device struct {
	Id             string                 `json:"id"`
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Kind           string                 `json:"kind"`
	ControllerName string                 `json:"controller_name"`
	Attributes     map[string]interface{} `json:"attributes"`
	StateProviders []string               `json:"state_providers"`
	State          map[string]*Property   `json:"state"`
}

// Property is defined in the .def file
type Property struct {
	Value         interface{} `json:"value"`
	Type          string      `json:"type"`
	Min           int32       `json:"min"`
	Max           int32       `json:"max"`
	Interpolation string      `json:"interpolation"`
}

// DeviceStateChangedEvent is defined in the .def file
type DeviceStateChangedEvent struct {
	Device *Device `json:"device"`
}

// Validate returns an error if any of the fields have bad values
func (m *Device) Validate() error {
	return nil
}

// Validate returns an error if any of the fields have bad values
func (m *Property) Validate() error {
	return nil
}

// Validate returns an error if any of the fields have bad values
func (m *DeviceStateChangedEvent) Validate() error {
	if err := m.Device.Validate(); err != nil {
		return err
	}

	if m.Device == nil {
		return errors.BadRequest("field device is required")
	}
	return nil
}
