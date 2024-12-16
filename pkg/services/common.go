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

func mapQRISDynamicEntityToModel(qrisDynamic *entities.QRISDynamic) *models.QRISDynamic {
	return &models.QRISDynamic{
		Version: models.Data{
			Tag:     qrisDynamic.Version.Tag,
			Content: qrisDynamic.Version.Content,
			Data:    qrisDynamic.Version.Data,
		},
		Category: models.Data{
			Tag:     qrisDynamic.Category.Tag,
			Content: qrisDynamic.Category.Content,
			Data:    qrisDynamic.Category.Data,
		},
		Acquirer: models.Acquirer{
			Tag:     qrisDynamic.Acquirer.Tag,
			Content: qrisDynamic.Acquirer.Content,
			Data:    qrisDynamic.Acquirer.Data,
			Detail: models.AcquirerDetail{
				Site: models.Data{
					Tag:     qrisDynamic.Acquirer.Detail.Site.Tag,
					Content: qrisDynamic.Acquirer.Detail.Site.Content,
					Data:    qrisDynamic.Acquirer.Detail.Site.Data,
				},
				MPAN: models.Data{
					Tag:     qrisDynamic.Acquirer.Detail.MPAN.Tag,
					Content: qrisDynamic.Acquirer.Detail.MPAN.Content,
					Data:    qrisDynamic.Acquirer.Detail.MPAN.Data,
				},
				TerminalID: models.Data{
					Tag:     qrisDynamic.Acquirer.Detail.TerminalID.Tag,
					Content: qrisDynamic.Acquirer.Detail.TerminalID.Content,
					Data:    qrisDynamic.Acquirer.Detail.TerminalID.Data,
				},
				Category: models.Data{
					Tag:     qrisDynamic.Acquirer.Detail.Category.Tag,
					Content: qrisDynamic.Acquirer.Detail.Category.Content,
					Data:    qrisDynamic.Acquirer.Detail.Category.Data,
				},
			},
		},
		Switching: models.Switching{
			Tag:     qrisDynamic.Switching.Tag,
			Content: qrisDynamic.Switching.Content,
			Data:    qrisDynamic.Switching.Data,
			Detail: models.SwitchingDetail{
				Site: models.Data{
					Tag:     qrisDynamic.Switching.Detail.Site.Tag,
					Content: qrisDynamic.Switching.Detail.Site.Content,
					Data:    qrisDynamic.Switching.Detail.Site.Data,
				},
				NMID: models.Data{
					Tag:     qrisDynamic.Switching.Detail.NMID.Tag,
					Content: qrisDynamic.Switching.Detail.NMID.Content,
					Data:    qrisDynamic.Switching.Detail.NMID.Data,
				},
				Category: models.Data{
					Tag:     qrisDynamic.Switching.Detail.Category.Tag,
					Content: qrisDynamic.Switching.Detail.Category.Content,
					Data:    qrisDynamic.Switching.Detail.Category.Data,
				},
			},
		},
		MerchantCategoryCode: models.Data{
			Tag:     qrisDynamic.MerchantCategoryCode.Tag,
			Content: qrisDynamic.MerchantCategoryCode.Content,
			Data:    qrisDynamic.MerchantCategoryCode.Data,
		},
		CurrencyCode: models.Data{
			Tag:     qrisDynamic.CurrencyCode.Tag,
			Content: qrisDynamic.CurrencyCode.Content,
			Data:    qrisDynamic.CurrencyCode.Data,
		},
		PaymentAmount: models.Data{
			Tag:     qrisDynamic.PaymentAmount.Tag,
			Content: qrisDynamic.PaymentAmount.Content,
			Data:    qrisDynamic.PaymentAmount.Data,
		},
		PaymentFeeCategory: models.Data{
			Tag:     qrisDynamic.PaymentFeeCategory.Tag,
			Content: qrisDynamic.PaymentFeeCategory.Content,
			Data:    qrisDynamic.PaymentFeeCategory.Data,
		},
		PaymentFee: models.Data{
			Tag:     qrisDynamic.PaymentFee.Tag,
			Content: qrisDynamic.PaymentFee.Content,
			Data:    qrisDynamic.PaymentFee.Data,
		},
		CountryCode: models.Data{
			Tag:     qrisDynamic.CountryCode.Tag,
			Content: qrisDynamic.CountryCode.Content,
			Data:    qrisDynamic.CountryCode.Data,
		},
		MerchantName: models.Data{
			Tag:     qrisDynamic.MerchantName.Tag,
			Content: qrisDynamic.MerchantName.Content,
			Data:    qrisDynamic.MerchantName.Data,
		},
		MerchantCity: models.Data{
			Tag:     qrisDynamic.MerchantCity.Tag,
			Content: qrisDynamic.MerchantCity.Content,
			Data:    qrisDynamic.MerchantCity.Data,
		},
		MerchantPostalCode: models.Data{
			Tag:     qrisDynamic.MerchantPostalCode.Tag,
			Content: qrisDynamic.MerchantPostalCode.Content,
			Data:    qrisDynamic.MerchantPostalCode.Data,
		},
		AdditionalInformation: models.Data{
			Tag:     qrisDynamic.AdditionalInformation.Tag,
			Content: qrisDynamic.AdditionalInformation.Content,
			Data:    qrisDynamic.AdditionalInformation.Data,
		},
		CRCCode: models.Data{
			Tag:     qrisDynamic.CRCCode.Tag,
			Content: qrisDynamic.CRCCode.Content,
			Data:    qrisDynamic.CRCCode.Data,
		},
	}
}
