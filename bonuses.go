package igcclient

import (
	"net/url"
	"strconv"

	. "github.com/moonwalker/igcclient/models"
)

type BonusesService service

// Gets the user current bonuses AuthenticationToken is required.
func (s *BonusesService) Bonuses(headers map[string]string) (response OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	err = s.client.apiPost("/bonuses", nil, &response, &headers)
	return
}

// Gets the user current bonuses X-Api-Key is required.
func (s *BonusesService) BonusesByUserID(userID int64, headers map[string]string) (response OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	i := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/bonuses?userid="+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Gets all user bonuses
func (s *BonusesService) History(headers map[string]string) (response OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	err = s.client.apiPost("/bonuses/history", nil, &response, &headers)
	return
}

// Gets all user bonuses - including claimed free spins
func (s *BonusesService) FullHistory(headers map[string]string) (response OperationResponseOfFrontAllUserBonuses, err error) {
	err = s.client.apiPost("/bonuses/fullhistory", nil, &response, &headers)
	return
}

// Credit bonus to user AuthenticationToken is required
func (s *BonusesService) Credit(promoCode string, headers map[string]string) (response OperationResponseOfListOfBonusDetails, err error) {
	err = s.client.apiPost("/bonuses/credit/"+url.QueryEscape(promoCode), nil, &response, &headers)
	return
}

// Credit bonus to user X-Api-Key is required
func (s *BonusesService) CreditByUserID(promoCode string, userID int64, headers map[string]string) (response OperationResponseOfListOfBonusDetails, err error) {
	i := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/bonuses/credit/"+url.QueryEscape(promoCode)+"?userid="+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Get a list of bonuses that are public to the users (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses')
func (s *BonusesService) GetPublicBonuses(headers map[string]string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiPost("/bonuses/getpublicbonuses", nil, &response, &headers)
	return
}

// Get a list of bonus types that are available
func (s *BonusesService) GetBonusTypes(headers map[string]string) (response OperationResponseOfListOfPublicBonusTypeObject, err error) {
	err = s.client.apiPost("/bonuses/getbonustypes", nil, &response, &headers)
	return
}

// Get a list of bonuses that are public and can be claimed by the logged in user (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses') AuthenticationToken header is required.
func (s *BonusesService) GetAvailable(deviceTypeID string, headers map[string]string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiPost("/bonuses/getavailable/"+url.QueryEscape(deviceTypeID), nil, &response, &headers)
	return
}

// Get a list of bonuses that are public and can be claimed by the logged in user (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses') X-Api-Key header is required.
func (s *BonusesService) GetAvailableByUserID(deviceTypeID string, userID int64, headers map[string]string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	i := strconv.FormatInt(userID, 10)
	err = s.client.apiPost("/bonuses/getavailable/"+url.QueryEscape(deviceTypeID)+"?userid="+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Forfeit user bonus
func (s *BonusesService) Delete(userBonusID int64, headers map[string]string) (response OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(userBonusID, 10)
	err = s.client.apiPost("/bonuses/delete/"+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Return the bonus details of a bonus code
func (s *BonusesService) BonusCodeDetails(bonusCode string, headers map[string]string) (response OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiPost("/bonuses/bonuscodedetails/"+url.QueryEscape(bonusCode), nil, &response, &headers)
	return
}

// Add a Promo Code to a Bonus which can be claimed only by the logged in user
func (s *BonusesService) PromoCodesAdd(bonusID int64, promoCode string, promoValidity int64, headers map[string]string) (response OperationResponseOfBoolean, err error) {
	bonus := strconv.FormatInt(bonusID, 10)
	validity := strconv.FormatInt(promoValidity, 10)
	err = s.client.apiPost("/bonuses/promocodes/add/"+url.QueryEscape(bonus)+"/"+url.QueryEscape(promoCode)+"/"+url.QueryEscape(validity), nil, &response, &headers)
	return
}

// Return a list of unclaimed Promo Codes assigned to this user
func (s *BonusesService) PromoCodesPending(headers map[string]string) (response OperationResponseOfListOfPublicPromoCodeObj, err error) {
	err = s.client.apiPost("/bonuses/promocodes/pending", nil, &response, &headers)
	return
}

// User rejects promo code
func (s *BonusesService) PromoCodesReject(promoID int64, headers map[string]string) (response OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(promoID, 10)
	err = s.client.apiPost("/bonuses/promocode/reject/"+url.QueryEscape(i), nil, &response, &headers)
	return
}

// Return a list of PromoCode Statuses
func (s *BonusesService) PromoCodesStatuses(headers map[string]string) (response OperationResponseOfListOfPromoStatuses, err error) {
	err = s.client.apiPost("/bonuses/promocodes/statuses", nil, &response, &headers)
	return
}
