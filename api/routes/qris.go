package routes

import (
	"github.com/fyvri/go-qris/api/handlers"
	"github.com/fyvri/go-qris/bootstrap"
	"github.com/fyvri/go-qris/internal/config"
	"github.com/fyvri/go-qris/internal/interface/controllers"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/utils"

	"github.com/gin-gonic/gin"
)

func NewQRISRouter(env *bootstrap.Env, group *gin.RouterGroup) {
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
	paymentFeeUsecase := usecases.NewPaymentFee(qrisTags, qrisPaymentFeeCategoryContents)
	fieldUsecase := usecases.NewField(acquirerUsecase, switchingUsecase, additionalInformationUsecase, qrisTags, qrisCategoryContents)
	crc16CCITTUsecase := usecases.NewCRC16CCITT()

	qrisUsecases := &usecases.QRISUsecases{
		Data:                  dataUsecase,
		Field:                 fieldUsecase,
		PaymentFee:            paymentFeeUsecase,
		AdditionalInformation: additionalInformationUsecase,
		CRC16CCITT:            crc16CCITTUsecase,
	}
	qrisUsecase := usecases.NewQRIS(
		qrisUsecases,
		qrisTags,
		qrisCategoryContents,
		qrisPaymentFeeCategoryContents,
	)
	qrCodeUtil := utils.NewQRCode()
	inputUtil := utils.NewInput()
	qrisController := controllers.NewQRIS(inputUtil, qrCodeUtil, qrisUsecase, env.QRCodeSize)
	qrisHandler := handlers.NewQRIS(qrisController)

	group.POST("/parse", qrisHandler.Parse)
	group.POST("/convert", qrisHandler.Convert)
	group.POST("/is-valid", qrisHandler.IsValid)
}
