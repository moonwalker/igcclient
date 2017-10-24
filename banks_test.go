package igcclient

import (
	"net/http"
	"testing"
)

func TestBanksService_Banks(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/banks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"BankId\":1},{\"BankId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	countries, err := client.Banks.Banks()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*countries.Data) != 2 {
		t.Errorf("Expected 2 banks")
	}
}

func TestBanksService_BankById(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/banks/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"BankId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	country, err := client.Banks.BankById(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if *country.Data.BankId != 1 {
		t.Errorf("BankId is not 1")
	}
}
