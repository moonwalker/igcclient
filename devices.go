package igcclient

import (
	. "github.com/moonwalker/igcclient/models"
	"strconv"
	"net/url"
)

type DevicesService service

// Get a single device by device ID
func (s *DevicesService) DevicesById(deviceId int) (response OperationResponseOfDeviceTypeObject, err error) {
	id := strconv.Itoa(deviceId)
	err = s.client.apiPost("/devices/"+url.QueryEscape(id), nil, &response, nil, nil)
	return
}

// Get all the device types
func (s *DevicesService) Devices() (response OperationResponseOfListOfDeviceTypeObject, err error) {
	err = s.client.apiPost("/devices", nil, &response, nil, nil)
	return
}

// Get the platform from the details passed in the DevicesDetectionModel
func (s *DevicesService) GetPlatform(body DevicesDetectionModel) (response OperationResponseOfDevicePlatformModel, err error) {
	err = s.client.apiPost("/devices/getplatform", &body, &response, nil, nil)
	return
}
