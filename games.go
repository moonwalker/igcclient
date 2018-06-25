package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type GamesService service

// Gets a list available of Games supported by the system.
func (s *GamesService) Games(data GameFilterModel) (response OperationResponseOfIEnumerableOfGameFront, err error) {
	err = s.client.apiPost("/games", &data, &response, nil, nil)
	return
}

// Gets the MD5 hash for the requested games list. This can be used to cache games request on the client.
func (s *GamesService) MD5(data GameFilterModel) (response OperationResponseOfString, err error) {
	err = s.client.apiPost("/games/md5", &data, &response, nil, nil)
	return
}

// Gets all the Jackpots for the requested Currency
func (s *GamesService) Jackpots(currency string) (response OperationResponseOfIEnumerableOfJackpot, err error) {
	err = s.client.apiPost("/games/jackpots/"+url.QueryEscape(currency), nil, &response, nil, nil)
	return
}

// Get a single Game by Game ID
func (s *GamesService) GameByID(gameID int64) (response OperationResponseOfGameFront, err error) {
	id := strconv.FormatInt(gameID, 10)
	err = s.client.apiPost("/games/"+url.QueryEscape(id), nil, &response, nil, nil)
	return
}

// Get the URL of the Game to be used in the iFrame or to redirect to
func (s *GamesService) URL(gameID int64, body GameURLModel) (response OperationResponseOfString, err error) {
	id := strconv.FormatInt(gameID, 10)
	err = s.client.apiPost("/games/url/"+url.QueryEscape(id), &body, &response, nil, nil)
	return
}

// Get all last played games
func (s *GamesService) LastPlayed(maxResults int64, body GameUserInteractionDataModel, authToken string) (response OperationResponseOfIEnumerableOfLastPlayed, err error) {
	max := strconv.FormatInt(maxResults, 10)
	err = s.client.apiPost("/games/lastplayed?maxresults="+url.QueryEscape(max), &body, &response, nil, &authToken)
	return
}

// Get all Game Categories by language alphaCode2
func (s *GamesService) Categories(alphaCode2 string) (response OperationResponseOfIEnumerableOfCategory, err error) {
	err = s.client.apiPost("/games/categories/"+url.QueryEscape(alphaCode2), nil, &response, nil, nil)
	return
}

// Get details of all the games.
func (s *GamesService) Details() (response OperationResponseOfIEnumerableOfGameDetails, err error) {
	err = s.client.apiPost("/games/details", nil, &response, nil, nil)
	return
}

// Get all recent game winners.
func (s *GamesService) RecentWinners(body RecentWinnersV2RequestModel) (response OperationResponseOfIEnumerableOfRecentWinnersResponseModel, err error) {
	err = s.client.apiPost("/v2/games/recentwinners", &body, &response, nil, nil)
	return
}

// Get a list of all the Vendors
func (s *GamesService) Vendors(enabledOnly bool) (response OperationResponseOfIEnumerableOfVendor, err error) {
	eoStr := strconv.FormatBool(enabledOnly)
	err = s.client.apiPost("/games/vendors?enabledonly="+url.QueryEscape(eoStr), nil, &response, nil, nil)
	return
}

func (s *GamesService) AffiliateGameDetails() (response OperationalResponseOfIEnumerableOfGameDetailsAffiliates, err error) {
	err = s.client.apiPost("/games/affiliategamedetails", nil, &response, nil, nil)
	return
}

func (s *GamesService) Transactions(body GameTransactionsSearchModel, authToken string) (response OperationResponseOfIEnumerableOfGameTransactionModel, err error) {
	err = s.client.apiPost("/games/transactions", &body, &response, nil, &authToken)
	return
}
