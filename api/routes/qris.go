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
	acquirerUsecase := usecases.NewAcquirer(dataUsecase, config.AcquirerDetailSiteTag, config.AcquirerDetailMPANTag, config.AcquirerDetailTerminalIDTag, config.AcquirerDetailCategoryTag)
	switchingUsecase := usecases.NewSwitching(dataUsecase, config.SwitchingDetailSiteTag, config.SwitchingDetailNMIDTag, config.SwitchingDetailCategoryTag)
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
	qrCodeUtil := utils.NewQRCode()
	inputUtil := utils.NewInput()
	qrisController := controllers.NewQRIS(inputUtil, qrCodeUtil, qrisUsecase, env.QRCodeSize)
	qrisHandler := handlers.NewQRIS(qrisController)

	group.POST("/parse", qrisHandler.Parse)
	group.POST("/convert", qrisHandler.Convert)
	group.POST("/is-valid", qrisHandler.IsValid)
}
