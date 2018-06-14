package igcclient

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/moonwalker/igcclient/models"
)

func TestDevicesService_Devices(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"DeviceTypeId\":1},{\"DeviceTypeId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Devices.Devices()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 devices")
	}
}

func TestDevicesService_DevicesByID(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/devices/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"DeviceTypeId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Devices.DevicesByID(1)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data.DeviceTypeID != 1 {
		t.Errorf("Expected device type ID to be 1")
	}
}

func TestDevicesService_GetPlatform(t *testing.T) {
	teardown := setup()
	defer teardown()

	userAgent := "userAgent"

	mux.HandleFunc("/devices/getplatform", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.DevicesDetectionModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.UserAgent != userAgent {
			t.Errorf("Wrong UserAgent")
			*success = false
		}
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"PlatformName\":\"" + *model.UserAgent + "\"},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.DevicesDetectionModel{
		UserAgent: &userAgent,
	}

	response, err := client.Devices.GetPlatform(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}

	if *response.Data.PlatformName != userAgent {
		t.Errorf("Expected user agent to be: " + userAgent)
	}
}
