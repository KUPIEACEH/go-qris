package controllers

import "github.com/fyvri/go-qris/internal/domain/entities"

var (
	expectedButGotMessage             = "Expected %v = %v, but got = %v"
	expectedErrorButGotMessage        = "Expected %v error = %v, but got = %v"
	expectedTypeAssertionErrorMessage = "Expected type assertion error, but got = %v"
	expectedReturnNonNil              = "Expected %v to return a non-nil %v"

	testVersionTag                  = "00"
	testCategoryTag                 = "01"
	testAcquirerTag                 = "26"
	testSwitchingTag                = "51"
	testMerchantCategoryCodeTag     = "52"
	testCurrencyCodeTag             = "53"
	testCountryCodeTag              = "58"
	testMerchantNameTag             = "59"
	testMerchantCityTag             = "60"
	testMerchantPostalCodeTag       = "61"
	testAdditionalInformationTag    = "62"
	testCRCCodeTag                  = "63"
	testAcquirerDetailSiteTag       = "00"
	testAcquirerDetailMPANTag       = "01"
	testAcquirerDetailTerminalIDTag = "02"
	testAcquirerDetailCategoryTag   = "03"
	testSwitchingDetailSiteTag      = "00"
	testSwitchingDetailNMIDTag      = "02"
	testSwitchingDetailCategoryTag  = "03"

	testAcquirerDetail = entities.AcquirerDetail{
		Site: entities.ExtractData{
			Tag:     testAcquirerDetailSiteTag,
			Content: "COM.MEMBASUH.WWW",
			Data:    testAcquirerDetailSiteTag + "16COM.MEMBASUH.WWW",
		},
		MPAN: entities.ExtractData{
			Tag:     testAcquirerDetailMPANTag,
			Content: "936009153022591481",
			Data:    testAcquirerDetailMPANTag + "18936009153022591481",
		},
		TerminalID: entities.ExtractData{
			Tag:     testAcquirerDetailTerminalIDTag,
			Content: "022591481",
			Data:    testAcquirerDetailTerminalIDTag + "09022591481",
		},
		Category: entities.ExtractData{
			Tag:     testAcquirerDetailCategoryTag,
			Content: "UMI",
			Data:    testAcquirerDetailCategoryTag + "03UMI",
		},
	}

	testSwitchingDetail = entities.SwitchingDetail{
		Site: entities.ExtractData{
			Tag:     testSwitchingDetailSiteTag,
			Content: "ID.CO.QRIS.WWW",
			Data:    testSwitchingDetailSiteTag + "14ID.CO.QRIS.WWW",
		},
		NMID: entities.ExtractData{
			Tag:     testSwitchingDetailNMIDTag,
			Content: "ID1020017611473",
			Data:    testSwitchingDetailNMIDTag + "15ID1020017611473",
		},
		Category: entities.ExtractData{
			Tag:     testSwitchingDetailCategoryTag,
			Content: "UMI",
			Data:    testSwitchingDetailCategoryTag + "03UMI",
		},
	}

	testQRISStatic = entities.QRISStatic{
		Version: entities.ExtractData{
			Tag:     testVersionTag,
			Content: "01",
			Data:    testVersionTag + "0201",
		},
		Category: entities.ExtractData{
			Tag:     testCategoryTag,
			Content: "11",
			Data:    testCategoryTag + "0211",
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
		MerchantCategoryCode: entities.ExtractData{
			Tag:     testMerchantCategoryCodeTag,
			Content: "4829",
			Data:    testMerchantCategoryCodeTag + "044829",
		},
		CurrencyCode: entities.ExtractData{
			Tag:     testCurrencyCodeTag,
			Content: "360",
			Data:    testCurrencyCodeTag + "03360",
		},
		CountryCode: entities.ExtractData{
			Tag:     testCountryCodeTag,
			Content: "ID",
			Data:    testCountryCodeTag + "02ID",
		},
		MerchantName: entities.ExtractData{
			Tag:     testMerchantNameTag,
			Content: "Sintas Store",
			Data:    testMerchantNameTag + "12Sintas Store",
		},
		MerchantCity: entities.ExtractData{
			Tag:     testMerchantCityTag,
			Content: "Kota Yogyakarta",
			Data:    testMerchantCityTag + "15Kota Yogyakarta",
		},
		MerchantPostalCode: entities.ExtractData{
			Tag:     testMerchantPostalCodeTag,
			Content: "55000",
			Data:    testMerchantPostalCodeTag + "0555000",
		},
		AdditionalInformation: entities.ExtractData{
			Tag:     testAdditionalInformationTag,
			Content: "01",
			Data:    testAdditionalInformationTag + "0201",
		},
		CRCCode: entities.ExtractData{
			Tag:     testCRCCodeTag,
			Content: "1FA2",
			Data:    testCRCCodeTag + "041FA2",
		},
	}

	testQRISStaticString = testQRISStatic.Version.Data +
		testQRISStatic.Category.Data +
		testQRISStatic.Acquirer.Data +
		testQRISStatic.Switching.Data +
		testQRISStatic.MerchantCategoryCode.Data +
		testQRISStatic.CurrencyCode.Data +
		testQRISStatic.CountryCode.Data +
		testQRISStatic.MerchantName.Data +
		testQRISStatic.MerchantCity.Data +
		testQRISStatic.MerchantPostalCode.Data +
		testQRISStatic.AdditionalInformation.Data +
		testQRISStatic.CRCCode.Data

	testQRCodeSize = 125
)

type mockQRISUsecase struct {
	ExtractStaticFunc          func(qrString string) (*entities.QRISStatic, error)
	StaticToDynamicFunc        func(qrisStatic *entities.QRISStatic, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) *entities.QRISDynamic
	DynamicToDynamicStringFunc func(QRISDynamic *entities.QRISDynamic) string
}

func (m *mockQRISUsecase) ExtractStatic(qrString string) (*entities.QRISStatic, error) {
	if m.ExtractStaticFunc != nil {
		return m.ExtractStaticFunc(qrString)
	}
	return nil, nil
}

func (m *mockQRISUsecase) StaticToDynamic(qrisStatic *entities.QRISStatic, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) *entities.QRISDynamic {
	if m.StaticToDynamicFunc != nil {
		return m.StaticToDynamicFunc(qrisStatic, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee)
	}
	return nil
}

func (m *mockQRISUsecase) DynamicToDynamicString(QRISDynamic *entities.QRISDynamic) string {
	if m.DynamicToDynamicStringFunc != nil {
		return m.DynamicToDynamicStringFunc(QRISDynamic)
	}
	return ""
}

type mockQRCodeUtil struct {
	StringToImageBase64Func func(qrString string, qrCodeSize int) (string, error)
}

func (m *mockQRCodeUtil) StringToImageBase64(qrString string, qrCodeSize int) (string, error) {
	if m.StringToImageBase64Func != nil {
		return m.StringToImageBase64Func(qrString, qrCodeSize)
	}
	return "", nil
}
