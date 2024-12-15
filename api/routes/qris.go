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
	dataUsecase := usecases.NewData()
	acquirerUsecase := usecases.NewAcquirer(dataUsecase, config.AcquirerDetailSiteTag, config.AcquirerDetailMPANTag, config.AcquirerDetailTerminalIDTag, config.AcquirerDetailCategoryTag)
	switchingUsecase := usecases.NewSwitching(dataUsecase, config.SwitchingDetailSiteTag, config.SwitchingDetailNMIDTag, config.SwitchingDetailCategoryTag)
	crc16CCITTUsecase := usecases.NewCRC16CCITT()
	qrisUsecase := usecases.NewQRIS(
		acquirerUsecase,
		switchingUsecase,
		dataUsecase,
		crc16CCITTUsecase,
		usecases.QRISTags{
			VersionTag:               config.VersionTag,
			CategoryTag:              config.CategoryTag,
			AcquirerTag:              config.AcquirerTag,
			SwitchingTag:             config.SwitchingTag,
			MerchantCategoryCodeTag:  config.MerchantCategoryCodeTag,
			CurrencyCodeTag:          config.CurrencyCodeTag,
			PaymentAmountTag:         config.PaymentAmountTag,
			PaymentFeeCategoryTag:    config.PaymentFeeCategoryTag,
			CountryCodeTag:           config.CountryCodeTag,
			MerchantNameTag:          config.MerchantNameTag,
			MerchantCityTag:          config.MerchantCityTag,
			MerchantPostalCodeTag:    config.MerchantPostalCodeTag,
			AdditionalInformationTag: config.AdditionalInformationTag,
			CRCCodeTag:               config.CRCCodeTag,
		},
		usecases.QRISCategoryContents{
			StaticContent:  config.CategoryStaticContent,
			DynamicContent: config.CategoryDynamicContent,
		},
		usecases.QRISDynamicPaymentFee{
			CategoryFixedContent:   config.PaymentFeeCategoryFixedContent,
			CategoryPercentContent: config.PaymentFeeCategoryPercentContent,
			FixedTag:               config.PaymentFeeFixedTag,
			PercentTag:             config.PaymentFeePercentTag,
		},
	)
	qrCodeUtil := utils.NewQRCode()
	qrisController := controllers.NewQRIS(qrisUsecase, qrCodeUtil, env.QRCodeSize)
	qrisHandler := handlers.NewQRIS(qrisController)

	group.POST("/extract-static", qrisHandler.ExtractStatic)
	group.POST("/static-to-dynamic", qrisHandler.StaticToDynamic)
}
