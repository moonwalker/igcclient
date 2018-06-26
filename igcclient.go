package igcclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/moonwalker/logger"
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
	Languages         *LanguagesService
	Payments          *PaymentsService
	RealityCheck      *RealityCheckService
	ResponsibleGaming *ResponsibleGamingService
	Roles             *RolesService
	SecurityQuestions *SecurityQuestionsService
	User              *UserService
	UserVerification  *UserVerificationService
	Validation        *ValidationService
	Wallet            *WalletService

	log logger.Logger

	logRequestBody  bool
	logResponseData bool
}

type service struct {
	client *IGCClient
}

func NewIGCClient(baseURL string, log logger.Logger, logRequestBody bool, logResponseData bool) (client *IGCClient, err error) {
	if baseURL == "" {
		err = errors.New("base url can not be empty")
		return
	}
	client = &IGCClient{
		HTTPClient:      http.DefaultClient,
		baseURL:         baseURL,
		log:             log,
		logRequestBody:  logRequestBody,
		logResponseData: logResponseData,
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
	client.Languages = (*LanguagesService)(&client.common)
	client.Payments = (*PaymentsService)(&client.common)
	client.RealityCheck = (*RealityCheckService)(&client.common)
	client.ResponsibleGaming = (*ResponsibleGamingService)(&client.common)
	client.Roles = (*RolesService)(&client.common)
	client.SecurityQuestions = (*SecurityQuestionsService)(&client.common)
	client.User = (*UserService)(&client.common)
	client.UserVerification = (*UserVerificationService)(&client.common)
	client.Validation = (*ValidationService)(&client.common)
	client.Wallet = (*WalletService)(&client.common)

	return
}

func (c IGCClient) apiPost(endpoint string, body interface{}, data interface{}, headers *map[string]string) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	logRequest := make(map[string]interface{})
	logResponse := make(map[string]interface{})

	if c.logRequestBody {
		logData, err := json.Marshal(body)
		if err == nil {
			logRequest["Body"] = string(logData)
		}
	}

	req, err := http.NewRequest("POST", c.baseURL+endpoint, b)
	if err != nil {
		return err
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, *v)
			logRequest[k] = *v
		}
	}

	if b != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	logRequest["URL"] = c.baseURL + endpoint

	if c.log != nil {
		c.log.Info("IGC Request", logRequest)
	}

	resp, e := c.HTTPClient.Do(req)
	if e != nil {
		return e
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()

	if c.logResponseData {
		logResponse["Response"] = s
	}

	logResponse["StatusCode"] = resp.StatusCode
	logResponse["URL"] = c.baseURL + endpoint

	if c.log != nil {
		c.log.Info("IGC Response", logResponse)
	}

	return json.Unmarshal([]byte(s), data)
}
