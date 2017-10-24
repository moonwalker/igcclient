package models

type PaymentMetaDataModel struct {
	DtoName *string  `json:"dto-name,omitempty"`
	Fields  *[]Field `json:"fields,omitempty"`
}
