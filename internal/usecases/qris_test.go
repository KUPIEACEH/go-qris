package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewQRIS(t *testing.T) {
	tests := []struct {
		name   string
		fields QRIS
		want   QRISInterface
	}{
		{
			name: "Success: No Field",
			fields: QRIS{
				dataUsecase:                    &Data{},
				fieldUsecase:                   &Field{},
				crc16CCITTUsecase:              &CRC16CCITT{},
				qrisTags:                       QRISTags{},
				qrisCategoryContents:           QRISCategoryContents{},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{},
			},
			want: &QRIS{
				dataUsecase:                    &Data{},
				fieldUsecase:                   &Field{},
				crc16CCITTUsecase:              &CRC16CCITT{},
				qrisTags:                       QRISTags{},
				qrisCategoryContents:           QRISCategoryContents{},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{},
			},
		},
		{
			name: "Success: With Field",
			fields: QRIS{
				dataUsecase:       &Data{},
				fieldUsecase:      &Field{},
				crc16CCITTUsecase: &CRC16CCITT{},
				qrisTags: QRISTags{
					VersionTag:               testVersionTag,
					CategoryTag:              testCategoryTag,
					AcquirerTag:              testAcquirerTag,
					SwitchingTag:             testSwitchingTag,
					MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
					CurrencyCodeTag:          testCurrencyCodeTag,
					PaymentAmountTag:         testPaymentAmountTag,
					PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
					CountryCodeTag:           testCountryCodeTag,
					MerchantNameTag:          testMerchantNameTag,
					MerchantCityTag:          testMerchantCityTag,
					MerchantPostalCodeTag:    testMerchantPostalCodeTag,
					AdditionalInformationTag: testAdditionalInformationTag,
					CRCCodeTag:               testCRCCodeTag,
				},
				qrisCategoryContents: QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			},
			want: &QRIS{
				dataUsecase:       &Data{},
				fieldUsecase:      &Field{},
				crc16CCITTUsecase: &CRC16CCITT{},
				qrisTags: QRISTags{
					VersionTag:               testVersionTag,
					CategoryTag:              testCategoryTag,
					AcquirerTag:              testAcquirerTag,
					SwitchingTag:             testSwitchingTag,
					MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
					CurrencyCodeTag:          testCurrencyCodeTag,
					PaymentAmountTag:         testPaymentAmountTag,
					PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
					CountryCodeTag:           testCountryCodeTag,
					MerchantNameTag:          testMerchantNameTag,
					MerchantCityTag:          testMerchantCityTag,
					MerchantPostalCodeTag:    testMerchantPostalCodeTag,
					AdditionalInformationTag: testAdditionalInformationTag,
					CRCCodeTag:               testCRCCodeTag,
				},
				qrisCategoryContents: QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRIS(test.fields.dataUsecase, test.fields.fieldUsecase, test.fields.crc16CCITTUsecase, test.fields.qrisTags, test.fields.qrisCategoryContents, test.fields.qrisPaymentFeeCategoryContents)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewQRIS", "QRISInterface")
			}

			got, ok := uc.(*QRIS)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*QRIS")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*QRIS", test.want, got)
			}
		})
	}
}

func TestQRISParse(t *testing.T) {
	type args struct {
		qrString string
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		want      *entities.QRIS
		wantError error
	}{
		{
			name: "Error: Parse",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid format code"),
		},
		{
			name: "Error: uc.fieldUsecase.Assign()",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return &testQRIS.Version, nil
					},
				},
				fieldUsecase: &mockFieldUsecase{
					AssignFunc: func(qris *entities.QRIS, data *entities.Data) error {
						return fmt.Errorf("invalid extract acquirer for content %s", testQRIS.Acquirer.Content)
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid extract acquirer for content %s", testQRIS.Acquirer.Content),
		},
		{
			name: "Error: uc.fieldUsecase.Validate()",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return &testQRIS.Version, nil
					},
				},
				fieldUsecase: &mockFieldUsecase{
					AssignFunc: func(qris *entities.QRIS, data *entities.Data) error {
						if data.Tag == testVersionTag {
							qris.Version = *data
						}
						return nil
					},
					ValidateFunc: func(qris *entities.QRIS, errs *[]string) {
						*errs = append(*errs, "Category tag is missing")
						return
					},
				},
			},
			args: args{
				qrString: testQRIS.Version.Data,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid QRIS format"),
		},
		{
			name: "Success",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return &testQRIS.Version, nil
					},
				},
				fieldUsecase: &mockFieldUsecase{
					AssignFunc: func(qris *entities.QRIS, data *entities.Data) error {
						if data.Tag == testVersionTag {
							qris.Version = *data
						}
						return nil
					},
				},
			},
			args: args{
				qrString: testQRIS.Version.Data,
			},
			want: &entities.QRIS{
				Version: testQRIS.Version,
			},
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				dataUsecase:       test.fields.dataUsecase,
				fieldUsecase:      test.fields.fieldUsecase,
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisTags: QRISTags{
					VersionTag:               testVersionTag,
					CategoryTag:              testCategoryTag,
					AcquirerTag:              testAcquirerTag,
					AcquirerBankTransferTag:  testAcquirerBankTransferTag,
					SwitchingTag:             testSwitchingTag,
					MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
					CurrencyCodeTag:          testCurrencyCodeTag,
					PaymentAmountTag:         testPaymentAmountTag,
					PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
					CountryCodeTag:           testCountryCodeTag,
					MerchantNameTag:          testMerchantNameTag,
					MerchantCityTag:          testMerchantCityTag,
					MerchantPostalCodeTag:    testMerchantPostalCodeTag,
					AdditionalInformationTag: testAdditionalInformationTag,
					CRCCodeTag:               testCRCCodeTag,
				},
				qrisCategoryContents: QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got, err, _ := uc.Parse(test.args.qrString)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Parse()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Parse()", test.want, got)
			}
		})
	}
}

func TestQRISToDynamic(t *testing.T) {
	var (
		testMerchantCityContent       = "New Merchant City"
		testMerchantPostalCodeContent = "55181"
		testPaymentAmountValue        = uint32(1337)
		testMerchantCity              = entities.Data{
			Tag:     testQRIS.MerchantCity.Tag,
			Content: testMerchantCityContent,
			Data:    testQRIS.MerchantCity.Tag + fmt.Sprintf("%02d", len(testMerchantCityContent)) + testMerchantCityContent,
		}
		testMerchantPostalCode = entities.Data{
			Tag:     testQRIS.MerchantPostalCode.Tag,
			Content: testMerchantPostalCodeContent,
			Data:    testQRIS.MerchantPostalCode.Tag + fmt.Sprintf("%02d", len(testMerchantPostalCodeContent)) + testMerchantPostalCodeContent,
		}
	)

	type args struct {
		qris               entities.QRIS
		paymentFeeCategory string
		paymentFee         uint32
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   *entities.QRISDynamic
	}{
		{
			name: "Success: Fixed Payment Fee",
			fields: QRIS{
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "AZ15"
					},
				},
				dataUsecase: &mockDataUsecase{
					ModifyContentFunc: func(extractData *entities.Data, content string) *entities.Data {
						length := len(content)
						switch extractData.Tag {
						case testQRIS.MerchantCity.Tag:
							return &entities.Data{
								Tag:     testQRIS.MerchantCity.Tag,
								Content: content,
								Data:    testQRIS.MerchantCity.Tag + fmt.Sprintf("%02d", length) + content,
							}
						case testQRIS.MerchantPostalCode.Tag:
							return &entities.Data{
								Tag:     testQRIS.MerchantPostalCode.Tag,
								Content: content,
								Data:    testQRIS.MerchantPostalCode.Tag + fmt.Sprintf("%02d", length) + content,
							}
						default:
							return &entities.Data{}
						}
					},
				},
			},
			args: args{
				qris:               testQRIS,
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
			},
			want: &entities.QRISDynamic{
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
			},
		},
		{
			name: "Success: Percent Payment Fee",
			fields: QRIS{
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "AZ15"
					},
				},
				dataUsecase: &mockDataUsecase{
					ModifyContentFunc: func(extractData *entities.Data, content string) *entities.Data {
						length := len(content)
						switch extractData.Tag {
						case testQRIS.MerchantCity.Tag:
							return &entities.Data{
								Tag:     testQRIS.MerchantCity.Tag,
								Content: content,
								Data:    testQRIS.MerchantCity.Tag + fmt.Sprintf("%02d", length) + content,
							}
						case testQRIS.MerchantPostalCode.Tag:
							return &entities.Data{
								Tag:     testQRIS.MerchantPostalCode.Tag,
								Content: content,
								Data:    testQRIS.MerchantPostalCode.Tag + fmt.Sprintf("%02d", length) + content,
							}
						default:
							return &entities.Data{}
						}
					},
				},
			},
			args: args{
				qris:               testQRIS,
				paymentFeeCategory: "PERCENT",
				paymentFee:         25,
			},
			want: &entities.QRISDynamic{
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
					Content: testPaymentFeeCategoryPercentContent,
					Data:    testPaymentFeeCategoryTag + fmt.Sprintf("%02d", len(testPaymentFeeCategoryPercentContent)) + testPaymentFeeCategoryPercentContent,
				},
				PaymentFee: entities.Data{
					Tag:     testPaymentFeePercentTag,
					Content: "25",
					Data:    testPaymentFeePercentTag + fmt.Sprintf("%02d", len("25")) + "25",
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
			},
		},
		{
			name: "Success: No Payment Fee",
			fields: QRIS{
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "AZ15"
					},
				},
				dataUsecase: &mockDataUsecase{
					ModifyContentFunc: func(extractData *entities.Data, content string) *entities.Data {
						length := len(content)
						switch extractData.Tag {
						case testQRIS.MerchantCity.Tag:
							return &entities.Data{
								Tag:     testQRIS.MerchantCity.Tag,
								Content: content,
								Data:    testQRIS.MerchantCity.Tag + fmt.Sprintf("%02d", length) + content,
							}
						case testQRIS.MerchantPostalCode.Tag:
							return &entities.Data{
								Tag:     testQRIS.MerchantPostalCode.Tag,
								Content: content,
								Data:    testQRIS.MerchantPostalCode.Tag + fmt.Sprintf("%02d", length) + content,
							}
						default:
							return &entities.Data{}
						}
					},
				},
			},
			args: args{
				qris:               testQRIS,
				paymentFeeCategory: "UNDEFINED",
				paymentFee:         1337,
			},
			want: &entities.QRISDynamic{
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
					Tag:     "",
					Content: "",
					Data:    "",
				},
				PaymentFee: entities.Data{
					Tag:     "",
					Content: "",
					Data:    "",
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
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				dataUsecase:       test.fields.dataUsecase,
				fieldUsecase:      test.fields.fieldUsecase,
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisTags: QRISTags{
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
				},
				qrisCategoryContents: QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.ToDynamic(&test.args.qris, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, test.args.paymentFeeCategory, test.args.paymentFee)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToDynamic()", test.want, got)
			}
		})
	}
}

func TestQRISDynamicToString(t *testing.T) {
	type args struct {
		qrisDynamic *entities.QRISDynamic
	}

	testQRISDynamic := &entities.QRISDynamic{
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
			Content: fmt.Sprintf("%d", 1337),
			Data:    testPaymentAmountTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", 1337))) + fmt.Sprintf("%d", 1337),
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
		CountryCode:           testQRIS.CountryCode,
		MerchantName:          testQRIS.MerchantName,
		MerchantCity:          testQRIS.MerchantCity,
		MerchantPostalCode:    testQRIS.MerchantPostalCode,
		AdditionalInformation: testQRIS.AdditionalInformation,
		CRCCode: entities.Data{
			Tag:     testCRCCodeTag,
			Content: "AZ15",
			Data:    testCRCCodeTag + fmt.Sprintf("%02d", len("AZ15")) + "AZ15",
		},
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   string
	}{
		{
			name:   "Success",
			fields: QRIS{},
			args: args{
				qrisDynamic: testQRISDynamic,
			},
			want: testQRISDynamic.Version.Data +
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
				testQRISDynamic.CRCCode.Data,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				dataUsecase:       test.fields.dataUsecase,
				fieldUsecase:      test.fields.fieldUsecase,
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisTags: QRISTags{
					VersionTag:               testVersionTag,
					CategoryTag:              testCategoryTag,
					AcquirerTag:              testAcquirerTag,
					SwitchingTag:             testSwitchingTag,
					MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
					CurrencyCodeTag:          testCurrencyCodeTag,
					PaymentAmountTag:         testPaymentAmountTag,
					PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
					CountryCodeTag:           testCountryCodeTag,
					MerchantNameTag:          testMerchantNameTag,
					MerchantCityTag:          testMerchantCityTag,
					MerchantPostalCodeTag:    testMerchantPostalCodeTag,
					AdditionalInformationTag: testAdditionalInformationTag,
					CRCCodeTag:               testCRCCodeTag,
				},
				qrisCategoryContents: QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.DynamicToString(test.args.qrisDynamic)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "DynamicToString()", test.want, got)
			}
		})
	}
}

func TestQRISValidate(t *testing.T) {
	type args struct {
		qris *entities.QRIS
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   bool
	}{
		{
			name: "Success: False",
			fields: QRIS{
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "AZ15"
					},
				},
			},
			args: args{
				qris: &testQRIS,
			},
			want: false,
		},
		{
			name: "Success: False",
			fields: QRIS{
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "1FA2"
					},
				},
			},
			args: args{
				qris: &testQRIS,
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				dataUsecase:       test.fields.dataUsecase,
				fieldUsecase:      test.fields.fieldUsecase,
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisTags: QRISTags{
					VersionTag:               testVersionTag,
					CategoryTag:              testCategoryTag,
					AcquirerTag:              testAcquirerTag,
					SwitchingTag:             testSwitchingTag,
					MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
					CurrencyCodeTag:          testCurrencyCodeTag,
					PaymentAmountTag:         testPaymentAmountTag,
					PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
					CountryCodeTag:           testCountryCodeTag,
					MerchantNameTag:          testMerchantNameTag,
					MerchantCityTag:          testMerchantCityTag,
					MerchantPostalCodeTag:    testMerchantPostalCodeTag,
					AdditionalInformationTag: testAdditionalInformationTag,
					CRCCodeTag:               testCRCCodeTag,
				},
				qrisCategoryContents: QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.Validate(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Validate()", test.want, got)
			}
		})
	}
}
