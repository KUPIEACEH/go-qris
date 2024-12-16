package usecases

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type QRIS struct {
	dataUsecase                    DataInterface
	fieldUsecase                   FieldInterface
	crc16CCITTUsecase              CRC16CCITTInterface
	qrisTags                       QRISTags
	qrisCategoryContents           QRISCategoryContents
	qrisPaymentFeeCategoryContents QRISPaymentFeeCategoryContents
}

type QRISInterface interface {
	Parse(qrString string) (*entities.QRIS, error, *[]string)
	ToDynamic(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic
	DynamicToString(qrisDynamic *entities.QRISDynamic) string
	Validate(qris *entities.QRIS) bool
}

func NewQRIS(dataUsecase DataInterface, fieldUsecase FieldInterface, crc16CCITTUsecase CRC16CCITTInterface, qrisTags QRISTags, qrisCategoryContents QRISCategoryContents, qrisPaymentFeeCategoryContents QRISPaymentFeeCategoryContents) QRISInterface {
	return &QRIS{
		dataUsecase:                    dataUsecase,
		fieldUsecase:                   fieldUsecase,
		crc16CCITTUsecase:              crc16CCITTUsecase,
		qrisTags:                       qrisTags,
		qrisCategoryContents:           qrisCategoryContents,
		qrisPaymentFeeCategoryContents: qrisPaymentFeeCategoryContents,
	}
}

func (uc *QRIS) Parse(qrString string) (*entities.QRIS, error, *[]string) {
	var qris entities.QRIS
	for len(qrString) > 0 {
		data, err := uc.dataUsecase.Parse(qrString)
		if err != nil {
			return nil, err, nil
		}

		if err := uc.fieldUsecase.Assign(&qris, data); err != nil {
			return nil, err, nil
		}

		qrString = qrString[4+len(data.Content):]
	}

	var errs []string
	if uc.fieldUsecase.Validate(&qris, &errs); errs != nil {
		return nil, fmt.Errorf("invalid QRIS format"), &errs
	}

	return &qris, nil, nil
}

func (uc *QRIS) ToDynamic(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic {
	var (
		content            string
		paymentFeeCategory entities.Data
		paymentFee         entities.Data
	)

	qris.Category = entities.Data{
		Tag:     qris.Category.Tag,
		Content: uc.qrisCategoryContents.Dynamic,
		Data:    qris.Category.Tag + fmt.Sprintf("%02d", len(uc.qrisCategoryContents.Dynamic)) + uc.qrisCategoryContents.Dynamic,
	}

	content = fmt.Sprintf("%d", paymentAmountValue)
	paymentAmount := entities.Data{
		Tag:     uc.qrisTags.PaymentAmountTag,
		Content: content,
		Data:    uc.qrisTags.PaymentAmountTag + fmt.Sprintf("%02d", len(content)) + content,
	}

	if qris.Acquirer.Tag == uc.qrisTags.AcquirerTag {
		if merchantCity != "" {
			qris.MerchantCity = *uc.dataUsecase.ModifyContent(&qris.MerchantCity, merchantCity)
		}
		if merchantPostalCode != "" {
			qris.MerchantPostalCode = *uc.dataUsecase.ModifyContent(&qris.MerchantPostalCode, merchantPostalCode)
		}

		if paymentFeeValue > 0 && paymentFeeCategoryValue != "" {
			paymentFeeCategoryTag := ""
			paymentFeeCategoryContent := ""
			paymentFeeCategoryContentLength := ""
			paymentFeeTag := ""
			if paymentFeeCategoryValue == "FIXED" {
				paymentFeeCategoryTag = uc.qrisTags.PaymentFeeCategoryTag
				paymentFeeCategoryContent = uc.qrisPaymentFeeCategoryContents.Fixed
				paymentFeeCategoryContentLength = fmt.Sprintf("%02d", len(paymentFeeCategoryContent))
				paymentFeeTag = uc.qrisTags.PaymentFeeFixedTag
			} else if paymentFeeCategoryValue == "PERCENT" {
				paymentFeeCategoryTag = uc.qrisTags.PaymentFeeCategoryTag
				paymentFeeCategoryContent = uc.qrisPaymentFeeCategoryContents.Percent
				paymentFeeCategoryContentLength = fmt.Sprintf("%02d", len(paymentFeeCategoryContent))
				paymentFeeTag = uc.qrisTags.PaymentFeePercentTag
			}

			paymentFeeCategory = entities.Data{
				Tag:     paymentFeeCategoryTag,
				Content: paymentFeeCategoryContent,
				Data:    paymentFeeCategoryTag + paymentFeeCategoryContentLength + paymentFeeCategoryContent,
			}
			if paymentFeeCategory.Tag != "" {
				content = fmt.Sprintf("%d", paymentFeeValue)
				paymentFee = entities.Data{
					Tag:     paymentFeeTag,
					Content: content,
					Data:    paymentFeeTag + fmt.Sprintf("%02d", len(content)) + content,
				}
			}
		}
	}

	qrStringConverted := qris.Version.Data +
		qris.Category.Data +
		qris.Acquirer.Data +
		qris.Switching.Data +
		qris.MerchantCategoryCode.Data +
		qris.CurrencyCode.Data +
		paymentAmount.Data +
		paymentFeeCategory.Data +
		paymentFee.Data +
		qris.CountryCode.Data +
		qris.MerchantName.Data +
		qris.MerchantCity.Data +
		qris.MerchantPostalCode.Data +
		qris.AdditionalInformation.Data +
		qris.CRCCode.Tag + "04"
	crc16CCITTCode := uc.crc16CCITTUsecase.GenerateCode(qrStringConverted)

	return &entities.QRISDynamic{
		Version:               qris.Version,
		Category:              qris.Category,
		Acquirer:              qris.Acquirer,
		Switching:             qris.Switching,
		MerchantCategoryCode:  qris.MerchantCategoryCode,
		CurrencyCode:          qris.CurrencyCode,
		PaymentAmount:         paymentAmount,
		PaymentFeeCategory:    paymentFeeCategory,
		PaymentFee:            paymentFee,
		CountryCode:           qris.CountryCode,
		MerchantName:          qris.MerchantName,
		MerchantCity:          qris.MerchantCity,
		MerchantPostalCode:    qris.MerchantPostalCode,
		AdditionalInformation: qris.AdditionalInformation,
		CRCCode: entities.Data{
			Tag:     uc.qrisTags.CRCCodeTag,
			Content: crc16CCITTCode,
			Data:    uc.qrisTags.CRCCodeTag + fmt.Sprintf("%02d", len(crc16CCITTCode)) + crc16CCITTCode,
		},
	}
}

func (uc *QRIS) DynamicToString(qrisDynamic *entities.QRISDynamic) string {
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

func (uc *QRIS) Validate(qris *entities.QRIS) bool {
	qrStringConverted := qris.Version.Data +
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

	return qris.CRCCode.Content == uc.crc16CCITTUsecase.GenerateCode(qrStringConverted)
}
