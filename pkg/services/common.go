package services

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/pkg/models"
)

func mapQRISEntityToModel(qris *entities.QRIS) *models.QRIS {
	return &models.QRIS{
		Version: models.Data{
			Tag:     qris.Version.Tag,
			Content: qris.Version.Content,
			Data:    qris.Version.Data,
		},
		Category: models.Data{
			Tag:     qris.Category.Tag,
			Content: qris.Category.Content,
			Data:    qris.Category.Data,
		},
		Acquirer: models.Acquirer{
			Tag:     qris.Acquirer.Tag,
			Content: qris.Acquirer.Content,
			Data:    qris.Acquirer.Data,
			Detail: models.AcquirerDetail{
				Site: models.Data{
					Tag:     qris.Acquirer.Detail.Site.Tag,
					Content: qris.Acquirer.Detail.Site.Content,
					Data:    qris.Acquirer.Detail.Site.Data,
				},
				MPAN: models.Data{
					Tag:     qris.Acquirer.Detail.MPAN.Tag,
					Content: qris.Acquirer.Detail.MPAN.Content,
					Data:    qris.Acquirer.Detail.MPAN.Data,
				},
				TerminalID: models.Data{
					Tag:     qris.Acquirer.Detail.TerminalID.Tag,
					Content: qris.Acquirer.Detail.TerminalID.Content,
					Data:    qris.Acquirer.Detail.TerminalID.Data,
				},
				Category: models.Data{
					Tag:     qris.Acquirer.Detail.Category.Tag,
					Content: qris.Acquirer.Detail.Category.Content,
					Data:    qris.Acquirer.Detail.Category.Data,
				},
			},
		},
		Switching: models.Switching{
			Tag:     qris.Switching.Tag,
			Content: qris.Switching.Content,
			Data:    qris.Switching.Data,
			Detail: models.SwitchingDetail{
				Site: models.Data{
					Tag:     qris.Switching.Detail.Site.Tag,
					Content: qris.Switching.Detail.Site.Content,
					Data:    qris.Switching.Detail.Site.Data,
				},
				NMID: models.Data{
					Tag:     qris.Switching.Detail.NMID.Tag,
					Content: qris.Switching.Detail.NMID.Content,
					Data:    qris.Switching.Detail.NMID.Data,
				},
				Category: models.Data{
					Tag:     qris.Switching.Detail.Category.Tag,
					Content: qris.Switching.Detail.Category.Content,
					Data:    qris.Switching.Detail.Category.Data,
				},
			},
		},
		MerchantCategoryCode: models.Data{
			Tag:     qris.MerchantCategoryCode.Tag,
			Content: qris.MerchantCategoryCode.Content,
			Data:    qris.MerchantCategoryCode.Data,
		},
		CurrencyCode: models.Data{
			Tag:     qris.CurrencyCode.Tag,
			Content: qris.CurrencyCode.Content,
			Data:    qris.CurrencyCode.Data,
		},
		PaymentAmount: models.Data{
			Tag:     qris.PaymentAmount.Tag,
			Content: qris.PaymentAmount.Content,
			Data:    qris.PaymentAmount.Data,
		},
		PaymentFeeCategory: models.Data{
			Tag:     qris.PaymentFeeCategory.Tag,
			Content: qris.PaymentFeeCategory.Content,
			Data:    qris.PaymentFeeCategory.Data,
		},
		PaymentFee: models.Data{
			Tag:     qris.PaymentFee.Tag,
			Content: qris.PaymentFee.Content,
			Data:    qris.PaymentFee.Data,
		},
		CountryCode: models.Data{
			Tag:     qris.CountryCode.Tag,
			Content: qris.CountryCode.Content,
			Data:    qris.CountryCode.Data,
		},
		MerchantName: models.Data{
			Tag:     qris.MerchantName.Tag,
			Content: qris.MerchantName.Content,
			Data:    qris.MerchantName.Data,
		},
		MerchantCity: models.Data{
			Tag:     qris.MerchantCity.Tag,
			Content: qris.MerchantCity.Content,
			Data:    qris.MerchantCity.Data,
		},
		MerchantPostalCode: models.Data{
			Tag:     qris.MerchantPostalCode.Tag,
			Content: qris.MerchantPostalCode.Content,
			Data:    qris.MerchantPostalCode.Data,
		},
		AdditionalInformation: models.Data{
			Tag:     qris.AdditionalInformation.Tag,
			Content: qris.AdditionalInformation.Content,
			Data:    qris.AdditionalInformation.Data,
		},
		CRCCode: models.Data{
			Tag:     qris.CRCCode.Tag,
			Content: qris.CRCCode.Content,
			Data:    qris.CRCCode.Data,
		},
	}
}

func mapQRISModelToEntity(qris *models.QRIS) *entities.QRIS {
	return &entities.QRIS{
		Version: entities.Data{
			Tag:     qris.Version.Tag,
			Content: qris.Version.Content,
			Data:    qris.Version.Data,
		},
		Category: entities.Data{
			Tag:     qris.Category.Tag,
			Content: qris.Category.Content,
			Data:    qris.Category.Data,
		},
		Acquirer: entities.Acquirer{
			Tag:     qris.Acquirer.Tag,
			Content: qris.Acquirer.Content,
			Data:    qris.Acquirer.Data,
			Detail: entities.AcquirerDetail{
				Site: entities.Data{
					Tag:     qris.Acquirer.Detail.Site.Tag,
					Content: qris.Acquirer.Detail.Site.Content,
					Data:    qris.Acquirer.Detail.Site.Data,
				},
				MPAN: entities.Data{
					Tag:     qris.Acquirer.Detail.MPAN.Tag,
					Content: qris.Acquirer.Detail.MPAN.Content,
					Data:    qris.Acquirer.Detail.MPAN.Data,
				},
				TerminalID: entities.Data{
					Tag:     qris.Acquirer.Detail.TerminalID.Tag,
					Content: qris.Acquirer.Detail.TerminalID.Content,
					Data:    qris.Acquirer.Detail.TerminalID.Data,
				},
				Category: entities.Data{
					Tag:     qris.Acquirer.Detail.Category.Tag,
					Content: qris.Acquirer.Detail.Category.Content,
					Data:    qris.Acquirer.Detail.Category.Data,
				},
			},
		},
		Switching: entities.Switching{
			Tag:     qris.Switching.Tag,
			Content: qris.Switching.Content,
			Data:    qris.Switching.Data,
			Detail: entities.SwitchingDetail{
				Site: entities.Data{
					Tag:     qris.Switching.Detail.Site.Tag,
					Content: qris.Switching.Detail.Site.Content,
					Data:    qris.Switching.Detail.Site.Data,
				},
				NMID: entities.Data{
					Tag:     qris.Switching.Detail.NMID.Tag,
					Content: qris.Switching.Detail.NMID.Content,
					Data:    qris.Switching.Detail.NMID.Data,
				},
				Category: entities.Data{
					Tag:     qris.Switching.Detail.Category.Tag,
					Content: qris.Switching.Detail.Category.Content,
					Data:    qris.Switching.Detail.Category.Data,
				},
			},
		},
		MerchantCategoryCode: entities.Data{
			Tag:     qris.MerchantCategoryCode.Tag,
			Content: qris.MerchantCategoryCode.Content,
			Data:    qris.MerchantCategoryCode.Data,
		},
		CurrencyCode: entities.Data{
			Tag:     qris.CurrencyCode.Tag,
			Content: qris.CurrencyCode.Content,
			Data:    qris.CurrencyCode.Data,
		},
		PaymentAmount: entities.Data{
			Tag:     qris.PaymentAmount.Tag,
			Content: qris.PaymentAmount.Content,
			Data:    qris.PaymentAmount.Data,
		},
		PaymentFeeCategory: entities.Data{
			Tag:     qris.PaymentFeeCategory.Tag,
			Content: qris.PaymentFeeCategory.Content,
			Data:    qris.PaymentFeeCategory.Data,
		},
		PaymentFee: entities.Data{
			Tag:     qris.PaymentFee.Tag,
			Content: qris.PaymentFee.Content,
			Data:    qris.PaymentFee.Data,
		},
		CountryCode: entities.Data{
			Tag:     qris.CountryCode.Tag,
			Content: qris.CountryCode.Content,
			Data:    qris.CountryCode.Data,
		},
		MerchantName: entities.Data{
			Tag:     qris.MerchantName.Tag,
			Content: qris.MerchantName.Content,
			Data:    qris.MerchantName.Data,
		},
		MerchantCity: entities.Data{
			Tag:     qris.MerchantCity.Tag,
			Content: qris.MerchantCity.Content,
			Data:    qris.MerchantCity.Data,
		},
		MerchantPostalCode: entities.Data{
			Tag:     qris.MerchantPostalCode.Tag,
			Content: qris.MerchantPostalCode.Content,
			Data:    qris.MerchantPostalCode.Data,
		},
		AdditionalInformation: entities.Data{
			Tag:     qris.AdditionalInformation.Tag,
			Content: qris.AdditionalInformation.Content,
			Data:    qris.AdditionalInformation.Data,
		},
		CRCCode: entities.Data{
			Tag:     qris.CRCCode.Tag,
			Content: qris.CRCCode.Content,
			Data:    qris.CRCCode.Data,
		},
	}
}
