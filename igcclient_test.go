package igcclient

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the IGC client being tested.
	client IGCClient

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
	c, err := NewIGCClient(parsed.String(), &testLogger{})
	if err != nil {
		fmt.Println("Error setting up client")
	}
	client = *c

	return func() {
		server.Close()
	}
}

func TestNewIGCClient(t *testing.T) {
	_, err := NewIGCClient("", &testLogger{})
	if err == nil {
		t.Errorf("Expected error since baseUrl is empty")
	}

	c, err := NewIGCClient("not://valid!url", &testLogger{})
	if err != nil {
		t.Errorf(err.Error())
	}

	err = c.apiPost("test", nil, &headers)
	if err == nil {
		t.Errorf("Expected error since the url should not be parseble")
	}

	teardown := setup()
	defer teardown()

	mux.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("{\"Data\":[{\"DeviceTypeId\":1},{\"DeviceTypeId\":2}],\"Success\":true,\"Errors\":[]}"))
	})

	client.Devices.Devices()
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
