package usecases

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type QRISTags struct {
	VersionTag               string
	CategoryTag              string
	AcquirerTag              string
	SwitchingTag             string
	MerchantCategoryCodeTag  string
	CurrencyCodeTag          string
	PaymentAmountTag         string
	PaymentFeeCategoryTag    string
	CountryCodeTag           string
	MerchantNameTag          string
	MerchantCityTag          string
	MerchantPostalCodeTag    string
	AdditionalInformationTag string
	CRCCodeTag               string
}

type QRISCategoryContents struct {
	StaticContent  string
	DynamicContent string
}

type QRISDynamicPaymentFee struct {
	CategoryFixedContent   string
	CategoryPercentContent string
	FixedTag               string
	PercentTag             string
}

type QRIS struct {
	acquirerUsecase       AcquirerInterface
	switchingUsecase      SwitchingInterface
	dataUsecase           DataInterface
	crc16CCITTUsecase     CRC16CCITTInterface
	qrisTags              QRISTags
	qrisCategoryContents  QRISCategoryContents
	qrisDynamicPaymentFee QRISDynamicPaymentFee
}

type QRISInterface interface {
	ExtractStatic(qrString string) (*entities.QRISStatic, error)
	StaticToDynamic(qrisStatic *entities.QRISStatic, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic
	DynamicToDynamicString(qrisDynamic *entities.QRISDynamic) string
}

func NewQRIS(acquirerUsecase AcquirerInterface, switchingUsecase SwitchingInterface, dataUsecase DataInterface, crc16CCITTUsecase CRC16CCITTInterface, qrisTags QRISTags, qrisCategoryContents QRISCategoryContents, qrisDynamicPaymentFee QRISDynamicPaymentFee) QRISInterface {
	return &QRIS{
		acquirerUsecase:       acquirerUsecase,
		switchingUsecase:      switchingUsecase,
		dataUsecase:           dataUsecase,
		crc16CCITTUsecase:     crc16CCITTUsecase,
		qrisTags:              qrisTags,
		qrisCategoryContents:  qrisCategoryContents,
		qrisDynamicPaymentFee: qrisDynamicPaymentFee,
	}
}

func (uc *QRIS) ExtractStatic(qrString string) (*entities.QRISStatic, error) {
	var qrisStatic entities.QRISStatic
	for len(qrString) > 0 {
		extractData, err := uc.dataUsecase.Extract(qrString)
		if err != nil {
			return nil, err
		}

		switch extractData.Tag {
		case uc.qrisTags.VersionTag:
			qrisStatic.Version = *extractData
		case uc.qrisTags.CategoryTag:
			qrisStatic.Category = *extractData
			if extractData.Content != uc.qrisCategoryContents.StaticContent {
				return nil, fmt.Errorf("not static QRIS content detected")
			}
		case uc.qrisTags.AcquirerTag:
			acquirerDetail, err := uc.acquirerUsecase.Extract(extractData.Content)
			if err != nil {
				return nil, fmt.Errorf("invalid extract acquirer for content %s", extractData.Content)
			}
			qrisStatic.Acquirer = entities.Acquirer{
				Tag:     extractData.Tag,
				Content: extractData.Content,
				Data:    extractData.Data,
				Detail:  *acquirerDetail,
			}
		case uc.qrisTags.SwitchingTag:
			switchingDetail, err := uc.switchingUsecase.Extract(extractData.Content)
			if err != nil {
				return nil, fmt.Errorf("invalid extract switching for content %s", extractData.Content)
			}
			qrisStatic.Switching = entities.Switching{
				Tag:     extractData.Tag,
				Content: extractData.Content,
				Data:    extractData.Data,
				Detail:  *switchingDetail,
			}
		case uc.qrisTags.MerchantCategoryCodeTag:
			qrisStatic.MerchantCategoryCode = *extractData
		case uc.qrisTags.CurrencyCodeTag:
			qrisStatic.CurrencyCode = *extractData
		case uc.qrisTags.CountryCodeTag:
			qrisStatic.CountryCode = *extractData
		case uc.qrisTags.MerchantNameTag:
			qrisStatic.MerchantName = *extractData
		case uc.qrisTags.MerchantCityTag:
			qrisStatic.MerchantCity = *extractData
		case uc.qrisTags.MerchantPostalCodeTag:
			qrisStatic.MerchantPostalCode = *extractData
		case uc.qrisTags.AdditionalInformationTag:
			qrisStatic.AdditionalInformation = *extractData
		case uc.qrisTags.CRCCodeTag:
			qrisStatic.CRCCode = *extractData
		}

		qrString = qrString[4+len(extractData.Content):]
	}

	return &qrisStatic, nil
}

func (uc *QRIS) StaticToDynamic(qrisStatic *entities.QRISStatic, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic {
	var (
		content            string
		paymentFeeCategory entities.ExtractData
		paymentFee         entities.ExtractData
	)

	if merchantCity != "" {
		qrisStatic.MerchantCity = uc.dataUsecase.ModifyContent(qrisStatic.MerchantCity, merchantCity)
	}

	if merchantPostalCode != "" {
		qrisStatic.MerchantPostalCode = uc.dataUsecase.ModifyContent(qrisStatic.MerchantPostalCode, merchantPostalCode)
	}

	qrisStatic.Category = entities.ExtractData{
		Tag:     qrisStatic.Category.Tag,
		Content: uc.qrisCategoryContents.DynamicContent,
		Data:    qrisStatic.Category.Tag + fmt.Sprintf("%02d", len(uc.qrisCategoryContents.DynamicContent)) + uc.qrisCategoryContents.DynamicContent,
	}

	content = fmt.Sprintf("%d", paymentAmountValue)
	paymentAmount := entities.ExtractData{
		Tag:     uc.qrisTags.PaymentAmountTag,
		Content: content,
		Data:    uc.qrisTags.PaymentAmountTag + fmt.Sprintf("%02d", len(content)) + content,
	}

	if paymentFeeValue > 0 && paymentFeeCategoryValue != "" {
		paymentFeeCategoryTag := ""
		paymentFeeCategoryContent := ""
		paymentFeeCategoryContentLength := ""
		paymentFeeTag := ""
		if paymentFeeCategoryValue == "FIXED" {
			paymentFeeCategoryTag = uc.qrisTags.PaymentFeeCategoryTag
			paymentFeeCategoryContent = uc.qrisDynamicPaymentFee.CategoryFixedContent
			paymentFeeCategoryContentLength = fmt.Sprintf("%02d", len(paymentFeeCategoryContent))
			paymentFeeTag = uc.qrisDynamicPaymentFee.FixedTag
		} else if paymentFeeCategoryValue == "PERCENT" {
			paymentFeeCategoryTag = uc.qrisTags.PaymentFeeCategoryTag
			paymentFeeCategoryContent = uc.qrisDynamicPaymentFee.CategoryPercentContent
			paymentFeeCategoryContentLength = fmt.Sprintf("%02d", len(paymentFeeCategoryContent))
			paymentFeeTag = uc.qrisDynamicPaymentFee.PercentTag
		}

		paymentFeeCategory = entities.ExtractData{
			Tag:     paymentFeeCategoryTag,
			Content: paymentFeeCategoryContent,
			Data:    paymentFeeCategoryTag + paymentFeeCategoryContentLength + paymentFeeCategoryContent,
		}
		if paymentFeeCategory.Tag != "" {
			content = fmt.Sprintf("%d", paymentFeeValue)
			paymentFee = entities.ExtractData{
				Tag:     paymentFeeTag,
				Content: content,
				Data:    paymentFeeTag + fmt.Sprintf("%02d", len(content)) + content,
			}
		}
	}

	qrStringConverted := qrisStatic.Version.Data +
		qrisStatic.Category.Data +
		qrisStatic.Acquirer.Data +
		qrisStatic.Switching.Data +
		qrisStatic.MerchantCategoryCode.Data +
		qrisStatic.CurrencyCode.Data +
		paymentAmount.Data +
		paymentFeeCategory.Data +
		paymentFee.Data +
		qrisStatic.CountryCode.Data +
		qrisStatic.MerchantName.Data +
		qrisStatic.MerchantCity.Data +
		qrisStatic.MerchantPostalCode.Data +
		qrisStatic.AdditionalInformation.Data +
		qrisStatic.CRCCode.Tag + "04"
	crc16CCITTCode := uc.crc16CCITTUsecase.GenerateCode(qrStringConverted)

	return &entities.QRISDynamic{
		Version:               qrisStatic.Version,
		Category:              qrisStatic.Category,
		Acquirer:              qrisStatic.Acquirer,
		Switching:             qrisStatic.Switching,
		MerchantCategoryCode:  qrisStatic.MerchantCategoryCode,
		CurrencyCode:          qrisStatic.CurrencyCode,
		PaymentAmount:         paymentAmount,
		PaymentFeeCategory:    paymentFeeCategory,
		PaymentFee:            paymentFee,
		CountryCode:           qrisStatic.CountryCode,
		MerchantName:          qrisStatic.MerchantName,
		MerchantCity:          qrisStatic.MerchantCity,
		MerchantPostalCode:    qrisStatic.MerchantPostalCode,
		AdditionalInformation: qrisStatic.AdditionalInformation,
		CRCCode: entities.ExtractData{
			Tag:     uc.qrisTags.CRCCodeTag,
			Content: crc16CCITTCode,
			Data:    uc.qrisTags.CRCCodeTag + fmt.Sprintf("%02d", len(crc16CCITTCode)) + crc16CCITTCode,
		},
	}
}

func (uc *QRIS) DynamicToDynamicString(qrisDynamic *entities.QRISDynamic) string {
	return qrisDynamic.Version.Data +
		qrisDynamic.Category.Data +
		qrisDynamic.Acquirer.Data +
		qrisDynamic.Switching.Data +
		qrisDynamic.MerchantCategoryCode.Data +
		qrisDynamic.CurrencyCode.Data +
		qrisDynamic.PaymentAmount.Data +
		qrisDynamic.PaymentFeeCategory.Data +
		qrisDynamic.PaymentFee.Data +
		qrisDynamic.CountryCode.Data +
		qrisDynamic.MerchantName.Data +
		qrisDynamic.MerchantCity.Data +
		qrisDynamic.MerchantPostalCode.Data +
		qrisDynamic.AdditionalInformation.Data +
		qrisDynamic.CRCCode.Data
}
