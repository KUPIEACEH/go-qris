package services

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/models"
)

var (
	expectedButGotMessage             = "Expected %v = %v, but got = %v"
	expectedErrorButGotMessage        = "Expected %v error = %v, but got = %v"
	expectedTypeAssertionErrorMessage = "Expected type assertion error, but got = %v"
	expectedReturnNonNil              = "Expected %v to return a non-nil %v"

	testVersionTag                       = "00"
	testCategoryTag                      = "01"
	testAcquirerTag                      = "26"
	testAcquirerBankTransferTag          = "40"
	testSwitchingTag                     = "51"
	testMerchantCategoryCodeTag          = "52"
	testCurrencyCodeTag                  = "53"
	testPaymentAmountTag                 = "54"
	testPaymentFeeCategoryTag            = "55"
	testPaymentFeeFixedTag               = "56"
	testPaymentFeePercentTag             = "57"
	testCountryCodeTag                   = "58"
	testMerchantNameTag                  = "59"
	testMerchantCityTag                  = "60"
	testMerchantPostalCodeTag            = "61"
	testAdditionalInformationTag         = "62"
	testCRCCodeTag                       = "63"
	testAcquirerDetailSiteTag            = "00"
	testAcquirerDetailMPANTag            = "01"
	testAcquirerDetailTerminalIDTag      = "02"
	testAcquirerDetailCategoryTag        = "03"
	testSwitchingDetailSiteTag           = "00"
	testSwitchingDetailNMIDTag           = "02"
	testSwitchingDetailCategoryTag       = "03"
	testCategoryStaticContent            = "11"
	testCategoryDynamicContent           = "12"
	testPaymentFeeCategoryFixedContent   = "02"
	testPaymentFeeCategoryPercentContent = "03"

	testQRISTags = usecases.QRISTags{
		VersionTag:               testVersionTag,
		CategoryTag:              testCategoryTag,
		AcquirerTag:              testAcquirerTag,
		AcquirerBankTransferTag:  testAcquirerBankTransferTag,
		SwitchingTag:             testSwitchingTag,
		MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
		CurrencyCodeTag:          testCurrencyCodeTag,
		PaymentAmountTag:         testPaymentAmountTag,
		PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
		PaymentFeeFixedTag:       testPaymentFeeFixedTag,
		PaymentFeePercentTag:     testPaymentFeePercentTag,
		CountryCodeTag:           testCountryCodeTag,
		MerchantNameTag:          testMerchantNameTag,
		MerchantCityTag:          testMerchantCityTag,
		MerchantPostalCodeTag:    testMerchantPostalCodeTag,
		AdditionalInformationTag: testAdditionalInformationTag,
		CRCCodeTag:               testCRCCodeTag,
	}
	testQRISCategoryContents = usecases.QRISCategoryContents{
		Static:  testCategoryStaticContent,
		Dynamic: testCategoryDynamicContent,
	}
	testQRISPaymentFeeCategoryContents = usecases.QRISPaymentFeeCategoryContents{
		Fixed:   testPaymentFeeCategoryFixedContent,
		Percent: testPaymentFeeCategoryPercentContent,
	}

	testAcquirerDetail = entities.AcquirerDetail{
		Site: entities.Data{
			Tag:     testAcquirerDetailSiteTag,
			Content: "COM.MEMBASUH.WWW",
			Data:    testAcquirerDetailSiteTag + "16COM.MEMBASUH.WWW",
		},
		MPAN: entities.Data{
			Tag:     testAcquirerDetailMPANTag,
			Content: "936009153022591481",
			Data:    testAcquirerDetailMPANTag + "18936009153022591481",
		},
		TerminalID: entities.Data{
			Tag:     testAcquirerDetailTerminalIDTag,
			Content: "022591481",
			Data:    testAcquirerDetailTerminalIDTag + "09022591481",
		},
		Category: entities.Data{
			Tag:     testAcquirerDetailCategoryTag,
			Content: "UMI",
			Data:    testAcquirerDetailCategoryTag + "03UMI",
		},
	}

	testSwitchingDetail = entities.SwitchingDetail{
		Site: entities.Data{
			Tag:     testSwitchingDetailSiteTag,
			Content: "ID.CO.QRIS.WWW",
			Data:    testSwitchingDetailSiteTag + "14ID.CO.QRIS.WWW",
		},
		NMID: entities.Data{
			Tag:     testSwitchingDetailNMIDTag,
			Content: "ID1020017611473",
			Data:    testSwitchingDetailNMIDTag + "15ID1020017611473",
		},
		Category: entities.Data{
			Tag:     testSwitchingDetailCategoryTag,
			Content: "UMI",
			Data:    testSwitchingDetailCategoryTag + "03UMI",
		},
	}

	testQRISEntity = entities.QRIS{
		Version: entities.Data{
			Tag:     testVersionTag,
			Content: "01",
			Data:    testVersionTag + "0201",
		},
		Category: entities.Data{
			Tag:     testCategoryTag,
			Content: testCategoryStaticContent,
			Data:    testCategoryTag + "02" + testCategoryStaticContent,
		},
		Acquirer: entities.Acquirer{
			Tag:     testAcquirerTag,
			Content: testAcquirerDetail.Site.Data + testAcquirerDetail.MPAN.Data + testAcquirerDetail.TerminalID.Data + testAcquirerDetail.Category.Data,
			Data:    testAcquirerTag + "62" + testAcquirerDetail.Site.Data + testAcquirerDetail.MPAN.Data + testAcquirerDetail.TerminalID.Data + testAcquirerDetail.Category.Data,
			Detail:  testAcquirerDetail,
		},
		Switching: entities.Switching{
			Tag:     testSwitchingTag,
			Content: testSwitchingDetail.Site.Data + testSwitchingDetail.NMID.Data + testSwitchingDetail.Category.Data,
			Data:    testSwitchingTag + "44" + testSwitchingDetail.Site.Data + testSwitchingDetail.NMID.Data + testSwitchingDetail.Category.Data,
			Detail:  testSwitchingDetail,
		},
		MerchantCategoryCode: entities.Data{
			Tag:     testMerchantCategoryCodeTag,
			Content: "4829",
			Data:    testMerchantCategoryCodeTag + "044829",
		},
		CurrencyCode: entities.Data{
			Tag:     testCurrencyCodeTag,
			Content: "360",
			Data:    testCurrencyCodeTag + "03360",
		},
		PaymentAmount: entities.Data{
			Tag:     "",
			Content: "",
			Data:    "",
		},
		PaymentFeeCategory: entities.Data{
			Tag:     "",
			Content: "",
			Data:    "",
		},
		PaymentFee: entities.Data{
			Tag:     "",
			Content: "",
			Data:    "",
		},
		CountryCode: entities.Data{
			Tag:     testCountryCodeTag,
			Content: "ID",
			Data:    testCountryCodeTag + "02ID",
		},
		MerchantName: entities.Data{
			Tag:     testMerchantNameTag,
			Content: "Sintas Store",
			Data:    testMerchantNameTag + "12Sintas Store",
		},
		MerchantCity: entities.Data{
			Tag:     testMerchantCityTag,
			Content: "Kota Yogyakarta",
			Data:    testMerchantCityTag + "15Kota Yogyakarta",
		},
		MerchantPostalCode: entities.Data{
			Tag:     testMerchantPostalCodeTag,
			Content: "55000",
			Data:    testMerchantPostalCodeTag + "0555000",
		},
		AdditionalInformation: entities.Data{
			Tag:     testAdditionalInformationTag,
			Content: "0703A01",
			Data:    testAdditionalInformationTag + "070703A01",
		},
		CRCCode: entities.Data{
			Tag:     testCRCCodeTag,
			Content: "1FA2",
			Data:    testCRCCodeTag + "041FA2",
		},
	}
	testQRISEntityString = testQRISEntity.Version.Data +
		testQRISEntity.Category.Data +
		testQRISEntity.Acquirer.Data +
		testQRISEntity.Switching.Data +
		testQRISEntity.MerchantCategoryCode.Data +
		testQRISEntity.CurrencyCode.Data +
		testQRISEntity.CountryCode.Data +
		testQRISEntity.MerchantName.Data +
		testQRISEntity.MerchantCity.Data +
		testQRISEntity.MerchantPostalCode.Data +
		testQRISEntity.AdditionalInformation.Data +
		testQRISEntity.CRCCode.Data

	testPaymentAmountValue        = 1337
	testPaymentFeeValue           = 666
	testMerchantCityContent       = "New Merchant City"
	testMerchantPostalCodeContent = "55181"
	testQRISEntityModified        = entities.QRIS{
		Version: testQRISEntity.Version,
		Category: entities.Data{
			Tag:     testCategoryTag,
			Content: testCategoryDynamicContent,
			Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
		},
		Acquirer:             testQRISEntity.Acquirer,
		Switching:            testQRISEntity.Switching,
		MerchantCategoryCode: testQRISEntity.MerchantCategoryCode,
		CurrencyCode:         testQRISEntity.CurrencyCode,
		PaymentAmount: entities.Data{
			Tag:     testPaymentAmountTag,
			Content: fmt.Sprintf("%d", testPaymentAmountValue),
			Data:    testPaymentAmountTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", testPaymentAmountValue))) + fmt.Sprintf("%d", testPaymentAmountValue),
		},
		PaymentFeeCategory: entities.Data{
			Tag:     testPaymentFeeCategoryTag,
			Content: testPaymentFeeCategoryFixedContent,
			Data:    testPaymentFeeCategoryTag + fmt.Sprintf("%02d", len(testPaymentFeeCategoryFixedContent)) + testPaymentFeeCategoryFixedContent,
		},
		PaymentFee: entities.Data{
			Tag:     testPaymentFeeFixedTag,
			Content: fmt.Sprintf("%d", testPaymentFeeValue),
			Data:    testPaymentFeeFixedTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", testPaymentFeeValue))) + fmt.Sprintf("%d", testPaymentFeeValue),
		},
		CountryCode:  testQRISEntity.CountryCode,
		MerchantName: testQRISEntity.MerchantName,
		MerchantCity: entities.Data{
			Tag:     testQRISEntity.MerchantCity.Tag,
			Content: testMerchantCityContent,
			Data:    testQRISEntity.MerchantCity.Tag + fmt.Sprintf("%02d", len(testMerchantCityContent)) + testMerchantCityContent,
		},
		MerchantPostalCode: entities.Data{
			Tag:     testQRISEntity.MerchantPostalCode.Tag,
			Content: testMerchantPostalCodeContent,
			Data:    testQRISEntity.MerchantPostalCode.Tag + fmt.Sprintf("%02d", len(testMerchantPostalCodeContent)) + testMerchantPostalCodeContent,
		},
		AdditionalInformation: testQRISEntity.AdditionalInformation,
		CRCCode:               testQRISEntity.CRCCode,
	}
	testQRISEntityModifiedString = testQRISEntityModified.Version.Data +
		testQRISEntityModified.Category.Data +
		testQRISEntityModified.Acquirer.Data +
		testQRISEntityModified.Switching.Data +
		testQRISEntityModified.MerchantCategoryCode.Data +
		testQRISEntityModified.CurrencyCode.Data +
		testQRISEntityModified.PaymentAmount.Data +
		testQRISEntityModified.PaymentFeeCategory.Data +
		testQRISEntityModified.PaymentFee.Data +
		testQRISEntityModified.CountryCode.Data +
		testQRISEntityModified.MerchantName.Data +
		testQRISEntityModified.MerchantCity.Data +
		testQRISEntityModified.MerchantPostalCode.Data +
		testQRISEntityModified.AdditionalInformation.Data +
		testQRISEntityModified.CRCCode.Data

	testQRISModel = models.QRIS{
		Version: models.Data{
			Tag:     testQRISEntity.Version.Tag,
			Content: testQRISEntity.Version.Content,
			Data:    testQRISEntity.Version.Data,
		},
		Category: models.Data{
			Tag:     testQRISEntity.Category.Tag,
			Content: testQRISEntity.Category.Content,
			Data:    testQRISEntity.Category.Data,
		},
		Acquirer: models.Acquirer{
			Tag:     testQRISEntity.Acquirer.Tag,
			Content: testQRISEntity.Acquirer.Content,
			Data:    testQRISEntity.Acquirer.Data,
			Detail: models.AcquirerDetail{
				Site: models.Data{
					Tag:     testQRISEntity.Acquirer.Detail.Site.Tag,
					Content: testQRISEntity.Acquirer.Detail.Site.Content,
					Data:    testQRISEntity.Acquirer.Detail.Site.Data,
				},
				MPAN: models.Data{
					Tag:     testQRISEntity.Acquirer.Detail.MPAN.Tag,
					Content: testQRISEntity.Acquirer.Detail.MPAN.Content,
					Data:    testQRISEntity.Acquirer.Detail.MPAN.Data,
				},
				TerminalID: models.Data{
					Tag:     testQRISEntity.Acquirer.Detail.TerminalID.Tag,
					Content: testQRISEntity.Acquirer.Detail.TerminalID.Content,
					Data:    testQRISEntity.Acquirer.Detail.TerminalID.Data,
				},
				Category: models.Data{
					Tag:     testQRISEntity.Acquirer.Detail.Category.Tag,
					Content: testQRISEntity.Acquirer.Detail.Category.Content,
					Data:    testQRISEntity.Acquirer.Detail.Category.Data,
				},
			},
		},
		Switching: models.Switching{
			Tag:     testQRISEntity.Switching.Tag,
			Content: testQRISEntity.Switching.Content,
			Data:    testQRISEntity.Switching.Data,
			Detail: models.SwitchingDetail{
				Site: models.Data{
					Tag:     testQRISEntity.Switching.Detail.Site.Tag,
					Content: testQRISEntity.Switching.Detail.Site.Content,
					Data:    testQRISEntity.Switching.Detail.Site.Data,
				},
				NMID: models.Data{
					Tag:     testQRISEntity.Switching.Detail.NMID.Tag,
					Content: testQRISEntity.Switching.Detail.NMID.Content,
					Data:    testQRISEntity.Switching.Detail.NMID.Data,
				},
				Category: models.Data{
					Tag:     testQRISEntity.Switching.Detail.Category.Tag,
					Content: testQRISEntity.Switching.Detail.Category.Content,
					Data:    testQRISEntity.Switching.Detail.Category.Data,
				},
			},
		},
		MerchantCategoryCode: models.Data{
			Tag:     testQRISEntity.MerchantCategoryCode.Tag,
			Content: testQRISEntity.MerchantCategoryCode.Content,
			Data:    testQRISEntity.MerchantCategoryCode.Data,
		},
		CurrencyCode: models.Data{
			Tag:     testQRISEntity.CurrencyCode.Tag,
			Content: testQRISEntity.CurrencyCode.Content,
			Data:    testQRISEntity.CurrencyCode.Data,
		},
		PaymentAmount: models.Data{
			Tag:     testQRISEntity.PaymentAmount.Tag,
			Content: testQRISEntity.PaymentAmount.Content,
			Data:    testQRISEntity.PaymentAmount.Data,
		},
		PaymentFeeCategory: models.Data{
			Tag:     testQRISEntity.PaymentFeeCategory.Tag,
			Content: testQRISEntity.PaymentFeeCategory.Content,
			Data:    testQRISEntity.PaymentFeeCategory.Data,
		},
		PaymentFee: models.Data{
			Tag:     testQRISEntity.PaymentFee.Tag,
			Content: testQRISEntity.PaymentFee.Content,
			Data:    testQRISEntity.PaymentFee.Data,
		},
		CountryCode: models.Data{
			Tag:     testQRISEntity.CountryCode.Tag,
			Content: testQRISEntity.CountryCode.Content,
			Data:    testQRISEntity.CountryCode.Data,
		},
		MerchantName: models.Data{
			Tag:     testQRISEntity.MerchantName.Tag,
			Content: testQRISEntity.MerchantName.Content,
			Data:    testQRISEntity.MerchantName.Data,
		},
		MerchantCity: models.Data{
			Tag:     testQRISEntity.MerchantCity.Tag,
			Content: testQRISEntity.MerchantCity.Content,
			Data:    testQRISEntity.MerchantCity.Data,
		},
		MerchantPostalCode: models.Data{
			Tag:     testQRISEntity.MerchantPostalCode.Tag,
			Content: testQRISEntity.MerchantPostalCode.Content,
			Data:    testQRISEntity.MerchantPostalCode.Data,
		},
		AdditionalInformation: models.Data{
			Tag:     testQRISEntity.AdditionalInformation.Tag,
			Content: testQRISEntity.AdditionalInformation.Content,
			Data:    testQRISEntity.AdditionalInformation.Data,
		},
		CRCCode: models.Data{
			Tag:     testQRISEntity.CRCCode.Tag,
			Content: testQRISEntity.CRCCode.Content,
			Data:    testQRISEntity.CRCCode.Data,
		},
	}
	testQRISModelModified = models.QRIS{
		Version: models.Data{
			Tag:     testQRISEntityModified.Version.Tag,
			Content: testQRISEntityModified.Version.Content,
			Data:    testQRISEntityModified.Version.Data,
		},
		Category: models.Data{
			Tag:     testQRISEntityModified.Category.Tag,
			Content: testQRISEntityModified.Category.Content,
			Data:    testQRISEntityModified.Category.Data,
		},
		Acquirer: models.Acquirer{
			Tag:     testQRISEntityModified.Acquirer.Tag,
			Content: testQRISEntityModified.Acquirer.Content,
			Data:    testQRISEntityModified.Acquirer.Data,
			Detail: models.AcquirerDetail{
				Site: models.Data{
					Tag:     testQRISEntityModified.Acquirer.Detail.Site.Tag,
					Content: testQRISEntityModified.Acquirer.Detail.Site.Content,
					Data:    testQRISEntityModified.Acquirer.Detail.Site.Data,
				},
				MPAN: models.Data{
					Tag:     testQRISEntityModified.Acquirer.Detail.MPAN.Tag,
					Content: testQRISEntityModified.Acquirer.Detail.MPAN.Content,
					Data:    testQRISEntityModified.Acquirer.Detail.MPAN.Data,
				},
				TerminalID: models.Data{
					Tag:     testQRISEntityModified.Acquirer.Detail.TerminalID.Tag,
					Content: testQRISEntityModified.Acquirer.Detail.TerminalID.Content,
					Data:    testQRISEntityModified.Acquirer.Detail.TerminalID.Data,
				},
				Category: models.Data{
					Tag:     testQRISEntityModified.Acquirer.Detail.Category.Tag,
					Content: testQRISEntityModified.Acquirer.Detail.Category.Content,
					Data:    testQRISEntityModified.Acquirer.Detail.Category.Data,
				},
			},
		},
		Switching: models.Switching{
			Tag:     testQRISEntityModified.Switching.Tag,
			Content: testQRISEntityModified.Switching.Content,
			Data:    testQRISEntityModified.Switching.Data,
			Detail: models.SwitchingDetail{
				Site: models.Data{
					Tag:     testQRISEntityModified.Switching.Detail.Site.Tag,
					Content: testQRISEntityModified.Switching.Detail.Site.Content,
					Data:    testQRISEntityModified.Switching.Detail.Site.Data,
				},
				NMID: models.Data{
					Tag:     testQRISEntityModified.Switching.Detail.NMID.Tag,
					Content: testQRISEntityModified.Switching.Detail.NMID.Content,
					Data:    testQRISEntityModified.Switching.Detail.NMID.Data,
				},
				Category: models.Data{
					Tag:     testQRISEntityModified.Switching.Detail.Category.Tag,
					Content: testQRISEntityModified.Switching.Detail.Category.Content,
					Data:    testQRISEntityModified.Switching.Detail.Category.Data,
				},
			},
		},
		MerchantCategoryCode: models.Data{
			Tag:     testQRISEntityModified.MerchantCategoryCode.Tag,
			Content: testQRISEntityModified.MerchantCategoryCode.Content,
			Data:    testQRISEntityModified.MerchantCategoryCode.Data,
		},
		CurrencyCode: models.Data{
			Tag:     testQRISEntityModified.CurrencyCode.Tag,
			Content: testQRISEntityModified.CurrencyCode.Content,
			Data:    testQRISEntityModified.CurrencyCode.Data,
		},
		PaymentAmount: models.Data{
			Tag:     testQRISEntityModified.PaymentAmount.Tag,
			Content: testQRISEntityModified.PaymentAmount.Content,
			Data:    testQRISEntityModified.PaymentAmount.Data,
		},
		PaymentFeeCategory: models.Data{
			Tag:     testQRISEntityModified.PaymentFeeCategory.Tag,
			Content: testQRISEntityModified.PaymentFeeCategory.Content,
			Data:    testQRISEntityModified.PaymentFeeCategory.Data,
		},
		PaymentFee: models.Data{
			Tag:     testQRISEntityModified.PaymentFee.Tag,
			Content: testQRISEntityModified.PaymentFee.Content,
			Data:    testQRISEntityModified.PaymentFee.Data,
		},
		CountryCode: models.Data{
			Tag:     testQRISEntityModified.CountryCode.Tag,
			Content: testQRISEntityModified.CountryCode.Content,
			Data:    testQRISEntityModified.CountryCode.Data,
		},
		MerchantName: models.Data{
			Tag:     testQRISEntityModified.MerchantName.Tag,
			Content: testQRISEntityModified.MerchantName.Content,
			Data:    testQRISEntityModified.MerchantName.Data,
		},
		MerchantCity: models.Data{
			Tag:     testQRISEntityModified.MerchantCity.Tag,
			Content: testQRISEntityModified.MerchantCity.Content,
			Data:    testQRISEntityModified.MerchantCity.Data,
		},
		MerchantPostalCode: models.Data{
			Tag:     testQRISEntityModified.MerchantPostalCode.Tag,
			Content: testQRISEntityModified.MerchantPostalCode.Content,
			Data:    testQRISEntityModified.MerchantPostalCode.Data,
		},
		AdditionalInformation: models.Data{
			Tag:     testQRISEntityModified.AdditionalInformation.Tag,
			Content: testQRISEntityModified.AdditionalInformation.Content,
			Data:    testQRISEntityModified.AdditionalInformation.Data,
		},
		CRCCode: models.Data{
			Tag:     testQRISEntityModified.CRCCode.Tag,
			Content: testQRISEntityModified.CRCCode.Content,
			Data:    testQRISEntityModified.CRCCode.Data,
		},
	}
	testQRISModelModifiedString = testQRISModelModified.Version.Data +
		testQRISModelModified.Category.Data +
		testQRISModelModified.Acquirer.Data +
		testQRISModelModified.Switching.Data +
		testQRISModelModified.MerchantCategoryCode.Data +
		testQRISModelModified.CurrencyCode.Data +
		testQRISModelModified.PaymentAmount.Data +
		testQRISModelModified.PaymentFeeCategory.Data +
		testQRISModelModified.PaymentFee.Data +
		testQRISModelModified.CountryCode.Data +
		testQRISModelModified.MerchantName.Data +
		testQRISModelModified.MerchantCity.Data +
		testQRISModelModified.MerchantPostalCode.Data +
		testQRISModelModified.AdditionalInformation.Data +
		testQRISModelModified.CRCCode.Data
)

type mockCRC16CCITTUsecase struct {
	GenerateCodeFunc func(code string) string
}

func (m *mockCRC16CCITTUsecase) GenerateCode(code string) string {
	if m.GenerateCodeFunc != nil {
		return m.GenerateCodeFunc(code)
	}
	return ""
}

type mockQRISUsecase struct {
	ParseFunc    func(qrString string) (*entities.QRIS, error, *[]string)
	IsValidFunc  func(qris *entities.QRIS) bool
	ModifyFunc   func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS
	ToStringFunc func(qris *entities.QRIS) string
}

func (m *mockQRISUsecase) Parse(qrString string) (*entities.QRIS, error, *[]string) {
	if m.ParseFunc != nil {
		return m.ParseFunc(qrString)
	}
	return nil, nil, nil
}

func (m *mockQRISUsecase) IsValid(qris *entities.QRIS) bool {
	if m.IsValidFunc != nil {
		return m.IsValidFunc(qris)
	}
	return false
}

func (m *mockQRISUsecase) Modify(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
	if m.ModifyFunc != nil {
		return m.ModifyFunc(qris, merchantCityValue, merchantPostalCodeValue, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)
	}
	return nil
}

func (m *mockQRISUsecase) ToString(qris *entities.QRIS) string {
	if m.ToStringFunc != nil {
		return m.ToStringFunc(qris)
	}
	return ""
}

type mockInputUtil struct {
	SanitizeFunc func(input string) string
}

func (m *mockInputUtil) Sanitize(input string) string {
	if m.SanitizeFunc != nil {
		return m.SanitizeFunc(input)
	}
	return ""
}
