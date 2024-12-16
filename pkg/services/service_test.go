package services

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
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

	testQRIS = entities.QRIS{
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
			Tag:     testPaymentAmountTag,
			Content: "1337",
			Data:    testPaymentAmountTag + "041337",
		},
		PaymentFeeCategory: entities.Data{
			Tag:     testPaymentFeeCategoryTag,
			Content: testPaymentFeeCategoryFixedContent,
			Data:    testPaymentFeeCategoryTag + "02" + testPaymentFeeCategoryFixedContent,
		},
		PaymentFee: entities.Data{
			Tag:     testPaymentFeeFixedTag,
			Content: "666",
			Data:    testPaymentFeeFixedTag + "03666",
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
	testQRISModel = *mapQRISEntityToModel(&testQRIS)

	testMerchantCityContent       = "New Merchant City"
	testMerchantPostalCodeContent = "55181"
	testPaymentAmountValue        = uint32(1337)
	testQRISString                = testQRIS.Version.Data +
		testQRIS.Category.Data +
		testQRIS.Acquirer.Data +
		testQRIS.Switching.Data +
		testQRIS.MerchantCategoryCode.Data +
		testQRIS.CurrencyCode.Data +
		testQRIS.CountryCode.Data +
		testQRIS.MerchantName.Data +
		testQRIS.MerchantCity.Data +
		testQRIS.MerchantPostalCode.Data +
		testQRIS.AdditionalInformation.Data +
		testQRIS.CRCCode.Data

	testMerchantCity = entities.Data{
		Tag:     testQRIS.MerchantCity.Tag,
		Content: testMerchantCityContent,
		Data:    testQRIS.MerchantCity.Tag + fmt.Sprintf("%02d", len(testMerchantCityContent)) + testMerchantCityContent,
	}
	testMerchantPostalCode = entities.Data{
		Tag:     testQRIS.MerchantPostalCode.Tag,
		Content: testMerchantPostalCodeContent,
		Data:    testQRIS.MerchantPostalCode.Tag + fmt.Sprintf("%02d", len(testMerchantPostalCodeContent)) + testMerchantPostalCodeContent,
	}
	testQRISDynamic = entities.QRISDynamic{
		Version: testQRIS.Version,
		Category: entities.Data{
			Tag:     testCategoryTag,
			Content: testCategoryDynamicContent,
			Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
		},
		Acquirer:             testQRIS.Acquirer,
		Switching:            testQRIS.Switching,
		MerchantCategoryCode: testQRIS.MerchantCategoryCode,
		CurrencyCode:         testQRIS.CurrencyCode,
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
			Content: "666",
			Data:    testPaymentFeeFixedTag + fmt.Sprintf("%02d", len("666")) + "666",
		},
		CountryCode:           testQRIS.CountryCode,
		MerchantName:          testQRIS.MerchantName,
		MerchantCity:          testMerchantCity,
		MerchantPostalCode:    testMerchantPostalCode,
		AdditionalInformation: testQRIS.AdditionalInformation,
		CRCCode: entities.Data{
			Tag:     testCRCCodeTag,
			Content: "AZ15",
			Data:    testCRCCodeTag + fmt.Sprintf("%02d", len("AZ15")) + "AZ15",
		},
	}

	testQRISDynamicModel  = *mapQRISDynamicEntityToModel(&testQRISDynamic)
	testQRISDynamicString = testQRISDynamic.Version.Data +
		testQRISDynamic.Category.Data +
		testQRISDynamic.Acquirer.Data +
		testQRISDynamic.Switching.Data +
		testQRISDynamic.MerchantCategoryCode.Data +
		testQRISDynamic.CurrencyCode.Data +
		testQRISDynamic.PaymentAmount.Data +
		testQRISDynamic.PaymentFeeCategory.Data +
		testQRISDynamic.PaymentFee.Data +
		testQRISDynamic.CountryCode.Data +
		testQRISDynamic.MerchantName.Data +
		testQRISDynamic.MerchantCity.Data +
		testQRISDynamic.MerchantPostalCode.Data +
		testQRISDynamic.AdditionalInformation.Data +
		testQRISDynamic.CRCCode.Data
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
	ParseFunc           func(qrString string) (*entities.QRIS, error, *[]string)
	ToDynamicFunc       func(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic
	DynamicToStringFunc func(qrisDynamic *entities.QRISDynamic) string
	ValidateFunc        func(qris *entities.QRIS) bool
}

func (m *mockQRISUsecase) Parse(qrString string) (*entities.QRIS, error, *[]string) {
	if m.ParseFunc != nil {
		return m.ParseFunc(qrString)
	}
	return nil, nil, nil
}

func (m *mockQRISUsecase) ToDynamic(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic {
	if m.ToDynamicFunc != nil {
		return m.ToDynamicFunc(qris, merchantCity, merchantPostalCode, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)
	}
	return nil
}

func (m *mockQRISUsecase) DynamicToString(qrisDynamic *entities.QRISDynamic) string {
	if m.DynamicToStringFunc != nil {
		return m.DynamicToStringFunc(qrisDynamic)
	}
	return ""
}

func (m *mockQRISUsecase) Validate(qris *entities.QRIS) bool {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(qris)
	}
	return false
}
