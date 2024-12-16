package entities

type QRIS struct {
	Version               Data      `json:"version"`
	Category              Data      `json:"category"`
	Acquirer              Acquirer  `json:"acquirer"`
	Switching             Switching `json:"switching"`
	MerchantCategoryCode  Data      `json:"merchant_category_code"`
	CurrencyCode          Data      `json:"currency_code"`
	PaymentAmount         Data      `json:"payment_amount"`
	PaymentFeeCategory    Data      `json:"payment_fee_category"`
	PaymentFee            Data      `json:"payment_fee"`
	CountryCode           Data      `json:"country_code"`
	MerchantName          Data      `json:"merchant_name"`
	MerchantCity          Data      `json:"merchant_city"`
	MerchantPostalCode    Data      `json:"merchant_postal_code"`
	AdditionalInformation Data      `json:"additional_information"`
	CRCCode               Data      `json:"crc_code"`
}
