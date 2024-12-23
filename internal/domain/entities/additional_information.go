package entities

type AdditionalInformation struct {
	Tag     string                      `json:"tag"`
	Content string                      `json:"content"`
	Data    string                      `json:"data"`
	Detail  AdditionalInformationDetail `json:"detail"`
}

type AdditionalInformationDetail struct {
	BillNumber                    Data `json:"bill_number"`
	MobileNumber                  Data `json:"mobile_number"`
	StoreLabel                    Data `json:"store_label"`
	LoyaltyNumber                 Data `json:"loyalty_number"`
	ReferenceLabel                Data `json:"reference_label"`
	CustomerLabel                 Data `json:"customer_label"`
	TerminalLabel                 Data `json:"terminal_label"`
	PurposeOfTransaction          Data `json:"purpose_of_transaction"`
	AdditionalConsumerDataRequest Data `json:"additional_consumer_data_request"`
	MerchantTaxID                 Data `json:"merchant_tax_id"`
	MerchantChannel               Data `json:"merchant_channel"`
	RFU                           Data `json:"rfu"`
	PaymentSystemSpecific         Data `json:"payment_system_specific"`
}
