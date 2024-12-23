package usecases

type QRISTags struct {
	Version               string
	Category              string
	Acquirer              string
	AcquirerBankTransfer  string
	Switching             string
	MerchantCategoryCode  string
	CurrencyCode          string
	PaymentAmount         string
	PaymentFeeCategory    string
	PaymentFeeFixed       string
	PaymentFeePercent     string
	CountryCode           string
	MerchantName          string
	MerchantCity          string
	MerchantPostalCode    string
	AdditionalInformation string
	CRCCode               string
}

type QRISCategoryContents struct {
	Static  string
	Dynamic string
}

type QRISPaymentFeeCategoryContents struct {
	Fixed   string
	Percent string
}
