package igcclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/moonwalker/logger"
)

const (
	timeout = 30 * time.Second
)

type IGCClient struct {
	HTTPClient *http.Client
	baseURL    string

	common            service
	Authentication    *AuthenticationService
	Banks             *BanksService
	Bonuses           *BonusesService
	Consent           *ConsentService
	Countries         *CountriesService
	Currencies        *CurrenciesService
	Devices           *DevicesService
	Games             *GamesService
	IPWhois           *IPWhoisService
	KYC               *KYCService
	Languages         *LanguagesService
	Payments          *PaymentsService
	RealityCheck      *RealityCheckService
	ResponsibleGaming *ResponsibleGamingService
	Roles             *RolesService
	SecurityQuestions *SecurityQuestionsService
	User              *UserService
	Validation        *ValidationService
	Wallet            *WalletService

	logRequestBody       bool
	logResponseData      bool
	logRequestBlacklist  []string
	logResponseBlacklist []string
	logBlacklist         []string

	debug bool
}

type service struct {
	client *IGCClient
}

func NewIGCClient(baseURL string, logRequestBody bool, logResponseData bool, logRequestBlacklist, logResponseBlacklist []string, logBlacklist []string, debug bool) (client *IGCClient, err error) {
	if baseURL == "" {
		err = errors.New("base url can not be empty")
		return
	}
	client = &IGCClient{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		baseURL:              baseURL,
		logRequestBody:       logRequestBody,
		logResponseData:      logResponseData,
		logRequestBlacklist:  logRequestBlacklist,
		logResponseBlacklist: logResponseBlacklist,
		logBlacklist:         logBlacklist,
		debug:                debug,
	}

	client.common.client = client
	client.Authentication = (*AuthenticationService)(&client.common)
	client.Banks = (*BanksService)(&client.common)
	client.Bonuses = (*BonusesService)(&client.common)
	client.Countries = (*CountriesService)(&client.common)
	client.Consent = (*ConsentService)(&client.common)
	client.Currencies = (*CurrenciesService)(&client.common)
	client.Devices = (*DevicesService)(&client.common)
	client.Games = (*GamesService)(&client.common)
	client.KYC = (*KYCService)(&client.common)
	client.IPWhois = (*IPWhoisService)(&client.common)
	client.Languages = (*LanguagesService)(&client.common)
	client.Payments = (*PaymentsService)(&client.common)
	client.RealityCheck = (*RealityCheckService)(&client.common)
	client.ResponsibleGaming = (*ResponsibleGamingService)(&client.common)
	client.Roles = (*RolesService)(&client.common)
	client.SecurityQuestions = (*SecurityQuestionsService)(&client.common)
	client.User = (*UserService)(&client.common)
	client.Validation = (*ValidationService)(&client.common)
	client.Wallet = (*WalletService)(&client.common)

	return
}

func (c IGCClient) doLog(endpoint string, blacklist []string) bool {
	for _, blacklisted := range blacklist {
		if strings.HasPrefix(strings.ToLower(endpoint), strings.ToLower(blacklisted)) {
			return false
		}
	}
	return true
}

func (c IGCClient) apiReq(method, endpoint string, params *url.Values, body interface{}, data interface{}, headers *map[string]string, log logger.Logger) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	logInfo := make(map[string]interface{})

	if c.logRequestBody && body != nil && c.doLog(endpoint, c.logRequestBlacklist) {
		logInfo["request"] = body
	}

	logInfo["method"] = method

	req, err := http.NewRequest(method, c.baseURL+endpoint, b)
	if err != nil {
		return err
	}

	if headers != nil {
		for k, v := range *headers {
			if v != "" {
				req.Header.Add(k, v)
				logInfo[k] = v
			}
		}
	}

	if b != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Accept", "application/json")

	query := endpoint[1:] //don't log the first '/'
	logInfo["query"] = query

	if params != nil {
		req.URL.RawQuery = params.Encode()
		logInfo["params"] = params.Encode()
	}

	startTime := time.Now()

	resp, e := c.HTTPClient.Do(req)
	if e != nil {
		logInfo["error"] = e
		log.Info(fmt.Sprintf("failed to make request to igc endpoint %s", query), logInfo)
		return e
	}

	duration := time.Since(startTime)

	logInfo["duration"] = durationToMilliseconds(duration)

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()

	err = json.Unmarshal([]byte(s), data)

	if c.logResponseData && c.doLog(endpoint, c.logResponseBlacklist) {
		logInfo["response"] = data
	}

	if log != nil && (c.doLog(endpoint, c.logBlacklist) || c.debug) {
		log.Info(query+" request", logInfo)
	}

	if err != nil && log != nil {
		logFields := make(map[string]interface{})
		logFields["error"] = err
		logFields["response"] = s
		log.Info(fmt.Sprintf("failed to parse response from igc endpoint %s", query), logFields)
	}

	return err
}

func durationToMilliseconds(duration time.Duration) float32 {
	return float32(duration.Nanoseconds()/1000) / 1000
}
