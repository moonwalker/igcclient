package igcclient

/**
func TestAuthenticationService_Login(t *testing.T) {
	teardown := setup()
	defer teardown()

	username := "TEST_USERNAME"
	password := "TEST_PASSWORD"

	mux.HandleFunc("/authentication/login", func(w http.ResponseWriter, r *http.Request) {
		loginModel := new(models.LoginModel)

		json.NewDecoder(r.Body).Decode(loginModel)

		if *loginModel.Username != username {
			t.Errorf("Wrong username")
		}
		if *loginModel.Password != password {
			t.Errorf("Wrong password")
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"" + *loginModel.Username + *loginModel.Password + "\",\"Success\":true,\"Errors\":[]}"))
	})

	loginModel := models.LoginModel{
		Username: &username,
		Password: &password,
	}

	authResponse, err := client.Authentication.Login(loginModel)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *authResponse.Data != username+password {
		t.Errorf("Wrong auth token")
	}
}

func TestAuthenticationService_IsLoggedIn(t *testing.T) {
	teardown := setup()
	defer teardown()

	username := "TEST_USERNAME"
	password := "TEST_PASSWORD"

	mux.HandleFunc("/authentication/isloggedin", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if r.Header.Get("AuthenticationToken") == authToken {
			w.Write([]byte("{\"Data\":true,\"Success\":true,\"Errors\":[]}"))
		} else {
			w.Write([]byte("{\"Data\":false,\"Success\":true,\"Errors\":[]}"))
		}
	})

	mux.HandleFunc("/authentication/login", func(w http.ResponseWriter, r *http.Request) {
		loginModel := new(models.LoginModel)

		json.NewDecoder(r.Body).Decode(loginModel)

		if *loginModel.Username != username {
			t.Errorf("Wrong username")
		}
		if *loginModel.Password != password {
			t.Errorf("Wrong password")
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"" + *loginModel.Username + *loginModel.Password + "\",\"Success\":true,\"Errors\":[]}"))
	})

	loginModel := models.LoginModel{
		Username: &username,
		Password: &password,
	}

	client.Authentication.Login(loginModel)

	response2, err := client.Authentication.IsLoggedIn(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response2.Data {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_Logout(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/authentication/logout", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":true,\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Authentication.Logout(authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_AuthenticateWithChallenge(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/authentication/authenticatewithchallenge", func(w http.ResponseWriter, r *http.Request) {
		values, _ := url.ParseQuery(r.URL.RawQuery)
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"" + values.Get("challengeToken") + "\",\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Authentication.AuthenticateWithChallenge("TOKEN")
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data != "TOKEN" {
		t.Errorf("Expected TOKEN as response")
	}
}

func TestAuthenticationService_GetPasswordRegex(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/authentication/getpasswordregex", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"REGEX\",\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Authentication.GetPasswordRegex()
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data != "REGEX" {
		t.Errorf("Expected REGEX")
	}
}

func TestAuthenticationService_VerifyActivationTokenGet(t *testing.T) {
	teardown := setup()
	defer teardown()

	email := "email@email.email"

	mux.HandleFunc("/v2/authentication/verify/activationtoken/get", func(w http.ResponseWriter, r *http.Request) {
		model := new(models.GetVerificationTokenModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Email != email {
			t.Errorf("Wrong email")
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"" + *model.Email + "\",\"Success\":true,\"Errors\":[]}"))
	})

	model := models.GetVerificationTokenModel{
		Email: &email,
	}

	response, err := client.Authentication.VerifyActivationTokenGet(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data != email {
		t.Errorf("Expected " + email)
	}
}

func TestAuthenticationService_VerifySMSSend(t *testing.T) {
	teardown := setup()
	defer teardown()

	prefix := "prefix"
	mobile := "mobile"

	mux.HandleFunc("/v2/authentication/verify/sms/send", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ShortMessageServiceModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.MobilePrefix != prefix {
			t.Errorf("Wrong mobile prefix")
			*success = false
		}

		if *model.Mobile != mobile {
			t.Errorf("Wrong mobile")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":" + strconv.FormatBool(*success) + ",\"Success\":true,\"Errors\":[]}"))
	})

	model := models.ShortMessageServiceModel{
		MobilePrefix: &prefix,
		Mobile:       &mobile,
	}

	response, err := client.Authentication.VerifySMSSend(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_VerifySMS(t *testing.T) {
	teardown := setup()
	defer teardown()

	prefix := "prefix"
	mobile := "mobile"
	code := "code"

	mux.HandleFunc("/v2/authentication/verify/sms", func(w http.ResponseWriter, r *http.Request) {
		model := new(models.VerifyUserBySMSModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.MobilePrefix != prefix {
			t.Errorf("Wrong mobile prefix")
		}

		if *model.Mobile != mobile {
			t.Errorf("Wrong mobile")
		}

		if *model.MobileVerificationCode != code {
			t.Errorf("Wrong mobile verification code")
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"" + prefix + mobile + code + "\",\"Success\":true,\"Errors\":[]}"))
	})

	model := models.VerifyUserBySMSModel{
		MobileVerificationCode: &code,
		MobilePrefix:           &prefix,
		Mobile:                 &mobile,
	}

	response, err := client.Authentication.VerifySMS(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data != prefix+mobile+code {
		t.Errorf("Expected " + prefix + mobile + code)
	}
}

func TestAuthenticationService_VerifyEmail(t *testing.T) {
	teardown := setup()
	defer teardown()

	code := "code"

	mux.HandleFunc("/Authentication/Verify/Email/"+code, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":\"" + code + "\",\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Authentication.VerifyEmail(code)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *response.Data != code {
		t.Errorf("Expected " + code)
	}
}

func TestAuthenticationService_VerifyEmailSend(t *testing.T) {
	teardown := setup()
	defer teardown()

	email := "email@email.email"

	mux.HandleFunc("/authentication/verify/email/send", func(w http.ResponseWriter, r *http.Request) {
		values, _ := url.ParseQuery(r.URL.RawQuery)
		if values.Get("email") != email {
			t.Errorf("Expected email: " + email)
		}
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":true,\"Success\":true,\"Errors\":[]}"))
	})

	response, err := client.Authentication.VerifyEmailSend(email)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Data {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_AllowIP(t *testing.T) {
	teardown := setup()
	defer teardown()

	ip := "127.0.0.1"
	userAgent := "useragent"

	mux.HandleFunc("/authentication/allow/ip", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.AllowIPModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.IP != ip {
			t.Errorf("Wrong IP")
			*success = false
		}

		if *model.UserAgent != userAgent {
			t.Errorf("Wrong UserAgnet")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.AllowIPModel{
		IP:        &ip,
		UserAgent: &userAgent,
	}

	response, err := client.Authentication.AllowIP(model, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_SecurityQuestion(t *testing.T) {
	teardown := setup()
	defer teardown()

	email := "email@email.email"

	mux.HandleFunc("/authentication/securityquestion", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.EmailModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Email != email {
			t.Errorf("Wrong email")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"ID\":1,\"Question\":\"" + *model.Email + "\"},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.EmailModel{
		Email: &email,
	}

	response, err := client.Authentication.SecurityQuestion(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}

	if *response.Data.ID != 1 {
		t.Errorf("Expected 1")
	}

	if *response.Data.Question != email {
		t.Errorf("Expected " + email)
	}
}

func TestAuthenticationService_ChangePassword(t *testing.T) {
	teardown := setup()
	defer teardown()

	oldPassword := "oldPassword"
	newPassword := "newPassword"

	mux.HandleFunc("/authentication/change/password", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ChangePasswordModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.OldPassword != oldPassword {
			t.Errorf("Wrong old password")
			*success = false
		}

		if *model.NewPassword != newPassword {
			t.Errorf("Wrong new password")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ChangePasswordModel{
		OldPassword: &oldPassword,
		NewPassword: &newPassword,
	}

	response, err := client.Authentication.ChangePassword(model, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_ChangeEmail(t *testing.T) {
	teardown := setup()
	defer teardown()

	password := "password"
	email := "email@email.email"

	mux.HandleFunc("/authentication/change/email", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ChangeEmailModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Password != password {
			t.Errorf("Wrong password")
			*success = false
		}

		if *model.Email != email {
			t.Errorf("Wrong email")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ChangeEmailModel{
		Password: &password,
		Email:    &email,
	}

	response, err := client.Authentication.ChangeEmail(model, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_ChangeSecurityQuestion(t *testing.T) {
	teardown := setup()
	defer teardown()

	password := "password"
	questionID := int64(1)
	answer := "answer"

	mux.HandleFunc("/authentication/change/securityquestion", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ChangeSecurityQuestionModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Password != password {
			t.Errorf("Wrong password")
			*success = false
		}

		if *model.SecurityQuestionID != questionID {
			t.Errorf("Wrong security question id")
			*success = false
		}

		if *model.SecurityAnswer != answer {
			t.Errorf("Wrong security answer")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ChangeSecurityQuestionModel{
		Password:           &password,
		SecurityQuestionID: &questionID,
		SecurityAnswer:     &answer,
	}

	response, err := client.Authentication.ChangeSecurityQuestion(model, authToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_ForgotPassword(t *testing.T) {
	teardown := setup()
	defer teardown()

	email := "email@email.email"
	answer := "answer"

	mux.HandleFunc("/authentication/forgotpassword", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ForgotPasswordModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Email != email {
			t.Errorf("Wrong email")
			*success = false
		}

		if *model.SecurityAnswer != answer {
			t.Errorf("Wrong security answer")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ForgotPasswordModel{
		Email:          &email,
		SecurityAnswer: &answer,
	}

	response, err := client.Authentication.ForgotPassword(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_ForgotPasswordSMS(t *testing.T) {
	teardown := setup()
	defer teardown()

	prefix := "prefix"
	mobile := "mobile"

	mux.HandleFunc("/authentication/forgotpassword/sms", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ShortMessageServiceModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Mobile != mobile {
			t.Errorf("Wrong mobile")
			*success = false
		}

		if *model.MobilePrefix != prefix {
			t.Errorf("Wrong mobile prefix")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ShortMessageServiceModel{
		MobilePrefix: &prefix,
		Mobile:       &mobile,
	}

	response, err := client.Authentication.ForgotPasswordSMS(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_ForgotPasswordChangeSMS(t *testing.T) {
	teardown := setup()
	defer teardown()

	prefix := "prefix"
	mobile := "mobile"
	password := "password"
	code := "code"

	mux.HandleFunc("/authentication/forgotpassword/change/sms", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ForgotPasswordChangeBySMSModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Mobile != mobile {
			t.Errorf("Wrong mobile")
			*success = false
		}

		if *model.MobilePrefix != prefix {
			t.Errorf("Wrong mobile prefix")
			*success = false
		}

		if *model.Password != password {
			t.Errorf("Wrong password")
			*success = false
		}

		if *model.RecoveryCode != code {
			t.Errorf("Wrong recovery code")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ForgotPasswordChangeBySMSModel{
		MobilePrefix: &prefix,
		Mobile:       &mobile,
		Password:     &password,
		RecoveryCode: &code,
	}

	response, err := client.Authentication.ForgotPasswordChangeSMS(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_ForgotPasswordChange(t *testing.T) {
	teardown := setup()
	defer teardown()

	password := "password"
	code := "code"

	mux.HandleFunc("/authentication/forgotpassword/change", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.ForgotPasswordChangeModel)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Password != password {
			t.Errorf("Wrong password")
			*success = false
		}

		if *model.RecoveryCode != code {
			t.Errorf("Wrong recovery code")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.ForgotPasswordChangeModel{
		Password:     &password,
		RecoveryCode: &code,
	}

	response, err := client.Authentication.ForgotPasswordChange(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !*response.Success {
		t.Errorf("Expected true")
	}
}

func TestAuthenticationService_Register(t *testing.T) {
	teardown := setup()
	defer teardown()

	username := "username"
	password := "password"
	postalcode := "postalcode"
	address1 := "address1"

	mux.HandleFunc("/v2/authentication/register", func(w http.ResponseWriter, r *http.Request) {
		s := true
		success := &s

		model := new(models.RegistrationData)

		json.NewDecoder(r.Body).Decode(model)

		if *model.Password != password {
			t.Errorf("Wrong password")
			*success = false
		}

		if *model.Username != username {
			t.Errorf("Wrong username")
			*success = false
		}

		if *model.PostalCode != postalcode {
			t.Errorf("Wrong postal code")
			*success = false
		}

		if *model.Address1 != address1 {
			t.Errorf("Wrong adress1")
			*success = false
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":{\"Username\":\"" + *model.Username + "\"},\"Success\":" + strconv.FormatBool(*success) + ",\"Errors\":[]}"))
	})

	model := models.RegistrationData{
		Username:   &username,
		Password:   &password,
		PostalCode: &postalcode,
		Address1:   &address1,
	}

	response, err := client.Authentication.Register(model)
	if err != nil {
		t.Errorf(err.Error())
	}

	stringMap := *response.Data

	if stringMap["Username"] != username {
		t.Errorf("Expected " + username)
	}
}
*/
