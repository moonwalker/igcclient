package igcclient

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/moonwalker/igcclient/models"
	"github.com/moonwalker/logger"
)

type BonusesService service

// Gets the user current bonuses AuthenticationToken is required.
func (s *BonusesService) Bonuses(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses", nil, nil, &response, &headers, log)
	return
}

// Gets the user current bonuses X-Api-Key is required.
func (s *BonusesService) BonusesByUserID(userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	i := strconv.FormatInt(userID, 10)
	q := url.Values{}
	q.Add("userid", i)
	err = s.client.apiReq(http.MethodPost, "/bonuses", &q, nil, &response, &headers, log)
	return
}

// Gets all user bonuses
func (s *BonusesService) History(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfFrontEndUserBonusObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/history", nil, nil, &response, &headers, log)
	return
}

// Gets all user bonuses - including claimed free spins
func (s *BonusesService) FullHistory(headers map[string]string, log logger.Logger) (response models.OperationResponseOfFrontAllUserBonuses, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/fullhistory", nil, nil, &response, &headers, log)
	return
}

// Credit bonus to user AuthenticationToken is required
func (s *BonusesService) Credit(promoCode string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfBonusDetails, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/credit/"+url.QueryEscape(promoCode), nil, nil, &response, &headers, log)
	return
}

// Credit bonus to user X-Api-Key is required
func (s *BonusesService) CreditByUserID(promoCode string, userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfBonusDetails, err error) {
	i := strconv.FormatInt(userID, 10)
	q := url.Values{}
	q.Add("userid", i)
	err = s.client.apiReq(http.MethodPost, "/bonuses/credit/"+url.QueryEscape(promoCode), &q, nil, &response, &headers, log)
	return
}

// Get a list of bonuses that are public to the users (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses')
func (s *BonusesService) GetPublicBonuses(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/getpublicbonuses", nil, nil, &response, &headers, log)
	return
}

// Get a list of bonus types that are available
func (s *BonusesService) GetBonusTypes(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicBonusTypeObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/getbonustypes", nil, nil, &response, &headers, log)
	return
}

// Get a list of bonuses that are public and can be claimed by the logged in user (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses') AuthenticationToken header is required.
func (s *BonusesService) GetAvailable(deviceTypeID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicBonusesObject, err error) {
	deviceID := strconv.FormatInt(deviceTypeID, 10)
	err = s.client.apiReq(http.MethodPost, "/bonuses/getavailable/"+url.QueryEscape(deviceID), nil, nil, &response, &headers, log)
	return
}

// Get a list of bonuses that are public and can be claimed by the logged in user (Bonuses that are 'IncludedInlist' and not 'Manual Bonuses') X-Api-Key header is required.
func (s *BonusesService) GetAvailableByUserID(deviceTypeID int64, userID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicBonusesObject, err error) {
	i := strconv.FormatInt(userID, 10)
	deviceID := strconv.FormatInt(deviceTypeID, 10)
	q := url.Values{}
	q.Add("userid", i)
	err = s.client.apiReq(http.MethodPost, "/bonuses/getavailable/"+url.QueryEscape(deviceID), &q, nil, &response, &headers, log)
	return
}

// Forfeit user bonus
func (s *BonusesService) Delete(userBonusID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(userBonusID, 10)
	err = s.client.apiReq(http.MethodPost, "/bonuses/delete/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Return the bonus details of a bonus code
func (s *BonusesService) BonusCodeDetails(bonusCode string, headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicBonusesObject, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/bonuscodedetails/"+url.QueryEscape(bonusCode), nil, nil, &response, &headers, log)
	return
}

// Add a Promo Code to a Bonus which can be claimed only by the logged in user
func (s *BonusesService) PromoCodesAdd(bonusID int64, promoCode string, promoValidity int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	bonus := strconv.FormatInt(bonusID, 10)
	validity := strconv.FormatInt(promoValidity, 10)
	err = s.client.apiReq(http.MethodPost, "/bonuses/promocodes/add/"+url.QueryEscape(bonus)+"/"+url.QueryEscape(promoCode)+"/"+url.QueryEscape(validity), nil, nil, &response, &headers, log)
	return
}

// Return a list of unclaimed Promo Codes assigned to this user
func (s *BonusesService) PromoCodesPending(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPublicPromoCodeObj, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/promocodes/pending", nil, nil, &response, &headers, log)
	return
}

// User rejects promo code
func (s *BonusesService) PromoCodesReject(promoID int64, headers map[string]string, log logger.Logger) (response models.OperationResponseOfBoolean, err error) {
	i := strconv.FormatInt(promoID, 10)
	err = s.client.apiReq(http.MethodPost, "/bonuses/promocodes/reject/"+url.QueryEscape(i), nil, nil, &response, &headers, log)
	return
}

// Return a list of PromoCode Statuses
func (s *BonusesService) PromoCodesStatuses(headers map[string]string, log logger.Logger) (response models.OperationResponseOfListOfPromoStatuses, err error) {
	err = s.client.apiReq(http.MethodPost, "/bonuses/promocodes/statuses", nil, nil, &response, &headers, log)
	return
}
