package services

import (
	"github.com/fyvri/go-qris/internal/config"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/models"
)

type QRIS struct {
	crc16CCITTUsecase usecases.CRC16CCITTInterface
	qrisUsecase       usecases.QRISInterface
}

type Schema struct {
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

	CategoryStaticContent  string
	CategoryDynamicContent string

	AcquirerDetailSiteTag       string
	AcquirerDetailMPANTag       string
	AcquirerDetailTerminalIDTag string
	AcquirerDetailCategoryTag   string

	SwitchingDetailSiteTag     string
	SwitchingDetailNMIDTag     string
	SwitchingDetailCategoryTag string

	PaymentFeeCategoryFixedContent   string
	PaymentFeeCategoryPercentContent string
}

type QRISInterface interface {
	Parse(qrisString string) (*models.QRIS, error, *[]string)
	Validate(qris *models.QRIS) bool
	ToDynamic(qris *models.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *models.QRISDynamic
	ToString(qris *models.QRIS) string
	Convert(qrisString string, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) (string, error, *[]string)
}

func NewQRIS(schema *Schema) QRISInterface {
	defineValue := func(value string, defaultValue string) string {
		if value == "" {
			return defaultValue
		}
		return value
	}

	if schema == nil {
		schema = &Schema{}
	}
	qrisTags := usecases.QRISTags{
		VersionTag:               defineValue(schema.VersionTag, config.VersionTag),
		CategoryTag:              defineValue(schema.CategoryTag, config.CategoryTag),
		AcquirerTag:              defineValue(schema.AcquirerTag, config.AcquirerTag),
		AcquirerBankTransferTag:  defineValue(schema.AcquirerBankTransferTag, config.AcquirerBankTransferTag),
		SwitchingTag:             defineValue(schema.SwitchingTag, config.SwitchingTag),
		MerchantCategoryCodeTag:  defineValue(schema.MerchantCategoryCodeTag, config.MerchantCategoryCodeTag),
		CurrencyCodeTag:          defineValue(schema.CurrencyCodeTag, config.CurrencyCodeTag),
		PaymentAmountTag:         defineValue(schema.PaymentAmountTag, config.PaymentAmountTag),
		PaymentFeeCategoryTag:    defineValue(schema.PaymentFeeCategoryTag, config.PaymentFeeCategoryTag),
		PaymentFeeFixedTag:       defineValue(schema.PaymentFeeFixedTag, config.PaymentFeeFixedTag),
		PaymentFeePercentTag:     defineValue(schema.PaymentFeePercentTag, config.PaymentFeePercentTag),
		CountryCodeTag:           defineValue(schema.CountryCodeTag, config.CountryCodeTag),
		MerchantNameTag:          defineValue(schema.MerchantNameTag, config.MerchantNameTag),
		MerchantCityTag:          defineValue(schema.MerchantCityTag, config.MerchantCityTag),
		MerchantPostalCodeTag:    defineValue(schema.MerchantPostalCodeTag, config.MerchantPostalCodeTag),
		AdditionalInformationTag: defineValue(schema.AdditionalInformationTag, config.AdditionalInformationTag),
		CRCCodeTag:               defineValue(schema.CRCCodeTag, config.CRCCodeTag),
	}
	qrisCategoryContents := usecases.QRISCategoryContents{
		Static:  defineValue(schema.CategoryStaticContent, config.CategoryStaticContent),
		Dynamic: defineValue(schema.CategoryDynamicContent, config.CategoryDynamicContent),
	}
	qrisPaymentFeeCategoryContents := usecases.QRISPaymentFeeCategoryContents{
		Fixed:   defineValue(schema.PaymentFeeCategoryFixedContent, config.PaymentFeeCategoryFixedContent),
		Percent: defineValue(schema.PaymentFeeCategoryPercentContent, config.PaymentFeeCategoryPercentContent),
	}

	dataUsecase := usecases.NewData()
	acquirerUsecase := usecases.NewAcquirer(dataUsecase, schema.AcquirerDetailSiteTag, schema.AcquirerDetailMPANTag, schema.AcquirerDetailTerminalIDTag, schema.AcquirerDetailCategoryTag)
	switchingUsecase := usecases.NewSwitching(dataUsecase, schema.SwitchingDetailSiteTag, schema.SwitchingDetailNMIDTag, schema.SwitchingDetailCategoryTag)
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

	return &QRIS{
		crc16CCITTUsecase: crc16CCITTUsecase,
		qrisUsecase:       qrisUsecase,
	}
}

func (s *QRIS) Parse(qrisString string) (*models.QRIS, error, *[]string) {
	qris, err, errs := s.qrisUsecase.Parse(qrisString)
	if err != nil {
		return nil, err, errs
	}

	return mapQRISEntityToModel(qris), nil, nil
}

func (s *QRIS) Validate(qris *models.QRIS) bool {
	internalQRIS := mapQRISModelToEntity(qris)

	return s.qrisUsecase.Validate(internalQRIS)
}

func (s *QRIS) ToDynamic(qris *models.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *models.QRISDynamic {
	internalQRIS := mapQRISModelToEntity(qris)
	qrisDynamic := s.qrisUsecase.ToDynamic(internalQRIS, merchantCity, merchantPostalCode, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)

	return mapQRISDynamicEntityToModel(qrisDynamic)
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

func (s *QRIS) Convert(qrisString string, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) (string, error, *[]string) {
	qris, err, errs := s.qrisUsecase.Parse(qrisString)
	if err != nil {
		return "", err, errs
	}

	return s.qrisUsecase.DynamicToString(s.qrisUsecase.ToDynamic(qris, merchantCity, merchantPostalCode, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)), nil, nil
}
