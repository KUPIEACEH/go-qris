package services

import (
	"fmt"
	"strings"

	"github.com/fyvri/go-qris/internal/config"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/models"
	"github.com/fyvri/go-qris/pkg/utils"
)

type QRIS struct {
	crc16CCITTUsecase usecases.CRC16CCITTInterface
	qrisUsecase       usecases.QRISInterface
	inputUtil         utils.InputInterface
}

type QRISInterface interface {
	Parse(qrisString string) (*models.QRIS, error, *[]string)
	IsValid(qris *models.QRIS) bool
	Modify(qris *models.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int, terminalLabelValue string) (*models.QRIS, error, *[]string)
	ToString(qris *models.QRIS) string
	Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int, terminalLabelValue string) (string, error, *[]string)
}

func NewQRIS() QRISInterface {
	qrisTags := &usecases.QRISTags{
		Version:               config.VersionTag,
		Category:              config.CategoryTag,
		Acquirer:              config.AcquirerTag,
		AcquirerBankTransfer:  config.AcquirerBankTransferTag,
		Switching:             config.SwitchingTag,
		MerchantCategoryCode:  config.MerchantCategoryCodeTag,
		CurrencyCode:          config.CurrencyCodeTag,
		PaymentAmount:         config.PaymentAmountTag,
		PaymentFeeCategory:    config.PaymentFeeCategoryTag,
		PaymentFeeFixed:       config.PaymentFeeFixedTag,
		PaymentFeePercent:     config.PaymentFeePercentTag,
		CountryCode:           config.CountryCodeTag,
		MerchantName:          config.MerchantNameTag,
		MerchantCity:          config.MerchantCityTag,
		MerchantPostalCode:    config.MerchantPostalCodeTag,
		AdditionalInformation: config.AdditionalInformationTag,
		CRCCode:               config.CRCCodeTag,
	}
	qrisCategoryContents := &usecases.QRISCategoryContents{
		Static:  config.CategoryStaticContent,
		Dynamic: config.CategoryDynamicContent,
	}
	qrisPaymentFeeCategoryContents := &usecases.QRISPaymentFeeCategoryContents{
		Fixed:   config.PaymentFeeCategoryFixedContent,
		Percent: config.PaymentFeeCategoryPercentContent,
	}
	acquirerDetailTags := &usecases.AcquirerDetailTags{
		Site:       config.AcquirerDetailSiteTag,
		MPAN:       config.AcquirerDetailMPANTag,
		TerminalID: config.AcquirerDetailTerminalIDTag,
		Category:   config.AcquirerDetailCategoryTag,
	}
	switchingDetailTags := &usecases.SwitchingDetailTags{
		Site:     config.SwitchingDetailSiteTag,
		NMID:     config.SwitchingDetailNMIDTag,
		Category: config.SwitchingDetailCategoryTag,
	}
	qrisAdditionalInformationDetailTags := &usecases.AdditionalInformationDetailTags{
		BillNumber:                    config.AdditionalInformationDetailBillNumberTag,
		MobileNumber:                  config.AdditionalInformationDetailMobileNumberTag,
		StoreLabel:                    config.AdditionalInformationDetailStoreLabelTag,
		LoyaltyNumber:                 config.AdditionalInformationDetailLoyaltyNumberTag,
		ReferenceLabel:                config.AdditionalInformationDetailReferenceLabelTag,
		CustomerLabel:                 config.AdditionalInformationDetailCustomerLabelTag,
		TerminalLabel:                 config.AdditionalInformationDetailTerminalLabelTag,
		PurposeOfTransaction:          config.AdditionalInformationDetailPurposeOfTransactionTag,
		AdditionalConsumerDataRequest: config.AdditionalInformationDetailAdditionalConsumerDataRequestTag,
		MerchantTaxID:                 config.AdditionalInformationDetailMerchantTaxIDTag,
		MerchantChannel:               config.AdditionalInformationDetailMerchantChannelTag,
		RFUStart:                      config.AdditionalInformationDetailRFUTagStart,
		RFUEnd:                        config.AdditionalInformationDetailRFUTagEnd,
		PaymentSystemSpecificStart:    config.AdditionalInformationDetailPaymentSystemSpecificTagStart,
		PaymentSystemSpecificEnd:      config.AdditionalInformationDetailPaymentSystemSpecificTagEnd,
	}

	dataUsecase := usecases.NewData()
	acquirerUsecase := usecases.NewAcquirer(dataUsecase, acquirerDetailTags)
	switchingUsecase := usecases.NewSwitching(dataUsecase, switchingDetailTags)
	additionalInformationUsecase := usecases.NewAdditionalInformation(dataUsecase, qrisAdditionalInformationDetailTags)
	fieldUsecase := usecases.NewField(acquirerUsecase, switchingUsecase, additionalInformationUsecase, qrisTags, qrisCategoryContents)
	paymentFeeUsecase := usecases.NewPaymentFee(qrisTags, qrisPaymentFeeCategoryContents)
	crc16CCITTUsecase := usecases.NewCRC16CCITT()

	qrisUsecases := &usecases.QRISUsecases{
		Data:                  dataUsecase,
		Field:                 fieldUsecase,
		PaymentFee:            paymentFeeUsecase,
		AdditionalInformation: additionalInformationUsecase,
		CRC16CCITT:            crc16CCITTUsecase,
	}
	qrisUsecase := usecases.NewQRIS(qrisUsecases, qrisTags, qrisCategoryContents, qrisPaymentFeeCategoryContents)
	inputUtil := utils.NewInput()

	return &QRIS{
		crc16CCITTUsecase: crc16CCITTUsecase,
		qrisUsecase:       qrisUsecase,
		inputUtil:         inputUtil,
	}
}

func (s *QRIS) Parse(qrisString string) (*models.QRIS, error, *[]string) {
	qrisString = s.inputUtil.Sanitize(qrisString)
	qris, err, errs := s.qrisUsecase.Parse(qrisString)
	if err != nil {
		return nil, err, errs
	}

	return mapQRISEntityToModel(qris), nil, nil
}

func (s *QRIS) IsValid(qris *models.QRIS) bool {
	qrisEntity := mapQRISModelToEntity(qris)

	return s.qrisUsecase.IsValid(qrisEntity)
}

func (s *QRIS) Modify(qris *models.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int, terminalLabelValue string) (*models.QRIS, error, *[]string) {
	errs := &[]string{}
	merchantCityValue = s.inputUtil.Sanitize(merchantCityValue)
	if len(merchantCityValue) > 15 {
		*errs = append(*errs, "merchant city exceeds 15 characters")
	}
	merchantPostalCodeValue = s.inputUtil.Sanitize(merchantPostalCodeValue)
	if len(merchantPostalCodeValue) > 10 {
		*errs = append(*errs, "merchant postal code exceeds 10 characters")
	}
	terminalLabelValue = s.inputUtil.Sanitize(terminalLabelValue)
	if len(terminalLabelValue) > 99 {
		*errs = append(*errs, "terminal label exceeds 99 characters")
	}
	if len(*errs) > 0 {
		return nil, fmt.Errorf("input length exceeds the maximum permitted characters"), errs
	}

	qrisEntity := mapQRISModelToEntity(qris)
	paymentFeeCategoryValue = strings.ToUpper(s.inputUtil.Sanitize(paymentFeeCategoryValue))
	qrisModel := s.qrisUsecase.Modify(qrisEntity, merchantCityValue, merchantPostalCodeValue, uint32(paymentAmountValue), paymentFeeCategoryValue, uint32(paymentFeeValue), terminalLabelValue)

	return mapQRISEntityToModel(qrisModel), nil, nil
}

func (s *QRIS) ToString(qris *models.QRIS) string {
	qrisString := qris.Version.Data +
		qris.Category.Data +
		qris.Acquirer.Data +
		qris.Switching.Data +
		qris.MerchantCategoryCode.Data +
		qris.CurrencyCode.Data +
		qris.PaymentAmount.Data +
		qris.PaymentFeeCategory.Data +
		qris.PaymentFee.Data +
		qris.CountryCode.Data +
		qris.MerchantName.Data +
		qris.MerchantCity.Data +
		qris.MerchantPostalCode.Data +
		qris.AdditionalInformation.Data +
		qris.CRCCode.Tag + "04"

	return qrisString + s.crc16CCITTUsecase.GenerateCode(qrisString)
}

func (s *QRIS) Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int, terminalLabelValue string) (string, error, *[]string) {
	errs := &[]string{}
	merchantCityValue = s.inputUtil.Sanitize(merchantCityValue)
	if len(merchantCityValue) > 15 {
		*errs = append(*errs, "merchant city exceeds 15 characters")
	}
	merchantPostalCodeValue = s.inputUtil.Sanitize(merchantPostalCodeValue)
	if len(merchantPostalCodeValue) > 10 {
		*errs = append(*errs, "merchant postal code exceeds 10 characters")
	}
	terminalLabelValue = s.inputUtil.Sanitize(terminalLabelValue)
	if len(terminalLabelValue) > 99 {
		*errs = append(*errs, "terminal label exceeds 99 characters")
	}
	if len(*errs) > 0 {
		return "", fmt.Errorf("input length exceeds the maximum permitted characters"), errs
	}

	qrisString = s.inputUtil.Sanitize(qrisString)
	qrisEntity, err, errs := s.qrisUsecase.Parse(qrisString)
	if err != nil {
		return "", err, errs
	}

	paymentFeeCategoryValue = strings.ToUpper(s.inputUtil.Sanitize(paymentFeeCategoryValue))
	qrisEntity = s.qrisUsecase.Modify(qrisEntity, merchantCityValue, merchantPostalCodeValue, uint32(paymentAmountValue), paymentFeeCategoryValue, uint32(paymentFeeValue), terminalLabelValue)

	return s.qrisUsecase.ToString(qrisEntity), nil, nil
}
