package services

import (
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
	Modify(qris *models.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int) *models.QRIS
	ToString(qris *models.QRIS) string
	Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int) (string, error, *[]string)
}

func NewQRIS() QRISInterface {
	qrisTags := &usecases.QRISTags{
		VersionTag:               config.VersionTag,
		CategoryTag:              config.CategoryTag,
		AcquirerTag:              config.AcquirerTag,
		AcquirerBankTransferTag:  config.AcquirerBankTransferTag,
		SwitchingTag:             config.SwitchingTag,
		MerchantCategoryCodeTag:  config.MerchantCategoryCodeTag,
		CurrencyCodeTag:          config.CurrencyCodeTag,
		PaymentAmountTag:         config.PaymentAmountTag,
		PaymentFeeCategoryTag:    config.PaymentFeeCategoryTag,
		PaymentFeeFixedTag:       config.PaymentFeeFixedTag,
		PaymentFeePercentTag:     config.PaymentFeePercentTag,
		CountryCodeTag:           config.CountryCodeTag,
		MerchantNameTag:          config.MerchantNameTag,
		MerchantCityTag:          config.MerchantCityTag,
		MerchantPostalCodeTag:    config.MerchantPostalCodeTag,
		AdditionalInformationTag: config.AdditionalInformationTag,
		CRCCodeTag:               config.CRCCodeTag,
	}
	qrisCategoryContents := &usecases.QRISCategoryContents{
		Static:  config.CategoryStaticContent,
		Dynamic: config.CategoryDynamicContent,
	}
	qrisPaymentFeeCategoryContents := &usecases.QRISPaymentFeeCategoryContents{
		Fixed:   config.PaymentFeeCategoryFixedContent,
		Percent: config.PaymentFeeCategoryPercentContent,
	}

	dataUsecase := usecases.NewData()
	acquirerUsecase := usecases.NewAcquirer(
		dataUsecase,
		config.AcquirerDetailSiteTag,
		config.AcquirerDetailMPANTag,
		config.AcquirerDetailTerminalIDTag,
		config.AcquirerDetailCategoryTag,
	)
	switchingUsecase := usecases.NewSwitching(
		dataUsecase,
		config.SwitchingDetailSiteTag,
		config.SwitchingDetailNMIDTag,
		config.SwitchingDetailCategoryTag,
	)
	fieldUsecase := usecases.NewField(acquirerUsecase, switchingUsecase, qrisTags, qrisCategoryContents)
	crc16CCITTUsecase := usecases.NewCRC16CCITT()
	qrisUsecase := usecases.NewQRIS(
		dataUsecase,
		fieldUsecase,
		crc16CCITTUsecase,
		qrisTags,
		qrisCategoryContents,
		qrisPaymentFeeCategoryContents,
	)
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

func (s *QRIS) Modify(qris *models.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int) *models.QRIS {
	qrisEntity := mapQRISModelToEntity(qris)

	merchantCityValue = s.inputUtil.Sanitize(merchantCityValue)
	merchantPostalCodeValue = s.inputUtil.Sanitize(merchantPostalCodeValue)
	paymentFeeCategoryValue = strings.ToUpper(s.inputUtil.Sanitize(paymentFeeCategoryValue))
	qrisModel := s.qrisUsecase.Modify(qrisEntity, merchantCityValue, merchantPostalCodeValue, uint32(paymentAmountValue), paymentFeeCategoryValue, uint32(paymentFeeValue))

	return mapQRISEntityToModel(qrisModel)
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

func (s *QRIS) Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int) (string, error, *[]string) {
	qrisString = s.inputUtil.Sanitize(qrisString)
	qrisEntity, err, errs := s.qrisUsecase.Parse(qrisString)
	if err != nil {
		return "", err, errs
	}

	merchantCityValue = s.inputUtil.Sanitize(merchantCityValue)
	merchantPostalCodeValue = s.inputUtil.Sanitize(merchantPostalCodeValue)
	paymentFeeCategoryValue = strings.ToUpper(s.inputUtil.Sanitize(paymentFeeCategoryValue))
	qrisEntity = s.qrisUsecase.Modify(qrisEntity, merchantCityValue, merchantPostalCodeValue, uint32(paymentAmountValue), paymentFeeCategoryValue, uint32(paymentFeeValue))

	return s.qrisUsecase.ToString(qrisEntity), nil, nil
}
