package usecases

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
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

	testQRISString = testQRIS.Version.Data +
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
)

type mockAcquirerUsecase struct {
	ParseFunc func(content string) (*entities.AcquirerDetail, error)
}

func (m *mockAcquirerUsecase) Parse(content string) (*entities.AcquirerDetail, error) {
	if m.ParseFunc != nil {
		return m.ParseFunc(content)
	}
	return nil, nil
}

type mockSwitchingUsecase struct {
	ParseFunc func(content string) (*entities.SwitchingDetail, error)
}

func (m *mockSwitchingUsecase) Parse(content string) (*entities.SwitchingDetail, error) {
	if m.ParseFunc != nil {
		return m.ParseFunc(content)
	}
	return nil, nil
}

type mockCRC16CCITTUsecase struct {
	GenerateCodeFunc func(code string) string
}

func (m *mockCRC16CCITTUsecase) GenerateCode(code string) string {
	if m.GenerateCodeFunc != nil {
		return m.GenerateCodeFunc(code)
	}
	return ""
}

type mockDataUsecase struct {
	ParseFunc         func(codeString string) (*entities.Data, error)
	ModifyContentFunc func(data *entities.Data, content string) *entities.Data
}

func (m *mockDataUsecase) Parse(codeString string) (*entities.Data, error) {
	if m.ParseFunc != nil {
		return m.ParseFunc(codeString)
	}
	return nil, nil
}

func (m *mockDataUsecase) ModifyContent(data *entities.Data, content string) *entities.Data {
	if m.ModifyContentFunc != nil {
		return m.ModifyContentFunc(data, content)
	}
	return &entities.Data{}
}

type mockFieldUsecase struct {
	AssignFunc  func(qris *entities.QRIS, data *entities.Data) error
	IsValidFunc func(qris *entities.QRIS, errs *[]string)
}

func (m *mockFieldUsecase) Assign(qris *entities.QRIS, data *entities.Data) error {
	if m.AssignFunc != nil {
		return m.AssignFunc(qris, data)
	}
	return nil
}

func (m *mockFieldUsecase) IsValid(qris *entities.QRIS, errs *[]string) {
	if m.IsValidFunc != nil {
		m.IsValidFunc(qris, errs)
		return
	}
	return
}
