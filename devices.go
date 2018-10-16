package igcclient

import (
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type DevicesService service

// Get a single device by device ID
func (s *DevicesService) DevicesByID(deviceId int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfDeviceTypeObject, err error) {
	id := strconv.FormatInt(deviceId, 10)
	err = s.client.apiPost("/devices/"+url.QueryEscape(id), nil, nil, &response, &headers, log)
	return
}

// Get all the device types
func (s *DevicesService) Devices(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfDeviceTypeObject, err error) {
	err = s.client.apiPost("/devices", nil, nil, &response, &headers, log)
	return
}

// Get the platform from the details passed in the DevicesDetectionModel
func (s *DevicesService) GetPlatform(body models.DevicesDetectionModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfDevicePlatformModel, err error) {
	err = s.client.apiPost("/devices/getplatform", nil, &body, &response, &headers, log)
	return
}
