package igcclient

import (
	"net/http"
	"testing"
)

func TestCurrenciesService_Currencies(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"CurrencyId\":1},{\"CurrencyId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	countries, err := client.Currencies.Currencies()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*countries.Data) != 2 {
		t.Errorf("Expected 2 currencies")
	}
}

func TestCurrenciesService_CurrencyById(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/currencies/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"CurrencyId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	country, err := client.Currencies.CurrencyById(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if *country.Data.CurrencyId != 1 {
		t.Errorf("CurrencyId is not 1")
	}
}

func TestCurrenciesService_CurrencyByCode(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/currencies/en", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"CurrencyId\":1},\"Success\":true,\"Errors\":[]}"))
	})

	country, err := client.Currencies.CurrencyByCode("en")
	if err != nil {
		t.Errorf(err.Error())
	}
	if *country.Data.CurrencyId != 1 {
		t.Errorf("CurrencyId is not 1")
	}
}