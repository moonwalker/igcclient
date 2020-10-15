package igcclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/moonwalker/logger"

	igcerr "github.com/moonwalker/igcclient/errors"
	"github.com/moonwalker/igcclient/models"
)

const (
	timeout                   = 30 * time.Second
	DefaultLogMaxResponseSize = 10000
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
	logMaxResponseSize   int64

	debug bool

	invalidAuthCallback *func(string)
}

type service struct {
	client *IGCClient
}

type Config struct {
	BaseURL              string
	LogRequestBody       bool
	LogResponseData      bool
	LogRequestBlacklist  []string
	LogResponseBlacklist []string
	LogBlacklist         []string
	LogMaxResponseSize   int64
	Debug                bool
	InvalidAuthCallback  *func(string)
}

func NewIGCClient(cfg Config) (client *IGCClient, err error) {
	if cfg.BaseURL == "" {
		err = errors.New("base url can not be empty")
		return
	}
	client = &IGCClient{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		baseURL:              cfg.BaseURL,
		logRequestBody:       cfg.LogRequestBody,
		logResponseData:      cfg.LogResponseData,
		logRequestBlacklist:  cfg.LogRequestBlacklist,
		logResponseBlacklist: cfg.LogResponseBlacklist,
		logBlacklist:         cfg.LogBlacklist,
		debug:                cfg.Debug,
		invalidAuthCallback:  cfg.InvalidAuthCallback,
		logMaxResponseSize:   cfg.LogMaxResponseSize,
	}

	if client.logMaxResponseSize == 0 {
		client.logMaxResponseSize = DefaultLogMaxResponseSize
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
			if k == "X-Api-Key" {
				if v != "" {
					req.Header.Add(k, v)
					logInfo[k] = c.obfuscate(v)
				}
			} else {
				if v != "" {
					req.Header.Add(k, v)
					logInfo[k] = v
				}
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
		return e
	}

	logInfo["duration"] = time.Since(startTime).Milliseconds()

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.Bytes()
	ls := int64(len(s))

	if headers != nil && (*headers)["AuthenticationToken"] != "" {
		authToken := (*headers)["AuthenticationToken"]
		c.checkForAuthError(s, authToken)
	}

	err = json.Unmarshal(s, data)

	if c.logResponseData && c.doLog(endpoint, c.logResponseBlacklist) {
		if ls < c.logMaxResponseSize {
			logInfo["response"] = data
		} else {
			logInfo["response"] = fmt.Sprintf("response data is too large to log (>=%d byte)", c.logMaxResponseSize)
		}
	}

	if log != nil && (c.doLog(endpoint, c.logBlacklist) || c.debug) {
		log.Info(query+" request", logInfo)
	}

	if err != nil && log != nil {
		logFields := make(map[string]interface{})
		logFields["error"] = err
		var response []byte
		response, err = base64.StdEncoding.DecodeString(string(s))
		if err != nil {
			response = s
		}
		if ls < c.logMaxResponseSize {
			logFields["response"] = response
		} else {
			logFields["response"] = response[:c.logMaxResponseSize]
		}
		log.Error(fmt.Sprintf("failed to parse response from igc endpoint %s", query), logFields)
	}

	return err
}

func (c IGCClient) checkForAuthError(data []byte, authToken string) {
	if c.invalidAuthCallback != nil {
		d := &models.OperationResponse{}
		if d.Errors != nil {
			for _, e := range *d.Errors {
				if e.ErrorCodeID != nil && *e.ErrorCodeID == igcerr.INVALID_AUTHENTICATION_TOKEN {
					// User is not logged in
					(*c.invalidAuthCallback)(authToken)
				}
			}
		}
	}
}

func (c IGCClient) obfuscate(val string) string {
	if val == "" || len(val) < 3 {
		return val
	}
	first3 := val[0:3]
	last3 := val[len(val)-3:]
	return fmt.Sprintf("%sxxxx%s", first3, last3)
}
