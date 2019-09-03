package igcclient

/*
func TestRolesService_Roles(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/roles", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"Id\":1,\"Name\":\"test1\"},{\"Id\":2,\"Name\":\"test2\"}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Roles.Roles(xAPIKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestRolesService_RolesByUserID(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/roles/user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("userId")
		userID, err := strconv.Atoi(id)
		if err != nil {
			t.Errorf("Expectet userID to be set")
		}
		if userID != 1 {
			t.Errorf("Expected userID to be 1")
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"Id\":1,\"Name\":\"test1\"}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Roles.RolesByUserID(1, xAPIKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 1 {
		t.Errorf("Expected 1 item")
	}
}
*/