package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type DevicesService service

// Get a single device by device ID
func (s *DevicesService) DevicesByID(deviceId int64) (response OperationResponseOfDeviceTypeObject, err error) {
	id := strconv.FormatInt(deviceId, 10)
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
