package igcclient

import (
	"net/http"
	"testing"
)

func TestUserVerificationService_RegistrationTypes(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/verify/registration/types", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"CountryID\":1,\"VerificationTypeID\":1,\"VerificationType\":{\"ID\":1}},{\"CountryID\":1,\"VerificationTypeID\":2,\"VerificationType\":{\"ID\":2}}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.UserVerification.RegistrationTypes()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}
