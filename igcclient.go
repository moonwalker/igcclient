package igcclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
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
	User              *UserService
	UserVerification  *UserVerificationService
	Validation        *ValidationService
	Wallet            *WalletService
}

type service struct {
	client *IGCClient
}

func NewIGCClient(baseUrl string) (client *IGCClient, err error) {
	if baseUrl == "" {
		err = errors.New("baseUrl can not be empty")
		return
	}
	client = &IGCClient{
		HttpClient: http.DefaultClient,
		baseUrl:    baseUrl,
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
	client.User = (*UserService)(&client.common)
	client.UserVerification = (*UserVerificationService)(&client.common)
	client.Validation = (*ValidationService)(&client.common)
	client.Wallet = (*WalletService)(&client.common)

	return
}

func (c IGCClient) apiPost(endpoint string, body interface{}, data interface{}, xApiKey *string, authToken *string) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest("POST", c.baseUrl+endpoint, b)
	if err != nil {
		return err
	}

	if authToken != nil && *authToken != "" {
		req.Header.Add("AuthenticationToken", *authToken)
	}

	if xApiKey != nil && *xApiKey != "" {
		req.Header.Add("X-Api-Key", *xApiKey)
	}

	if b != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, e := c.HttpClient.Do(req)
	if e != nil {
		return e
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}
