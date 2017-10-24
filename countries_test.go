package igcclient

import (
	"net/http"
	"testing"
)

func TestCountriesService_CountryById(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/countries/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"CountryID\":1},\"Success\":true,\"Errors\":[]}"))
	})

	country, err := client.Countries.CountryById(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if *country.Data.CountryID != 1 {
		t.Errorf("CountryID is not 1")
	}
}

func TestCountriesService_CountriesTop(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/countries/top", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"CountryID\":1},{\"CountryID\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	countries, err := client.Countries.CountriesTop()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*countries.Data) != 2 {
		t.Errorf("Expected 2 countries")
	}
}

func TestCountriesService_CountryByAlphaCode2(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/countries/en", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"CountryID\":1},\"Success\":true,\"Errors\":[]}"))
	})

	country, err := client.Countries.CountryByAlphaCode2("en")
	if err != nil {
		t.Errorf(err.Error())
	}
	if *country.Data.CountryID != 1 {
		t.Errorf("CountryID is not 1")
	}
}

func TestCountriesService_CountryByAlphaCode3(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/countries/eng", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"CountryID\":1},\"Success\":true,\"Errors\":[]}"))
	})

	country, err := client.Countries.CountryByAlphaCode3("eng")
	if err != nil {
		t.Errorf(err.Error())
	}
	if *country.Data.CountryID != 1 {
		t.Errorf("CountryID is not 1")
	}
}

func TestCountriesService_Countries(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"CountryID\":1},{\"CountryID\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	countries, err := client.Countries.Countries()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*countries.Data) != 2 {
		t.Errorf("Expected 2 countries")
	}
}
