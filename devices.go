package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type DevicesService service

// Get a single device by device ID
func (s *DevicesService) DevicesByID(deviceId int64, headers map[string]string) (response OperationResponseOfDeviceTypeObject, err error) {
	id := strconv.FormatInt(deviceId, 10)
	err = s.client.apiPost("/devices/"+url.QueryEscape(id), nil, &response, &headers)
	return
}

// Get all the device types
func (s *DevicesService) Devices(headers map[string]string) (response OperationResponseOfListOfDeviceTypeObject, err error) {
	err = s.client.apiPost("/devices", nil, &response, &headers)
	return
}

// Get the platform from the details passed in the DevicesDetectionModel
func (s *DevicesService) GetPlatform(body DevicesDetectionModel, headers map[string]string) (response OperationResponseOfDevicePlatformModel, err error) {
	err = s.client.apiPost("/devices/getplatform", &body, &response, &headers)
	return
}
