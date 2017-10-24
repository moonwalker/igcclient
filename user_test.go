package igcclient

import (
	"bytes"
	"encoding/json"
	"github.com/moonwalker/igcclient/models"
	"net/http"
	"strconv"
	"testing"
)

func TestUserService_AcceptTerms(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/acceptterms", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":true,\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.AcceptTerms(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_AddKYCRequirements(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/addkycrequirements", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":true,\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.AddKYCRequirements(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_AddUserNote(t *testing.T) {
	teardown := setup()
	defer teardown()

	note := "note with space"

	mux.HandleFunc("/user/addusernote", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("note")
		if q == note {
			w.WriteHeader(200)
			w.Write([]byte("{\"Data\":true,\"Success\":true,\"Errors\":[]}"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("{\"Data\":false,\"Success\":true,\"Errors\":[]}"))
		}
	})

	response, err := client.User.AddUserNote(note, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_CheckCombination(t *testing.T) {
	teardown := setup()
	defer teardown()

	firstname := "firstname"
	lastname := "lastname"
	address := "address"
	phone := "phone"

	mux.HandleFunc("/user/check/combination", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.UserObject)

		json.NewDecoder(r.Body).Decode(model)

		if *model.FirstName != firstname {
			t.Errorf("Wrong FirstName")
			*success = false
		}
		if *model.LastName != lastname {
			t.Errorf("Wrong LastName")
			*success = false
		}
		if *model.Address1 != address {
			t.Errorf("Wrong Address")
			*success = false
		}
		if *model.Phone != phone {
			t.Errorf("Wrong Phone")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + strconv.FormatBool(*success) + ",\"Success\":true,\"Errors\":[]}"))
	})

	model := models.UserObject{
		FirstName: &firstname,
		LastName:  &lastname,
		Address1:  &address,
		Phone:     &phone,
	}

	response, err := client.User.CheckCombination(model)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_CheckEmail(t *testing.T) {
	teardown := setup()
	defer teardown()

	email := "email@email.email"

	mux.HandleFunc("/user/check/email", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.UserObject)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Email != email {
			t.Errorf("Wrong Email")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + strconv.FormatBool(*success) + ",\"Success\":true,\"Errors\":[]}"))
	})

	model := models.CheckUser{
		Email: &email,
	}

	response, err := client.User.CheckEmail(model)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_CheckUsername(t *testing.T) {
	teardown := setup()
	defer teardown()

	username := "username"

	mux.HandleFunc("/user/check/username", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.UserObject)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Username != username {
			t.Errorf("Wrong Username")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + strconv.FormatBool(*success) + ",\"Success\":true,\"Errors\":[]}"))
	})

	model := models.CheckUser{
		Username: &username,
	}

	response, err := client.User.CheckUsername(model)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_GetLoginHistory(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/getloginhistory", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"AdminUserID\":1},{\"AdminUserID\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.GetLoginHistory(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_GetSimilarUserCount(t *testing.T) {
	teardown := setup()
	defer teardown()

	firstname := "firstname"
	lastname := "lastname"

	mux.HandleFunc("/user/getsimilarusercount", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.SimilarUser)

		json.NewDecoder(r.Body).Decode(model)

		if *model.FirstName != firstname {
			t.Errorf("Wrong Firstname")
			*success = false
		}
		if *model.LastName != lastname {
			t.Errorf("Wrong Lastname")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":2,\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.SimilarUser{
		FirstName: &firstname,
		LastName:  &lastname,
	}

	response, err := client.User.GetSimilarUserCount(model)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Success {
		t.Errorf("Expected true")
	}
	if *response.Data != 2 {
		t.Errorf("Expected 2")
	}
}

func TestUserService_GetUsersWithNoActivity(t *testing.T) {
	teardown := setup()
	defer teardown()

	fromdate := "2017-09-19"
	limit := 2

	mux.HandleFunc("/user/getuserswithnoactivity", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		q1 := r.URL.Query().Get("fromdate")
		q2 := r.URL.Query().Get("limit")

		if q1 != fromdate {
			t.Errorf("Wrong from date")
			*success = false
		}
		q2int, _ := strconv.Atoi(q2)
		if q2int != limit {
			t.Errorf("Wrong limit")
			*success = false
		}
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"UserId\":1},{\"UserId\":2}],\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	response, err := client.User.GetUsersWithNoActivity(fromdate, limit, xApiKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Success {
		t.Errorf("Expected true")
	}
	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_KYC(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/kyc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"KycId\":1},{\"KycId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.KYC(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_KYCByUserId(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/kyc", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("userid")
		qInt, _ := strconv.Atoi(q)
		if qInt != 1 {
			t.Errorf("Expected userId to be 1")
			w.WriteHeader(200)
			w.Write([]byte("{\"Data\":[],\"Success\":false,\"Errors\":[]}"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("{\"Data\":[{\"KycId\":1},{\"KycId\":2}],\"Success\":true,\"Errors\":[]}"))
		}
	})

	response, err := client.User.KYCByUserId(1, xApiKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Success {
		t.Errorf("Expected true")
	}
	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_KYCUpload(t *testing.T) {
	teardown := setup()
	defer teardown()

	data := []byte("data")
	kycid := 1
	extension := "extension"

	mux.HandleFunc("/user/kyc/upload", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.KYCUpload)

		json.NewDecoder(r.Body).Decode(model)

		if bytes.Compare(*model.Data, data) != 0 {
			t.Errorf("Wrong Data")
			*success = false
		}
		if *model.KYCID != kycid {
			t.Errorf("Wrong KYCID")
			*success = false
		}
		if *model.Extension != extension {
			t.Errorf("Wrong Extension")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":true,\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.KYCUpload{
		Data:      &data,
		KYCID:     &kycid,
		Extension: &extension,
	}

	response, err := client.User.KYCUpload(model, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Success {
		t.Errorf("Expected true")
	}
	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestUserService_LimitsGetLimitDurations(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/limits/getlimitdurations", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"LimitDurationType\":1},{\"LimitDurationType\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.LimitsGetLimitDurations()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_LimitsGetLimits(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/limits/getlimits", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"LimitType\":1},{\"LimitType\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.LimitsGetLimits()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_LimitsGetUserLimits(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/limits/getlimits", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"LimitType\":1},{\"LimitType\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.LimitsGetLimits()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_LimitsGetUserLimitsByUserId(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/limits/getuserlimits", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("userId")
		qInt, _ := strconv.Atoi(q)
		if qInt != 1 {
			t.Errorf("Expected userId to be 1")
		}
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"Id\":1},{\"Id\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.User.LimitsGetUserLimitsByUserId(1, xApiKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(*response.Data) != 2 {
		t.Errorf("Expected 2 items")
	}
}

func TestUserService_LimitsSetUserLimit(t *testing.T) {
	teardown := setup()
	defer teardown()

	userid := 1
	limittype := models.LT_CasinoNetLoss
	duration := models.LD_1Week
	limitamount := 0.5

	mux.HandleFunc("/user/limits/setuserlimit", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.SetUserLimitModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.UserId != userid {
			t.Errorf("Wrong UserId")
			*success = false
		}
		if *model.Type != limittype {
			t.Errorf("Wrong Type")
			*success = false
		}
		if *model.Duration != duration {
			t.Errorf("Wrong Duration")
			*success = false
		}
		if *model.LimitAmount != limitamount {
			t.Errorf("Wrong LimitAmount")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":true,\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.SetUserLimitModel{
		UserId:      &userid,
		Type:        &limittype,
		Duration:    &duration,
		LimitAmount: &limitamount,
	}

	response, err := client.User.LimitsSetUserLimit(model, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !*response.Success {
		t.Errorf("Expected true")
	}
}


func TestUserService_SendEmail(t *testing.T) {

}

func TestUserService_SessionTurnover(t *testing.T) {

}

func TestUserService_SetAffiliateReference(t *testing.T) {

}

func TestUserService_Update(t *testing.T) {

}

func TestUserService_User(t *testing.T) {

}

func TestUserService_UserById(t *testing.T) {

}
