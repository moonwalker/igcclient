package igcclient

/*
func TestValidationService_Email(t *testing.T) {
	teardown := setup()
	defer teardown()

	email := "email@email.email"

	mux.HandleFunc("/validate/email", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ValidationEmailModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Email != email {
			t.Errorf("Wrong Email")
			*success = false
		}

		q := r.URL.Query().Get("ignoreExisting")
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + q + ",\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ValidationEmailModel{
		Email: &email,
	}

	testIgnoreExistingFalse := func() {
		response, err := client.Validation.Email(model, false)
		if err != nil {
			t.Errorf(err.Error())
		}
		if *response.Data {
			t.Errorf("Expected false")
		}
		if !*response.Success {
			t.Errorf("Expected true")
		}
	}

	testIgnoreExistingTrue := func() {
		response, err := client.Validation.Email(model, true)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !*response.Data {
			t.Errorf("Expected true")
		}
		if !*response.Success {
			t.Errorf("Expected true")
		}
	}

	testIgnoreExistingFalse()
	testIgnoreExistingTrue()
}

func TestValidationService_Mobile(t *testing.T) {
	teardown := setup()
	defer teardown()

	prefix := "prefix"
	mobile := "mobile"

	mux.HandleFunc("/validate/mobile", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ValidationMobileModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Prefix != prefix {
			t.Errorf("Wrong Prefix")
			*success = false
		}

		if *model.Mobile != mobile {
			t.Errorf("Wrong Mobile")
			*success = false
		}

		q := r.URL.Query().Get("ignoreExisting")
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + q + ",\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ValidationMobileModel{
		Prefix: &prefix,
		Mobile: &mobile,
	}

	testIgnoreExistingFalse := func() {
		response, err := client.Validation.Mobile(model, false)
		if err != nil {
			t.Errorf(err.Error())
		}
		if *response.Data {
			t.Errorf("Expected false")
		}
		if !*response.Success {
			t.Errorf("Expected true")
		}
	}

	testIgnoreExistingTrue := func() {
		response, err := client.Validation.Mobile(model, true)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !*response.Data {
			t.Errorf("Expected true")
		}
		if !*response.Success {
			t.Errorf("Expected true")
		}
	}

	testIgnoreExistingFalse()
	testIgnoreExistingTrue()
}

func TestValidationService_Username(t *testing.T) {
	teardown := setup()
	defer teardown()

	username := "username"

	mux.HandleFunc("/validate/username", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ValidationUsernameModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Username != username {
			t.Errorf("Wrong Username")
			*success = false
		}

		q := r.URL.Query().Get("ignoreExisting")
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + q + ",\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ValidationUsernameModel{
		Username: &username,
	}

	testIgnoreExistingFalse := func() {
		response, err := client.Validation.Username(model, false)
		if err != nil {
			t.Errorf(err.Error())
		}
		if *response.Data {
			t.Errorf("Expected false")
		}
		if !*response.Success {
			t.Errorf("Expected true")
		}
	}

	testIgnoreExistingTrue := func() {
		response, err := client.Validation.Username(model, true)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !*response.Data {
			t.Errorf("Expected true")
		}
		if !*response.Success {
			t.Errorf("Expected true")
		}
	}

	testIgnoreExistingFalse()
	testIgnoreExistingTrue()
}
*/