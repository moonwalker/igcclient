package igcclient

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the IGC client being tested.
	client *IGCClient

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server

	xAPIKey   = "xApiKey"
	authToken = "authToken"
)

// setup sets up a test HTTP server along with a igcclient.IGCClient that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() func() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	parsed, _ := url.Parse(server.URL)
	// IGCClient client configured to use test server
	c, err := NewClient(parsed.String())
	if err != nil {
		log.Panic(err)
	}
	client = c

	return func() {
		server.Close()
	}
}

func NewClient(serverHost string) (*IGCClient, error) {
	logRequestBlacklist := []string{
		"/user",
		"/user/closeaccount",
		"/v2/authentication/login",
		"/authentication/change/password",
		"/authentication/change/email",
		"/authentication/change/securityquestion",
		"/authentication/forgotpassword/change/sms",
		"/authentication/forgotpassword/change",
		"/v2/authentication/register",
	}

	logResponseBlacklist := []string{
		"/consent/getconsents",
		"/consent/userconsents",
		"/countries",
		"/user",
	}

	logBlacklist := []string{
		"/authentication/isloggedin",
	}

	invalidAuthCallback := tokenInvalid

	igcClient, err := NewIGCClient(Config{
		BaseURL:              fmt.Sprintf("%s", serverHost),
		LogRequestBody:       true,
		LogResponseData:      true,
		LogRequestBlacklist:  logRequestBlacklist,
		LogResponseBlacklist: logResponseBlacklist,
		LogBlacklist:         logBlacklist,
		Debug:                true,
		InvalidAuthCallback:  &invalidAuthCallback,
	})

	if err != nil {
		return nil, err
	}

	return igcClient, nil
}

func tokenInvalid(token string) {

}

func TestTimeoutError(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(31 * time.Second)
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"DeviceTypeId\":1},{\"DeviceTypeId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	logger := testLogger{}

	_, err := client.Devices.Devices(make(map[string]string), logger)

	if err != nil && err.Error() == ErrorClientTimeout {
		fmt.Println("Yey, we got timeout error as we wanted!")
	} else if err != nil {
		t.Errorf("Expected ErrorClientTimeout but got: %s", err)
	} else {
		t.Error("Expected ErrorClientTimeout")
	}
}

type testLogger struct{}

func (l testLogger) DebugLevel(debug bool) {}

func (l testLogger) SetContext(context map[string]interface{}) {}

func (l testLogger) Println(args ...interface{}) {
	log.Println(args...)
}

func (l testLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l testLogger) Debug(msg string, fields map[string]interface{}) {
	log.Printf(msg, fields)
}

func (l testLogger) Info(msg string, fields map[string]interface{}) {
	log.Printf(msg, fields)
}

func (l testLogger) Error(msg string, err error) {
	log.Printf(msg, err)
}

func (l testLogger) Fatal(msg string, err error) {
	log.Fatal(msg, err)
}
