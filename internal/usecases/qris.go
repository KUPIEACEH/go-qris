package usecases

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type QRIS struct {
	dataUsecase                    DataInterface
	fieldUsecase                   FieldInterface
	crc16CCITTUsecase              CRC16CCITTInterface
	qrisTags                       *QRISTags
	qrisCategoryContents           *QRISCategoryContents
	qrisPaymentFeeCategoryContents *QRISPaymentFeeCategoryContents
}

type QRISInterface interface {
	Parse(qrString string) (*entities.QRIS, error, *[]string)
	IsValid(qris *entities.QRIS) bool
	Modify(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS
	ToString(qris *entities.QRIS) string
}

func NewQRIS(dataUsecase DataInterface, fieldUsecase FieldInterface, crc16CCITTUsecase CRC16CCITTInterface, qrisTags *QRISTags, qrisCategoryContents *QRISCategoryContents, qrisPaymentFeeCategoryContents *QRISPaymentFeeCategoryContents) QRISInterface {
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
	if uc.fieldUsecase.IsValid(&qris, &errs); errs != nil {
		return nil, fmt.Errorf("invalid QRIS format"), &errs
	}

	return &qris, nil, nil
}

func (uc *QRIS) IsValid(qris *entities.QRIS) bool {
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

func (uc *QRIS) Modify(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
	qris.Category = entities.Data{
		Tag:     qris.Category.Tag,
		Content: uc.qrisCategoryContents.Dynamic,
		Data:    qris.Category.Tag + fmt.Sprintf("%02d", len(uc.qrisCategoryContents.Dynamic)) + uc.qrisCategoryContents.Dynamic,
	}

	content := fmt.Sprintf("%d", paymentAmountValue)
	qris.PaymentAmount = *uc.dataUsecase.ModifyContent(&entities.Data{
		Tag:     uc.qrisTags.PaymentAmountTag,
		Content: content,
		Data:    uc.qrisTags.PaymentAmountTag + fmt.Sprintf("%02d", len(content)) + content,
	}, content)

	if qris.Acquirer.Tag == uc.qrisTags.AcquirerTag {
		if merchantCityValue != "" {
			qris.MerchantCity = *uc.dataUsecase.ModifyContent(&qris.MerchantCity, merchantCityValue)
		}
		if merchantPostalCodeValue != "" {
			qris.MerchantPostalCode = *uc.dataUsecase.ModifyContent(&qris.MerchantPostalCode, merchantPostalCodeValue)
		}

		qris.PaymentFee = entities.Data{}
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

			qris.PaymentFeeCategory = entities.Data{
				Tag:     paymentFeeCategoryTag,
				Content: paymentFeeCategoryContent,
				Data:    paymentFeeCategoryTag + paymentFeeCategoryContentLength + paymentFeeCategoryContent,
			}
			if qris.PaymentFeeCategory.Tag != "" {
				content = fmt.Sprintf("%d", paymentFeeValue)
				qris.PaymentFee = entities.Data{
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
		qris.PaymentAmount.Data +
		qris.PaymentFeeCategory.Data +
		qris.PaymentFee.Data +
		qris.CountryCode.Data +
		qris.MerchantName.Data +
		qris.MerchantCity.Data +
		qris.MerchantPostalCode.Data +
		qris.AdditionalInformation.Data +
		qris.CRCCode.Tag + "04"
	content = uc.crc16CCITTUsecase.GenerateCode(qrStringConverted)
	qris.CRCCode = *uc.dataUsecase.ModifyContent(&qris.CRCCode, content)

	return qris
}

func (uc *QRIS) ToString(qris *entities.QRIS) string {
	return qris.Version.Data +
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
		qris.CRCCode.Data
}
