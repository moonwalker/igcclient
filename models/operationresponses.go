package models

type OperationResponse struct {
	Success *bool        `json:"Success"`
	Errors  *[]ErrorCode `json:"Errors"`
}

type OperationResponseOfIEnumerableOfCategory struct {
	*OperationResponse
	Data *[]Category `json:"Data"`
}

type OperationResponseOfIEnumerableOfLastPlayed struct {
	*OperationResponse
	Data *[]LastPlayed `json:"Data"`
}

type OperationResponseOfIEnumerableOfGameFront struct {
	*OperationResponse
	Data *[]GameFront `json:"Data"`
}

type OperationResponseOfBoolean struct {
	*OperationResponse
	Data *bool `json:"Data"`
}

type OperationResponseOfNullableIfGuid struct {
	*OperationResponse
	Data *string `json:"Data"`
}

type OperationResponseOfString struct {
	*OperationResponse
	Data *string `json:"Data"`
}

type OperationResponseOfListOfBankObject struct {
	*OperationResponse
	Data *[]BankObject `json:"Data"`
}

type OperationResponseOfBankObject struct {
	*OperationResponse
	Data *BankObject `json:"Data"`
}

type OperationResponseOfListOfDBTokenIP struct {
	*OperationResponse
	Data *[]DBTokenIP `json:"Data"`
}

type OperationResponseOfSafeUserDetails struct {
	*OperationResponse
	Data *SafeUserDetails `json:"Data"`
}

type OperationResponseOfListOfLanguageObject struct {
	*OperationResponse
	Data *[]LanguageObject `json:"Data"`
}

type OperationResponseOfLanguageObject struct {
	*OperationResponse
	Data *LanguageObject `json:"Data"`
}

type OperationResponseOfListOfCountriesObject struct {
	*OperationResponse
	Data *[]CountriesObject `json:"Data"`
}

type OperationResponseOfCountriesObject struct {
	*OperationResponse
	Data *CountriesObject `json:"Data"`
}

type OperationResponseOfListOfCurrenciesObject struct {
	*OperationResponse
	Data *[]CurrenciesObject `json:"Data"`
}

type OperationResponseOfCurrenciesObject struct {
	*OperationResponse
	Data *CurrenciesObject `json:"Data"`
}

type OperationResponseOfListOfCountryVerificationTypeWhitelist struct {
	*OperationResponse
	Data *[]CountryVerificationTypeWhitelist `json:"Data"`
}

type OperationResponseOfDictionaryOfStringAndString struct {
	*OperationResponse
	Data *map[string]string `json:"Data"`
}

type OperationResponseOfGuid struct {
	*OperationResponse
	Data *string `json:"Data"`
}

type OperationResponseOfSecurityQuestion struct {
	*OperationResponse
	Data *SecurityQuestion `json:"Data"`
}

type OperationResponseOfIEnumerableOfRoleResponse struct {
	*OperationResponse
	Data *[]RolesResponse `json:"Data"`
}

type OperationResponseOfListOfDeviceTypeObject struct {
	*OperationResponse
	Data *[]DeviceTypeObject `json:"Data"`
}

type OperationResponseOfDeviceTypeObject struct {
	*OperationResponse
	Data *DeviceTypeObject `json:"Data"`
}

type OperationResponseOfDevicePlatformModel struct {
	*OperationResponse
	Data *DevicePlatformModel `json:"Data"`
}

type OperationResponseOfDecimal struct {
	*OperationResponse
	Data *float64 `json:"Data"`
}

type OperationResponseOfIEnumerableOfUserWithNoActivity struct {
	*OperationResponse
	Data *[]UserWithNoActivity `json:"Data"`
}

type OperationResponseOfInt32 struct {
	*OperationResponse
	Data *int32 `json:"Data"`
}

type OperationResponseOfListOfKYCLinkedObject struct {
	*OperationResponse
	Data *[]KYCLinkedObject `json:"Data"`
}

type OperationResponseOfIEnumerableOfLimitDurationObject struct {
	*OperationResponse
	Data *[]LimitDurationObject `json:"Data"`
}

type OperationResponseOfIEnumerableOfLimit struct {
	*OperationResponse
	Data *[]Limit `json:"Data"`
}

type OperationResponseOfIEnumerableOfUserLimit struct {
	*OperationResponse
	Data *[]UserLimit `json:"Data"`
}

type OperationResponseOfIEnumerableOfJackpot struct {
	*OperationResponse
	Data *[]Jackpot `json:"Data"`
}

type OperationResponseOfGameFront struct {
	*OperationResponse
	Data *GameFront `json:"Data"`
}

type OperationResponseOfIEnumerableOfGameDetails struct {
	*OperationResponse
	Data *[]GameDetails `json:"Data"`
}

type OperationResponseOfIEnumerableOfRecentWinnersResponseModel struct {
	*OperationResponse
	Data *[]RecentWinnersResponseModel `json:"Data"`
}

type OperationResponseOfIEnumerableOfVendor struct {
	*OperationResponse
	Data *[]Vendor `json:"Data"`
}

type OperationalResponseOfIEnumerableOfGameDetailsAffiliates struct {
	*OperationResponse
	Data *[]GameDetailsAffiliates `json:"Data"`
}

type OperationResponseOfIEnumerableOfPaymentMethod struct {
	*OperationResponse
	Data *[]PaymentMethod `json:"Data"`
}

type OperationResponseOfPaymentApiResponse struct {
	*OperationResponse
	Data *[]PaymentApiResponse `json:"Data"`
}

type OperationResponseOfPaymentResponse struct {
	*OperationResponse
	Data *PaymentResponse `json:"Data"`
}

type OperationResponseOfTransactionModel struct {
	*OperationResponse
	Data *TransactionModel `json:"Data"`
}

type OperationResponseOfPaymentMetaDataModel struct {
	*OperationResponse
	Data *PaymentMetaDataModel `json:"Data"`
}

type OperationResponseOfWalletResponse struct {
	*OperationResponse
	Data *WalletResponse `json:"Data"`
}

type OperationResponseOfOMUserWalletPublic struct {
	*OperationResponse
	Data *OMUserWalletPublic `json:"Data"`
}

type OperationResponseOfBonus struct {
	*OperationResponse
	Data *Bonus `json:"Data"`
}

type OperationResponseOfUserWalletPublic struct {
	*OperationResponse
	Data *UserWalletPublic `json:"Data"`
}

type OperationResponseOfIEnumerableOfWalletTransactionsModel struct {
	*OperationResponse
	Data *[]WalletTransactionsModel `json:"Data"`
}

type OperationResponseOfListOfFrontEndUserBonusObject struct {
	*OperationResponse
	Data *[]FrontEndUserBonusObject `json:"Data"`
}

type OperationResponseOfFrontAllUserBonuses struct {
	*OperationResponse
	Data *FrontAllUserBonuses `json:"Data"`
}

type OperationResponseOfListOfBonusDetails struct {
	*OperationResponse
	Data *[]BonusDetails `json:"Data"`
}

type OperationResponseOfListOfPublicBonusesObject struct {
	*OperationResponse
	Data *[]PublicBonusesObject `json:"Data"`
}

type OperationResponseOfListOfPublicBonusTypeObject struct {
	*OperationResponse
	Data *[]PublicBonusTypeObject `json:"Data"`
}

type OperationResponseOfListOfPublicPromoCodeObj struct {
	*OperationResponse
	Data *[]PublicPromoCodeObj `json:"Data"`
}

type OperationResponseOfListOfPromoStatuses struct {
	*OperationResponse
	Data *[]PromoStatuses `json:"Data"`
}

type OperationResponseOfRealityCheckUserSetting struct {
	*OperationResponse
	Data *RealityCheckUserSetting `json:"Data"`
}

type OperationResponseOfListOfRealityCheckOption struct {
	*OperationResponse
	Data *[]RealityCheckOption `json:"Data"`
}

type OperationResponseOfListOfSelfExclusionCategory struct {
	*OperationResponse
	Data *[]SelfExclusionCategory `json:"Data"`
}

type OperationResponseOfListOfSelfExclusionCategoryDuration struct {
	*OperationResponse
	Data *[]SelfExclusionCategoryDuration `json:"Data"`
}

type OperationResponseOfObject struct {
	*OperationResponse
	Data *string `json:"Data"`
}

type OperationResponseOfListOfSecurityQuestionModel struct {
	*OperationResponse
	Data *[]SecurityQuestionModel `json:"Data"`
}

type OperationResponseOfIEnumerableOfUserLimitResponse struct {
	*OperationResponse
	Data *[]UserLimitResponse `json:"Data"`
}

type OperationResponseOfListOfPublicConsentModel struct {
	*OperationResponse
	Data *[]PublicConsentModel `json:"Data"`
}

type OperationResponseOfPublicUserConsentsListModel struct {
	*OperationResponse
	Data *PublicUserConsentsListModel `json:"Data"`
}
