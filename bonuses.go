package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type BonusesService service

// Gets the user current bonuses AuthenticationToken is required.
func (s *BonusesService) Bonuses(authToken string) (response OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	err = s.client.apiPost("/bonuses", nil, &response, nil, &authToken)
	return
}

// Gets the user current bonuses X-Api-Key is required.
func (s *BonusesService) BonusesByUserID(userID int64, xAPIKey string) (response OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	i := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/bonuses?userid="+url.QueryEscape(i), nil, &response, &xAPIKey, nil)
	return
}

// Gets all user bonuses
func (s *BonusesService) History(authToken string) (response OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	err = s.client.apiPost("/bonuses/history", nil, &response, nil, &authToken)
	return
}

// Gets all user bonuses - including claimed free spins
func (s *BonusesService) FullHistory(authToken string) (response OperationResponseOfFrontAllUserBonuses, err error) {
	err = s.client.apiPost("/bonuses/fullhistory", nil, &response, nil, &authToken)
	return
}

// Credit bonus to user AuthenticationToken is required
func (s *BonusesService) Credit(promoCode string, authToken string) (response OperationResponseOfListOfBonusDetails, err error) {
	err = s.client.apiPost("/bonuses/credit/"+url.QueryEscape(promoCode), nil, &response, nil, &authToken)
	return
}

// Credit bonus to user X-Api-Key is required
func (s *BonusesService) CreditByUserID(promoCode string, userID int64, xAPIKey string) (response OperationResponseOfListOfBonusDetails, err error) {
	i := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/bonuses/credit/"+url.QueryEscape(promoCode)+"?userid="+url.QueryEscape(i), nil, &response, &xAPIKey, nil)
	return
}

// Get a list of bonuses that are public to the users (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses')
func (s *BonusesService) GetPublicBonuses(authToken string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiPost("/bonuses/getpublicbonuses", nil, &response, nil, &authToken)
	return
}

// Get a list of bonus types that are available
func (s *BonusesService) GetBonusTypes(authToken string) (response OperationResponseOfListOfPublicBonusTypeObject, err error) {
	err = s.client.apiPost("/bonuses/getbonustypes", nil, &response, nil, &authToken)
	return
}

// Get a list of bonuses that are public and can be claimed by the logged in user (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses') AuthenticationToken header is required.
func (s *BonusesService) GetAvailable(deviceTypeID string, xAPIKey string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiPost("/bonuses/getavailable/"+url.QueryEscape(deviceTypeID), nil, &response, &xAPIKey, nil)
	return
}

// Get a list of bonuses that are public and can be claimed by the logged in user (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses') X-Api-Key header is required.
func (s *BonusesService) GetAvailableByUserID(deviceTypeID string, userID int64, authToken string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	i := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/bonuses/getavailable/"+url.QueryEscape(deviceTypeID)+"?userid="+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Forfeit user bonus
func (s *BonusesService) Delete(userBonusID int64, authToken string) (response OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(userBonusID, 10)
	err = s.client.apiPost("/bonuses/delete/"+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Return the bonus details of a bonus code
func (s *BonusesService) BonusCodeDetails(bonusCode string, authToken string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiPost("/bonuses/bonuscodedetails/"+url.QueryEscape(bonusCode), nil, &response, nil, &authToken)
	return
}

// Add a Promo Code to a Bonus which can be claimed only by the logged in user
func (s *BonusesService) PromoCodesAdd(bonusID int64, promoCode string, promoValidity int64, authToken string) (response OperationResponseOfBoolean, err error) {
	bonus := strconv.FormatInt(bonusID, 10)
	validity := strconv.FormatInt(promoValidity, 10)
	err = s.client.apiPost("/bonuses/promocodes/add/"+url.QueryEscape(bonus)+"/"+url.QueryEscape(promoCode)+"/"+url.QueryEscape(validity), nil, &response, nil, &authToken)
	return
}

// Return a list of unclaimed Promo Codes assigned to this user
func (s *BonusesService) PromoCodesPending(authToken string) (response OperationResponseOfListOfPublicPromoCodeObj, err error) {
	err = s.client.apiPost("/bonuses/promocodes/pending", nil, &response, nil, &authToken)
	return
}

// User rejects promo code
func (s *BonusesService) PromoCodesReject(promoID int64, authToken string) (response OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(promoID, 10)
	err = s.client.apiPost("/bonuses/promocode/reject/"+url.QueryEscape(i), nil, &response, nil, &authToken)
	return
}

// Return a list of PromoCode Statuses
func (s *BonusesService) PromoCodesStatuses(authToken string) (response OperationResponseOfListOfPromoStatuses, err error) {
	err = s.client.apiPost("/bonuses/promocodes/statuses", nil, &response, nil, &authToken)
	return
}
