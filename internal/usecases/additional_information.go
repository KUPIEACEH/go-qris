package usecases

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
)

type AdditionalInformation struct {
	dataUsecase                     DataInterface
	additionalInformationDetailTags *AdditionalInformationDetailTags
}

type AdditionalInformationDetailTags struct {
	BillNumber                    string
	MobileNumber                  string
	StoreLabel                    string
	LoyaltyNumber                 string
	ReferenceLabel                string
	CustomerLabel                 string
	TerminalLabel                 string
	PurposeOfTransaction          string
	AdditionalConsumerDataRequest string
	MerchantTaxID                 string
	MerchantChannel               string
	RFUStart                      string
	RFUEnd                        string
	PaymentSystemSpecificStart    string
	PaymentSystemSpecificEnd      string
}

type AdditionalInformationInterface interface {
	Parse(content string) (*entities.AdditionalInformationDetail, error)
	ToString(additionalInformationDetail *entities.AdditionalInformationDetail) string
}

func NewAdditionalInformation(dataUsecase DataInterface, additionalInformationDetailTags *AdditionalInformationDetailTags) AdditionalInformationInterface {
	return &AdditionalInformation{
		dataUsecase:                     dataUsecase,
		additionalInformationDetailTags: additionalInformationDetailTags,
	}
}

func (uc *AdditionalInformation) Parse(content string) (*entities.AdditionalInformationDetail, error) {
	var detail entities.AdditionalInformationDetail
	for len(content) > 0 {
		data, err := uc.dataUsecase.Parse(content)
		if err != nil {
			return nil, err
		}

		switch {
		case data.Tag == uc.additionalInformationDetailTags.BillNumber:
			detail.BillNumber = *data
		case data.Tag == uc.additionalInformationDetailTags.MobileNumber:
			detail.MobileNumber = *data
		case data.Tag == uc.additionalInformationDetailTags.StoreLabel:
			detail.StoreLabel = *data
		case data.Tag == uc.additionalInformationDetailTags.LoyaltyNumber:
			detail.LoyaltyNumber = *data
		case data.Tag == uc.additionalInformationDetailTags.ReferenceLabel:
			detail.ReferenceLabel = *data
		case data.Tag == uc.additionalInformationDetailTags.CustomerLabel:
			detail.CustomerLabel = *data
		case data.Tag == uc.additionalInformationDetailTags.TerminalLabel:
			detail.TerminalLabel = *data
		case data.Tag == uc.additionalInformationDetailTags.PurposeOfTransaction:
			detail.PurposeOfTransaction = *data
		case data.Tag == uc.additionalInformationDetailTags.AdditionalConsumerDataRequest:
			detail.AdditionalConsumerDataRequest = *data
		case data.Tag == uc.additionalInformationDetailTags.MerchantTaxID:
			detail.MerchantTaxID = *data
		case data.Tag == uc.additionalInformationDetailTags.MerchantChannel:
			detail.MerchantChannel = *data
		case inRange(data.Tag, uc.additionalInformationDetailTags.RFUStart, uc.additionalInformationDetailTags.RFUEnd):
			detail.RFU = *data
		case inRange(data.Tag, uc.additionalInformationDetailTags.PaymentSystemSpecificStart, uc.additionalInformationDetailTags.PaymentSystemSpecificEnd):
			detail.PaymentSystemSpecific = *data
		}

		content = content[4+len(data.Content):]
	}

	return &detail, nil
}

func (uc *AdditionalInformation) ToString(additionalInformationDetail *entities.AdditionalInformationDetail) string {
	return additionalInformationDetail.BillNumber.Data +
		additionalInformationDetail.MobileNumber.Data +
		additionalInformationDetail.StoreLabel.Data +
		additionalInformationDetail.LoyaltyNumber.Data +
		additionalInformationDetail.ReferenceLabel.Data +
		additionalInformationDetail.CustomerLabel.Data +
		additionalInformationDetail.TerminalLabel.Data +
		additionalInformationDetail.PurposeOfTransaction.Data +
		additionalInformationDetail.AdditionalConsumerDataRequest.Data +
		additionalInformationDetail.MerchantTaxID.Data +
		additionalInformationDetail.MerchantChannel.Data +
		additionalInformationDetail.RFU.Data +
		additionalInformationDetail.PaymentSystemSpecific.Data
}
