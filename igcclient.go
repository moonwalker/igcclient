package igcclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/moonwalker/logger"
)

type IGCClient struct {
	HttpClient *http.Client
	baseUrl    string

	common            service
	Authentication    *AuthenticationService
	Banks             *BanksService
	Bonuses           *BonusesService
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

	log *logger.Logger
}

type service struct {
	client *IGCClient
}

func NewIGCClient(baseUrl string, log *logger.Logger) (client *IGCClient, err error) {
	if baseUrl == "" {
		err = errors.New("baseUrl can not be empty")
		return
	}
	client = &IGCClient{
		HttpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		log:        log,
	}

	client.common.client = client
	client.Authentication = (*AuthenticationService)(&client.common)
	client.Banks = (*BanksService)(&client.common)
	client.Bonuses = (*BonusesService)(&client.common)
	client.Countries = (*CountriesService)(&client.common)
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

func (c IGCClient) apiPost(endpoint string, body interface{}, data interface{}, xApiKey *string, authToken *string) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	logInfo := make(map[string]interface{})

	req, err := http.NewRequest("POST", c.baseUrl+endpoint, b)
	if err != nil {
		return err
	}

	if authToken != nil && *authToken != "" {
		req.Header.Add("AuthenticationToken", *authToken)
		logInfo["AuthenticationToken"] = *authToken
	}

	if xApiKey != nil && *xApiKey != "" {
		req.Header.Add("X-Api-Key", *xApiKey)
		logInfo["X-Api-Key"] = true
	}

	if b != nil {
		req.Header.Add("Content-Type", "application/json")
		logInfo["Data"] = data
	}

	resp, e := c.HttpClient.Do(req)
	if e != nil {
		return e
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()

	logInfo["Response"] = s
	logInfo["StatusCode"] = resp.StatusCode
	logInfo["URL"] = c.baseUrl + endpoint

	if c.log != nil {
		(*c.log).Info("Request to IGC api", logInfo)
	}

	return json.Unmarshal([]byte(s), data)
}
