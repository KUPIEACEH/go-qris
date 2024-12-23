package usecases

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type QRIS struct {
	qrisUsecases                   *QRISUsecases
	qrisTags                       *QRISTags
	qrisCategoryContents           *QRISCategoryContents
	qrisPaymentFeeCategoryContents *QRISPaymentFeeCategoryContents
}

type QRISUsecases struct {
	Data                  DataInterface
	Field                 FieldInterface
	PaymentFee            PaymentFeeInterface
	AdditionalInformation AdditionalInformationInterface
	CRC16CCITT            CRC16CCITTInterface
}

type QRISInterface interface {
	Parse(qrString string) (*entities.QRIS, error, *[]string)
	IsValid(qris *entities.QRIS) bool
	Modify(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS
	ToString(qris *entities.QRIS) string
}

func NewQRIS(qrisUsecases *QRISUsecases, qrisTags *QRISTags, qrisCategoryContents *QRISCategoryContents, qrisPaymentFeeCategoryContents *QRISPaymentFeeCategoryContents) QRISInterface {
	return &QRIS{
		qrisUsecases:                   qrisUsecases,
		qrisTags:                       qrisTags,
		qrisCategoryContents:           qrisCategoryContents,
		qrisPaymentFeeCategoryContents: qrisPaymentFeeCategoryContents,
	}
}

func (uc *QRIS) Parse(qrString string) (*entities.QRIS, error, *[]string) {
	var qris entities.QRIS
	for len(qrString) > 0 {
		data, err := uc.qrisUsecases.Data.Parse(qrString)
		if err != nil {
			return nil, err, nil
		}

		if err := uc.qrisUsecases.Field.Assign(&qris, data); err != nil {
			return nil, err, nil
		}

		qrString = qrString[4+len(data.Content):]
	}

	var errs []string
	if uc.qrisUsecases.Field.IsValid(&qris, &errs); errs != nil {
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

	return qris.CRCCode.Content == uc.qrisUsecases.CRC16CCITT.GenerateCode(qrStringConverted)
}

func (uc *QRIS) Modify(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
	qris.Category = entities.Data{
		Tag:     qris.Category.Tag,
		Content: uc.qrisCategoryContents.Dynamic,
		Data:    qris.Category.Tag + fmt.Sprintf("%02d", len(uc.qrisCategoryContents.Dynamic)) + uc.qrisCategoryContents.Dynamic,
	}

	content := fmt.Sprintf("%d", paymentAmountValue)
	qris.PaymentAmount = *uc.qrisUsecases.Data.ModifyContent(&entities.Data{
		Tag:     uc.qrisTags.PaymentAmount,
		Content: content,
		Data:    uc.qrisTags.PaymentAmount + fmt.Sprintf("%02d", len(content)) + content,
	}, content)

	qris.PaymentFeeCategory = entities.Data{}
	qris.PaymentFee = entities.Data{}
	if qris.Acquirer.Tag == uc.qrisTags.Acquirer {
		if merchantCityValue != "" {
			qris.MerchantCity = *uc.qrisUsecases.Data.ModifyContent(&qris.MerchantCity, merchantCityValue)
		}
		if merchantPostalCodeValue != "" {
			qris.MerchantPostalCode = *uc.qrisUsecases.Data.ModifyContent(&qris.MerchantPostalCode, merchantPostalCodeValue)
		}

		if paymentFeeCategoryValue != "" && paymentFeeValue > 0 {
			qris = uc.qrisUsecases.PaymentFee.Modify(qris, paymentFeeCategoryValue, paymentFeeValue)
		}
	}

	if qris.AdditionalInformation.Detail.TerminalLabel.Tag != "" && terminalLabelValue != "" {
		qris.AdditionalInformation.Detail.TerminalLabel = *uc.qrisUsecases.Data.ModifyContent(&qris.AdditionalInformation.Detail.TerminalLabel, terminalLabelValue)
		qrisAdditionalInformationContent := uc.qrisUsecases.AdditionalInformation.ToString(&qris.AdditionalInformation.Detail)
		qris.AdditionalInformation = entities.AdditionalInformation{
			Tag:     uc.qrisTags.AdditionalInformation,
			Content: qrisAdditionalInformationContent,
			Data:    uc.qrisTags.AdditionalInformation + fmt.Sprintf("%02d", len(qrisAdditionalInformationContent)) + qrisAdditionalInformationContent,
			Detail:  qris.AdditionalInformation.Detail,
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
	content = uc.qrisUsecases.CRC16CCITT.GenerateCode(qrStringConverted)
	qris.CRCCode = *uc.qrisUsecases.Data.ModifyContent(&qris.CRCCode, content)

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
