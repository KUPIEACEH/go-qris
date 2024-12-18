package usecases

type QRISTags struct {
	VersionTag               string
	CategoryTag              string
	AcquirerTag              string
	AcquirerBankTransferTag  string
	SwitchingTag             string
	MerchantCategoryCodeTag  string
	CurrencyCodeTag          string
	PaymentAmountTag         string
	PaymentFeeCategoryTag    string
	PaymentFeeFixedTag       string
	PaymentFeePercentTag     string
	CountryCodeTag           string
	MerchantNameTag          string
	MerchantCityTag          string
	MerchantPostalCodeTag    string
	AdditionalInformationTag string
	CRCCodeTag               string
}

type QRISCategoryContents struct {
	Static  string
	Dynamic string
}

type QRISPaymentFeeCategoryContents struct {
	Fixed   string
	Percent string
}
