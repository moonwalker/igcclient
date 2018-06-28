package igcclient

import (
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
)

type GamesService service

// Gets a list available of Games supported by the system.
func (s *GamesService) Games(data models.GameFilterModel, headers map[string]string) (response models.OperationResponseOfIEnumerableOfGameFront, err error) {
	err = s.client.apiPost("/games", &data, &response, &headers)
	return
}

// Gets the MD5 hash for the requested games list. This can be used to cache games request on the client.
func (s *GamesService) MD5(data models.GameFilterModel, headers map[string]string) (response models.OperationResponseOfString, err error) {
	err = s.client.apiPost("/games/md5", &data, &response, &headers)
	return
}

// Gets all the Jackpots for the requested Currency
func (s *GamesService) Jackpots(currency string, headers map[string]string) (response models.OperationResponseOfIEnumerableOfJackpot, err error) {
	err = s.client.apiPost("/games/jackpots/"+url.QueryEscape(currency), nil, &response, &headers)
	return
}

// Get a single Game by Game ID
func (s *GamesService) GameByID(gameID int64, headers map[string]string) (response models.OperationResponseOfGameFront, err error) {
	id := strconv.FormatInt(gameID, 10)
	err = s.client.apiPost("/games/"+url.QueryEscape(id), nil, &response, &headers)
	return
}

// Get the URL of the Game to be used in the iFrame or to redirect to
func (s *GamesService) URL(gameID int64, body models.GameURLModel, headers map[string]string) (response models.OperationResponseOfString, err error) {
	id := strconv.FormatInt(gameID, 10)
	err = s.client.apiPost("/games/url/"+url.QueryEscape(id), &body, &response, &headers)
	return
}

// Get all last played games
func (s *GamesService) LastPlayed(maxResults int64, body models.GameUserInteractionDataModel, headers map[string]string) (response models.OperationResponseOfIEnumerableOfLastPlayed, err error) {
	max := strconv.FormatInt(maxResults, 10)
	err = s.client.apiPost("/games/lastplayed?maxresults="+url.QueryEscape(max), &body, &response, &headers)
	return
}

// Get all Game Categories by language alphaCode2
func (s *GamesService) Categories(alphaCode2 string, headers map[string]string) (response models.OperationResponseOfIEnumerableOfCategory, err error) {
	err = s.client.apiPost("/games/categories/"+url.QueryEscape(alphaCode2), nil, &response, &headers)
	return
}

// Get details of all the games.
func (s *GamesService) Details(headers map[string]string) (response models.OperationResponseOfIEnumerableOfGameDetails, err error) {
	err = s.client.apiPost("/games/details", nil, &response, &headers)
	return
}

// Get all recent game winners.
func (s *GamesService) RecentWinners(body models.RecentWinnersV2RequestModel, headers map[string]string) (response models.OperationResponseOfIEnumerableOfRecentWinnersResponseModel, err error) {
	err = s.client.apiPost("/v2/games/recentwinners", &body, &response, &headers)
	return
}

// Get a list of all the Vendors
func (s *GamesService) Vendors(enabledOnly bool, headers map[string]string) (response models.OperationResponseOfIEnumerableOfVendor, err error) {
	eoStr := strconv.FormatBool(enabledOnly)
	err = s.client.apiPost("/games/vendors?enabledonly="+url.QueryEscape(eoStr), nil, &response, &headers)
	return
}

func (s *GamesService) AffiliateGameDetails(headers map[string]string) (response models.OperationalResponseOfIEnumerableOfGameDetailsAffiliates, err error) {
	err = s.client.apiPost("/games/affiliategamedetails", nil, &response, &headers)
	return
}

func (s *GamesService) Transactions(body models.GameTransactionsSearchModel, headers map[string]string) (response models.OperationResponseOfIEnumerableOfGameTransactionModel, err error) {
	err = s.client.apiPost("/games/transactions", &body, &response, &headers)
	return
}
