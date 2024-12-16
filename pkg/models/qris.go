package models

type QRIS struct {
	Version               Data
	Category              Data
	Acquirer              Acquirer
	Switching             Switching
	MerchantCategoryCode  Data
	CurrencyCode          Data
	PaymentAmount         Data
	PaymentFeeCategory    Data
	PaymentFee            Data
	CountryCode           Data
	MerchantName          Data
	MerchantCity          Data
	MerchantPostalCode    Data
	AdditionalInformation Data
	CRCCode               Data
}
