package igcclient

/*
func TestBonusesService_BonusCodeDetails(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/bonuses/bonuscodedetails/abc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"BonusId\":1},{\"BonusId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	details, err := client.Bonuses.BonusCodeDetails("abc", authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*details.Data) != 2 {
		t.Errorf("Expected 2 public bonus objects")
	}
}

func TestBonusesService_Bonuses(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/bonuses", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"BonusId\":1},{\"BonusId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	details, err := client.Bonuses.Bonuses(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*details.Data) != 2 {
		t.Errorf("Expected 2 bonus objects")
	}
}

func TestBonusesService_BonusesByUserId(t *testing.T) {
	teardown := setup()
	defer teardown()

	userID := int64(1337)

	mux.HandleFunc("/bonuses", func(w http.ResponseWriter, r *http.Request) {
		values, _ := url.ParseQuery(r.URL.RawQuery)
		id, _ := strconv.Atoi(values.Get("userid"))
		if int64(id) != userID {
			t.Errorf("Expected user ID: %v", userID)
		}
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"BonusId\":1},{\"BonusId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	details, err := client.Bonuses.BonusesByUserID(userID, xAPIKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*details.Data) != 2 {
		t.Errorf("Expected 2 bonus objects")
	}
}

func TestBonusesService_Credit(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/bonuses/credit/abc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"UserBonuses\":{}},{\"UserBonuses\":{}}],\"Success\":true,\"Errors\":[]}"))
	})

	details, err := client.Bonuses.Credit("abc", authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*details.Data) != 2 {
		t.Errorf("Expected 2 bonus objects")
	}
}
*/
