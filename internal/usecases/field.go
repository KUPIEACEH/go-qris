package usecases

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Field struct {
	acquirerUsecase              AcquirerInterface
	switchingUsecase             SwitchingInterface
	additionalInformationUsecase AdditionalInformationInterface
	qrisTags                     *QRISTags
	qrisCategoryContents         *QRISCategoryContents
}

type FieldInterface interface {
	Assign(qris *entities.QRIS, data *entities.Data) error
	IsValid(qris *entities.QRIS, errs *[]string)
}

func NewField(acquirerUsecase AcquirerInterface, switchingUsecase SwitchingInterface, additionalInformationUsecase AdditionalInformationInterface, qrisTags *QRISTags, qrisCategoryContents *QRISCategoryContents) FieldInterface {
	return &Field{
		acquirerUsecase:              acquirerUsecase,
		switchingUsecase:             switchingUsecase,
		additionalInformationUsecase: additionalInformationUsecase,
		qrisTags:                     qrisTags,
		qrisCategoryContents:         qrisCategoryContents,
	}
}

func (uc *Field) Assign(qris *entities.QRIS, data *entities.Data) error {
	switch data.Tag {
	case uc.qrisTags.Version:
		qris.Version = *data
	case uc.qrisTags.Category:
		qris.Category = *data
	case uc.qrisTags.Acquirer, uc.qrisTags.AcquirerBankTransfer:
		detail, err := uc.acquirerUsecase.Parse(data.Content)
		if err != nil {
			return fmt.Errorf("invalid parse acquirer for content %s", data.Content)
		}
		qris.Acquirer = entities.Acquirer{
			Tag:     data.Tag,
			Content: data.Content,
			Data:    data.Data,
			Detail:  *detail,
		}
	case uc.qrisTags.Switching:
		detail, err := uc.switchingUsecase.Parse(data.Content)
		if err != nil {
			return fmt.Errorf("invalid parse switching for content %s", data.Content)
		}
		qris.Switching = entities.Switching{
			Tag:     data.Tag,
			Content: data.Content,
			Data:    data.Data,
			Detail:  *detail,
		}
	case uc.qrisTags.MerchantCategoryCode:
		qris.MerchantCategoryCode = *data
	case uc.qrisTags.CurrencyCode:
		qris.CurrencyCode = *data
	case uc.qrisTags.PaymentAmount:
		qris.PaymentAmount = *data
	case uc.qrisTags.PaymentFeeCategory:
		qris.PaymentFeeCategory = *data
	case uc.qrisTags.PaymentFeeFixed, uc.qrisTags.PaymentFeePercent:
		qris.PaymentFee = *data
	case uc.qrisTags.CountryCode:
		qris.CountryCode = *data
	case uc.qrisTags.MerchantName:
		qris.MerchantName = *data
	case uc.qrisTags.MerchantCity:
		qris.MerchantCity = *data
	case uc.qrisTags.MerchantPostalCode:
		qris.MerchantPostalCode = *data
	case uc.qrisTags.AdditionalInformation:
		detail, err := uc.additionalInformationUsecase.Parse(data.Content)
		if err != nil {
			return fmt.Errorf("invalid parse additional information for content %s", data.Content)
		}
		qris.AdditionalInformation = entities.AdditionalInformation{
			Tag:     data.Tag,
			Content: data.Content,
			Data:    data.Data,
			Detail:  *detail,
		}
	case uc.qrisTags.CRCCode:
		qris.CRCCode = *data
	default:
		// Ignore unrecognized tags
	}

	return nil
}

func (uc *Field) IsValid(qris *entities.QRIS, errs *[]string) {
	isValidField := func(errs *[]string, tag, message string) {
		if tag == "" {
			*errs = append(*errs, message)
		}
	}

	isValidField(errs, qris.Version.Tag, "Version tag is missing")
	isValidField(errs, qris.Category.Tag, "Category tag is missing")

	if qris.Category.Content != uc.qrisCategoryContents.Static &&
		qris.Category.Content != uc.qrisCategoryContents.Dynamic {
		*errs = append(*errs, "Category content undefined")
	}

	if qris.Acquirer.Tag == "" {
		*errs = append(*errs, "Acquirer tag is missing")
	} else {
		isValidField(errs, qris.Acquirer.Detail.Site.Tag, "Acquirer site tag is missing")
		isValidField(errs, qris.Acquirer.Detail.MPAN.Tag, "Acquirer MPAN tag is missing")
		isValidField(errs, qris.Acquirer.Detail.TerminalID.Tag, "Acquirer terminal id tag is missing")
		if qris.Acquirer.Tag == uc.qrisTags.Acquirer {
			isValidField(errs, qris.Acquirer.Detail.Category.Tag, "Acquirer category tag is missing")

			if qris.Switching.Tag == "" {
				*errs = append(*errs, "Switching tag is missing")
			} else {
				isValidField(errs, qris.Switching.Detail.Site.Tag, "Switching site tag is missing")
				isValidField(errs, qris.Switching.Detail.NMID.Tag, "Switching NMID tag is missing")
				isValidField(errs, qris.Switching.Detail.Category.Tag, "Switching category tag is missing")
			}
		}
	}

	uc.IsValidPaymentFee(qris, errs)

	isValidField(errs, qris.MerchantCategoryCode.Tag, "Merchant category tag is missing")
	isValidField(errs, qris.CurrencyCode.Tag, "Currency code tag is missing")
	isValidField(errs, qris.CountryCode.Tag, "Country code tag is missing")
	isValidField(errs, qris.MerchantName.Tag, "Merchant name tag is missing")
	isValidField(errs, qris.MerchantCity.Tag, "Merchant city tag is missing")
	isValidField(errs, qris.MerchantPostalCode.Tag, "Merchant postal code tag is missing")
	isValidField(errs, qris.CRCCode.Tag, "CRC code tag is missing")
}

func (uc *Field) IsValidPaymentFee(qris *entities.QRIS, errs *[]string) {
	if qris.PaymentFeeCategory.Tag == "" && qris.PaymentFee.Tag != "" {
		*errs = append(*errs, "Payment fee category tag is missing")
	}
	if qris.PaymentFeeCategory.Tag != "" && qris.PaymentFee.Tag == "" {
		*errs = append(*errs, "Payment fee tag is missing")
	}
}
