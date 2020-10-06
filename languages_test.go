package igcclient

/*
func TestLanguagesService_Languages(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("isLiveOnly")
		isLiveOnly, err := strconv.ParseBool(q)
		if err != nil {
			t.Errorf("Expected isLiveOnly parameter to be set as boolean")
		}

		if isLiveOnly {
			w.WriteHeader(200)
			w.Write([]byte("{\"Data\":[{\"LanguageId\":1,\"IsLive\":true}],\"Success\":true,\"Errors\":[]}"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("{\"Data\":[{\"LanguageId\":1,\"IsLive\":true},{\"LanguageID\":2,\"IsLive\":false}],\"Success\":true,\"Errors\":[]}"))
		}
	})

	testIsLiveOnly := func() {
		response, err := client.Languages.Languages(true)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(*response.Data) != 1 {
			t.Errorf("Expected 1 device")
		}
	}

	testIsNotLiveOnly := func() {
		response, err := client.Languages.Languages(false)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(*response.Data) != 2 {
			t.Errorf("Expected 2 devices")
		}
	}

	testIsLiveOnly()
	testIsNotLiveOnly()
}

func TestLanguagesService_LanguagesByAlphaCode2(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/languages/en", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"LanguageId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Languages.LanguagesByAlphaCode2("en")
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data.LanguageID != 1 {
		t.Errorf("Expected language id to be 1")
	}
}

func TestLanguagesService_LanguagesByAlphaCode3(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/languages/eng", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"LanguageId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Languages.LanguagesByAlphaCode3("eng")
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data.LanguageID != 1 {
		t.Errorf("Expected language id to be 1")
	}
}

func TestLanguagesService_LanguagesByID(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/languages/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"LanguageId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Languages.LanguagesByID(1)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data.LanguageID != 1 {
		t.Errorf("Expected language ID to be 1")
	}
}

func TestLanguagesService_LanguagesByName(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/languages/english", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"LanguageId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Languages.LanguagesByName("english")
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data.LanguageID != 1 {
		t.Errorf("Expected language ID to be 1")
	}
}
*/
