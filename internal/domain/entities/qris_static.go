package entities

type QRISStatic struct {
	Version               ExtractData `json:"version"`
	Category              ExtractData `json:"category"`
	Acquirer              Acquirer    `json:"acquirer"`
	Switching             Switching   `json:"switching"`
	MerchantCategoryCode  ExtractData `json:"merchant_category_code"`
	CurrencyCode          ExtractData `json:"currency_code"`
	CountryCode           ExtractData `json:"country_code"`
	MerchantName          ExtractData `json:"merchant_name"`
	MerchantCity          ExtractData `json:"merchant_city"`
	MerchantPostalCode    ExtractData `json:"merchant_postal_code"`
	AdditionalInformation ExtractData `json:"additional_information"`
	CRCCode               ExtractData `json:"crc_code"`
}
