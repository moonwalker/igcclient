package igcclient

import (
	"net/http"
	"strconv"
	"testing"
)

func TestRolesService_Roles(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"Id\":1,\"Name\":\"test1\"},{\"Id\":2,\"Name\":\"test2\"}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Roles.Roles(xApiKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestRolesService_RolesByUserId(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/roles/user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("userId")
		userId, err := strconv.Atoi(id)
		if err != nil {
			t.Errorf("Expectet userId to be set")
		}
		if userId != 1 {
			t.Errorf("Expected userId to be 1")
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"Id\":1,\"Name\":\"test1\"}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Roles.RolesByUserId(1, xApiKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 1 {
		t.Errorf("Expected 1 item")
	}
}
