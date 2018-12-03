package igcclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type GamesService service

// Gets a list available of Games supported by the system.
func (s *GamesService) Games(data models.GameFilterModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfGameFront, err error) {
	err = s.client.apiReq(http.MethodPost, "/games", nil, &data, &response, &headers, log)
	return
}

// Gets the MD5 hash for the requested games list. This can be used to cache games request on the client.
func (s *GamesService) MD5(data models.GameFilterModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfString, err error) {
	err = s.client.apiReq(http.MethodPost, "/games/md5", nil, &data, &response, &headers, log)
	return
}

// Gets all the Jackpots for the requested Currency
func (s *GamesService) Jackpots(currency string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfJackpot, err error) {
	err = s.client.apiReq(http.MethodPost, "/games/jackpots/"+url.QueryEscape(currency), nil, nil, &response, &headers, log)
	return
}

// Get a single Game by Game ID
func (s *GamesService) GameByID(gameID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfGameFront, err error) {
	id := strconv.FormatInt(gameID, 10)
	err = s.client.apiReq(http.MethodPost, "/games/"+url.QueryEscape(id), nil, nil, &response, &headers, log)
	return
}

// Get the URL of the Game to be used in the iFrame or to redirect to
func (s *GamesService) URL(gameID int64, body models.GameURLModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfString, err error) {
	id := strconv.FormatInt(gameID, 10)
	err = s.client.apiReq(http.MethodPost, "/games/url/"+url.QueryEscape(id), nil, &body, &response, &headers, log)
	return
}

// Get all last played games
func (s *GamesService) LastPlayed(maxResults int64, body models.GameUserInteractionDataModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfLastPlayed, err error) {
	q := url.Values{}
	q.Add("maxresults", fmt.Sprintf("%d", maxResults))
	err = s.client.apiReq(http.MethodPost, "/games/lastplayed", &q, &body, &response, &headers, log)
	return
}

// Get all Game Categories by language alphaCode2
func (s *GamesService) Categories(alphaCode2 string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfCategory, err error) {
	err = s.client.apiReq(http.MethodPost, "/games/categories/"+url.QueryEscape(alphaCode2), nil, nil, &response, &headers, log)
	return
}

// Get details of all the games.
func (s *GamesService) Details(headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfGameDetails, err error) {
	err = s.client.apiReq(http.MethodPost, "/games/details", nil, nil, &response, &headers, log)
	return
}

// Get all recent game winners.
func (s *GamesService) RecentWinners(body models.RecentWinnersV2RequestModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfRecentWinnersResponseModel, err error) {
	err = s.client.apiReq(http.MethodPost, "/v2/games/recentwinners", nil, &body, &response, &headers, log)
	return
}

// Get a list of all the Vendors
func (s *GamesService) Vendors(enabledOnly bool, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfVendor, err error) {
	eoStr := strconv.FormatBool(enabledOnly)
	q := url.Values{}
	q.Add("enabledonly", eoStr)
	err = s.client.apiReq(http.MethodPost, "/games/vendors", nil, nil, &response, &headers, log)
	return
}

func (s *GamesService) AffiliateGameDetails(headers map[string]string, log logger.Logger) (response models.OperationalResponseOfIEnumerableOfGameDetailsAffiliates, err error) {
	err = s.client.apiReq(http.MethodPost, "/games/affiliategamedetails", nil, nil, &response, &headers, log)
	return
}

func (s *GamesService) Transactions(body models.GameTransactionsSearchModel, headers map[string]string, log logger.Logger) (response models.OperationResponseOfIEnumerableOfGameTransactionModel, err error) {
	err = s.client.apiReq(http.MethodPost, "/games/transactions", nil, &body, &response, &headers, log)
	return
}
