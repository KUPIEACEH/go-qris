package usecases

import (
	"errors"
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
				acquirerUsecase:       &Acquirer{},
				switchingUsecase:      &Switching{},
				dataUsecase:           &Data{},
				crc16CCITTUsecase:     &CRC16CCITT{},
				qrisTags:              QRISTags{},
				qrisCategoryContents:  QRISCategoryContents{},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{},
			},
			want: &QRIS{
				acquirerUsecase:       &Acquirer{},
				switchingUsecase:      &Switching{},
				dataUsecase:           &Data{},
				crc16CCITTUsecase:     &CRC16CCITT{},
				qrisTags:              QRISTags{},
				qrisCategoryContents:  QRISCategoryContents{},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{},
			},
		},
		{
			name: "Success: With Field",
			fields: QRIS{
				acquirerUsecase:   &Acquirer{},
				switchingUsecase:  &Switching{},
				dataUsecase:       &Data{},
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
					StaticContent:  testCategoryStaticContent,
					DynamicContent: testCategoryDynamicContent,
				},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{
					CategoryFixedContent:   testPaymentFeeCategoryFixedContent,
					CategoryPercentContent: testPaymentFeeCategoryPercentContent,
					FixedTag:               testPaymentFeeFixedTag,
					PercentTag:             testPaymentFeePercentTag,
				},
			},
			want: &QRIS{
				acquirerUsecase:   &Acquirer{},
				switchingUsecase:  &Switching{},
				dataUsecase:       &Data{},
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
					StaticContent:  testCategoryStaticContent,
					DynamicContent: testCategoryDynamicContent,
				},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{
					CategoryFixedContent:   testPaymentFeeCategoryFixedContent,
					CategoryPercentContent: testPaymentFeeCategoryPercentContent,
					FixedTag:               testPaymentFeeFixedTag,
					PercentTag:             testPaymentFeePercentTag,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRIS(test.fields.acquirerUsecase, test.fields.switchingUsecase, test.fields.dataUsecase, test.fields.crc16CCITTUsecase, test.fields.qrisTags, test.fields.qrisCategoryContents, test.fields.qrisDynamicPaymentFee)

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

func TestQRISExtractStatic(t *testing.T) {
	type args struct {
		qrString string
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		want      *entities.QRISStatic
		wantError error
	}{
		{
			name: "Error: Parse",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want:      nil,
			wantError: errors.New("invalid format code"),
		},
		{
			name: "Error: Not Static QRIS Format",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						switch codeString[:2] {
						case testQRISStatic.Version.Tag:
							return &testQRISStatic.Version, nil
						case testQRISStatic.Category.Tag:
							return &entities.ExtractData{
								Tag:     testQRISStatic.Category.Tag,
								Content: "12",
								Data:    testQRISStatic.Category.Tag + "0212",
							}, nil
						default:
							return nil, nil
						}
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want:      nil,
			wantError: errors.New("not static QRIS content detected"),
		},
		{
			name: "Error: Invalid Extract Acquirer",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						switch codeString[:2] {
						case testQRISStatic.Version.Tag:
							return &testQRISStatic.Version, nil
						case testQRISStatic.Category.Tag:
							return &testQRISStatic.Category, nil
						case testQRISStatic.Acquirer.Tag:
							return &entities.ExtractData{
								Tag:     testQRISStatic.Acquirer.Tag,
								Content: testQRISStatic.Acquirer.Content,
								Data:    testQRISStatic.Acquirer.Data,
							}, nil
						default:
							return nil, nil
						}
					},
				},
				acquirerUsecase: &mockAcquirerUsecase{
					ExtractFunc: func(content string) (*entities.AcquirerDetail, error) {
						return nil, fmt.Errorf("invalid acquirer account for content %s", testQRISStatic.Acquirer.Content)
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid extract acquirer for content %s", testQRISStatic.Acquirer.Content),
		},
		{
			name: "Error: Invalid Extract Switching",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						switch codeString[:2] {
						case testQRISStatic.Version.Tag:
							return &testQRISStatic.Version, nil
						case testQRISStatic.Category.Tag:
							return &testQRISStatic.Category, nil
						case testQRISStatic.Acquirer.Tag:
							return &entities.ExtractData{
								Tag:     testQRISStatic.Acquirer.Tag,
								Content: testQRISStatic.Acquirer.Content,
								Data:    testQRISStatic.Acquirer.Data,
							}, nil
						case testQRISStatic.Switching.Tag:
							return &entities.ExtractData{
								Tag:     testQRISStatic.Switching.Tag,
								Content: testQRISStatic.Switching.Content,
								Data:    testQRISStatic.Switching.Data,
							}, nil
						default:
							return nil, nil
						}
					},
				},
				acquirerUsecase: &mockAcquirerUsecase{
					ExtractFunc: func(content string) (*entities.AcquirerDetail, error) {
						return &testAcquirerDetail, nil
					},
				},
				switchingUsecase: &mockSwitchingUsecase{
					ExtractFunc: func(content string) (*entities.SwitchingDetail, error) {
						return nil, fmt.Errorf("invalid extract switching for content %s", testQRISStatic.Switching.Content)
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid extract switching for content %s", testQRISStatic.Switching.Content),
		},
		{
			name: "Success",
			fields: QRIS{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						switch codeString[:2] {
						case testQRISStatic.Version.Tag:
							return &testQRISStatic.Version, nil
						case testQRISStatic.Category.Tag:
							return &testQRISStatic.Category, nil
						case testQRISStatic.Acquirer.Tag:
							return &entities.ExtractData{
								Tag:     testQRISStatic.Acquirer.Tag,
								Content: testQRISStatic.Acquirer.Content,
								Data:    testQRISStatic.Acquirer.Data,
							}, nil
						case testQRISStatic.Switching.Tag:
							return &entities.ExtractData{
								Tag:     testQRISStatic.Switching.Tag,
								Content: testQRISStatic.Switching.Content,
								Data:    testQRISStatic.Switching.Data,
							}, nil
						case testQRISStatic.MerchantCategoryCode.Tag:
							return &testQRISStatic.MerchantCategoryCode, nil
						case testQRISStatic.CurrencyCode.Tag:
							return &testQRISStatic.CurrencyCode, nil
						case testQRISStatic.CountryCode.Tag:
							return &testQRISStatic.CountryCode, nil
						case testQRISStatic.MerchantName.Tag:
							return &testQRISStatic.MerchantName, nil
						case testQRISStatic.MerchantCity.Tag:
							return &testQRISStatic.MerchantCity, nil
						case testQRISStatic.MerchantPostalCode.Tag:
							return &testQRISStatic.MerchantPostalCode, nil
						case testQRISStatic.AdditionalInformation.Tag:
							return &testQRISStatic.AdditionalInformation, nil
						case testQRISStatic.CRCCode.Tag:
							return &testQRISStatic.CRCCode, nil
						default:
							return nil, nil
						}
					},
				},
				acquirerUsecase: &mockAcquirerUsecase{
					ExtractFunc: func(content string) (*entities.AcquirerDetail, error) {
						return &testQRISStatic.Acquirer.Detail, nil
					},
				},
				switchingUsecase: &mockSwitchingUsecase{
					ExtractFunc: func(content string) (*entities.SwitchingDetail, error) {
						return &testQRISStatic.Switching.Detail, nil
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want:      &testQRISStatic,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				acquirerUsecase:   test.fields.acquirerUsecase,
				switchingUsecase:  test.fields.switchingUsecase,
				dataUsecase:       test.fields.dataUsecase,
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
					StaticContent:  testCategoryStaticContent,
					DynamicContent: testCategoryDynamicContent,
				},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{
					CategoryFixedContent:   testPaymentFeeCategoryFixedContent,
					CategoryPercentContent: testPaymentFeeCategoryPercentContent,
					FixedTag:               testPaymentFeeFixedTag,
					PercentTag:             testPaymentFeePercentTag,
				},
			}

			got, err := uc.ExtractStatic(test.args.qrString)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "ExtractStatic()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ExtractStatic()", test.want, got)
			}
		})
	}
}

func TestQRISStaticToDynamic(t *testing.T) {
	var (
		testMerchantCityContent       = "New Merchant City"
		testMerchantPostalCodeContent = "55181"
		testPaymentAmountValue        = uint32(1337)
		testMerchantCity              = entities.ExtractData{
			Tag:     testQRISStatic.MerchantCity.Tag,
			Content: testMerchantCityContent,
			Data:    testQRISStatic.MerchantCity.Tag + fmt.Sprintf("%02d", len(testMerchantCityContent)) + testMerchantCityContent,
		}
		testMerchantPostalCode = entities.ExtractData{
			Tag:     testQRISStatic.MerchantPostalCode.Tag,
			Content: testMerchantPostalCodeContent,
			Data:    testQRISStatic.MerchantPostalCode.Tag + fmt.Sprintf("%02d", len(testMerchantPostalCodeContent)) + testMerchantPostalCodeContent,
		}
	)

	type args struct {
		qrisStatic         entities.QRISStatic
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
					ModifyContentFunc: func(extractData entities.ExtractData, content string) entities.ExtractData {
						length := len(content)
						switch extractData.Tag {
						case testQRISStatic.MerchantCity.Tag:
							return entities.ExtractData{
								Tag:     testQRISStatic.MerchantCity.Tag,
								Content: content,
								Data:    testQRISStatic.MerchantCity.Tag + fmt.Sprintf("%02d", length) + content,
							}
						case testQRISStatic.MerchantPostalCode.Tag:
							return entities.ExtractData{
								Tag:     testQRISStatic.MerchantPostalCode.Tag,
								Content: content,
								Data:    testQRISStatic.MerchantPostalCode.Tag + fmt.Sprintf("%02d", length) + content,
							}
						default:
							return entities.ExtractData{}
						}
					},
				},
			},
			args: args{
				qrisStatic:         testQRISStatic,
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
			},
			want: &entities.QRISDynamic{
				Version: testQRISStatic.Version,
				Category: entities.ExtractData{
					Tag:     testCategoryTag,
					Content: testCategoryDynamicContent,
					Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
				},
				Acquirer:             testQRISStatic.Acquirer,
				Switching:            testQRISStatic.Switching,
				MerchantCategoryCode: testQRISStatic.MerchantCategoryCode,
				CurrencyCode:         testQRISStatic.CurrencyCode,
				PaymentAmount: entities.ExtractData{
					Tag:     testPaymentAmountTag,
					Content: fmt.Sprintf("%d", testPaymentAmountValue),
					Data:    testPaymentAmountTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", testPaymentAmountValue))) + fmt.Sprintf("%d", testPaymentAmountValue),
				},
				PaymentFeeCategory: entities.ExtractData{
					Tag:     testPaymentFeeCategoryTag,
					Content: testPaymentFeeCategoryFixedContent,
					Data:    testPaymentFeeCategoryTag + fmt.Sprintf("%02d", len(testPaymentFeeCategoryFixedContent)) + testPaymentFeeCategoryFixedContent,
				},
				PaymentFee: entities.ExtractData{
					Tag:     testPaymentFeeFixedTag,
					Content: "666",
					Data:    testPaymentFeeFixedTag + fmt.Sprintf("%02d", len("666")) + "666",
				},
				CountryCode:           testQRISStatic.CountryCode,
				MerchantName:          testQRISStatic.MerchantName,
				MerchantCity:          testMerchantCity,
				MerchantPostalCode:    testMerchantPostalCode,
				AdditionalInformation: testQRISStatic.AdditionalInformation,
				CRCCode: entities.ExtractData{
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
					ModifyContentFunc: func(extractData entities.ExtractData, content string) entities.ExtractData {
						length := len(content)
						switch extractData.Tag {
						case testQRISStatic.MerchantCity.Tag:
							return entities.ExtractData{
								Tag:     testQRISStatic.MerchantCity.Tag,
								Content: content,
								Data:    testQRISStatic.MerchantCity.Tag + fmt.Sprintf("%02d", length) + content,
							}
						case testQRISStatic.MerchantPostalCode.Tag:
							return entities.ExtractData{
								Tag:     testQRISStatic.MerchantPostalCode.Tag,
								Content: content,
								Data:    testQRISStatic.MerchantPostalCode.Tag + fmt.Sprintf("%02d", length) + content,
							}
						default:
							return entities.ExtractData{}
						}
					},
				},
			},
			args: args{
				qrisStatic:         testQRISStatic,
				paymentFeeCategory: "PERCENT",
				paymentFee:         25,
			},
			want: &entities.QRISDynamic{
				Version: testQRISStatic.Version,
				Category: entities.ExtractData{
					Tag:     testCategoryTag,
					Content: testCategoryDynamicContent,
					Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
				},
				Acquirer:             testQRISStatic.Acquirer,
				Switching:            testQRISStatic.Switching,
				MerchantCategoryCode: testQRISStatic.MerchantCategoryCode,
				CurrencyCode:         testQRISStatic.CurrencyCode,
				PaymentAmount: entities.ExtractData{
					Tag:     testPaymentAmountTag,
					Content: fmt.Sprintf("%d", testPaymentAmountValue),
					Data:    testPaymentAmountTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", testPaymentAmountValue))) + fmt.Sprintf("%d", testPaymentAmountValue),
				},
				PaymentFeeCategory: entities.ExtractData{
					Tag:     testPaymentFeeCategoryTag,
					Content: testPaymentFeeCategoryPercentContent,
					Data:    testPaymentFeeCategoryTag + fmt.Sprintf("%02d", len(testPaymentFeeCategoryPercentContent)) + testPaymentFeeCategoryPercentContent,
				},
				PaymentFee: entities.ExtractData{
					Tag:     testPaymentFeePercentTag,
					Content: "25",
					Data:    testPaymentFeePercentTag + fmt.Sprintf("%02d", len("25")) + "25",
				},
				CountryCode:           testQRISStatic.CountryCode,
				MerchantName:          testQRISStatic.MerchantName,
				MerchantCity:          testMerchantCity,
				MerchantPostalCode:    testMerchantPostalCode,
				AdditionalInformation: testQRISStatic.AdditionalInformation,
				CRCCode: entities.ExtractData{
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
					ModifyContentFunc: func(extractData entities.ExtractData, content string) entities.ExtractData {
						length := len(content)
						switch extractData.Tag {
						case testQRISStatic.MerchantCity.Tag:
							return entities.ExtractData{
								Tag:     testQRISStatic.MerchantCity.Tag,
								Content: content,
								Data:    testQRISStatic.MerchantCity.Tag + fmt.Sprintf("%02d", length) + content,
							}
						case testQRISStatic.MerchantPostalCode.Tag:
							return entities.ExtractData{
								Tag:     testQRISStatic.MerchantPostalCode.Tag,
								Content: content,
								Data:    testQRISStatic.MerchantPostalCode.Tag + fmt.Sprintf("%02d", length) + content,
							}
						default:
							return entities.ExtractData{}
						}
					},
				},
			},
			args: args{
				qrisStatic:         testQRISStatic,
				paymentFeeCategory: "UNDEFINED",
				paymentFee:         1337,
			},
			want: &entities.QRISDynamic{
				Version: testQRISStatic.Version,
				Category: entities.ExtractData{
					Tag:     testCategoryTag,
					Content: testCategoryDynamicContent,
					Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
				},
				Acquirer:             testQRISStatic.Acquirer,
				Switching:            testQRISStatic.Switching,
				MerchantCategoryCode: testQRISStatic.MerchantCategoryCode,
				CurrencyCode:         testQRISStatic.CurrencyCode,
				PaymentAmount: entities.ExtractData{
					Tag:     testPaymentAmountTag,
					Content: fmt.Sprintf("%d", testPaymentAmountValue),
					Data:    testPaymentAmountTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", testPaymentAmountValue))) + fmt.Sprintf("%d", testPaymentAmountValue),
				},
				PaymentFeeCategory: entities.ExtractData{
					Tag:     "",
					Content: "",
					Data:    "",
				},
				PaymentFee: entities.ExtractData{
					Tag:     "",
					Content: "",
					Data:    "",
				},
				CountryCode:           testQRISStatic.CountryCode,
				MerchantName:          testQRISStatic.MerchantName,
				MerchantCity:          testMerchantCity,
				MerchantPostalCode:    testMerchantPostalCode,
				AdditionalInformation: testQRISStatic.AdditionalInformation,
				CRCCode: entities.ExtractData{
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
				acquirerUsecase:   test.fields.acquirerUsecase,
				switchingUsecase:  test.fields.switchingUsecase,
				dataUsecase:       test.fields.dataUsecase,
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
					StaticContent:  testCategoryStaticContent,
					DynamicContent: testCategoryDynamicContent,
				},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{
					CategoryFixedContent:   testPaymentFeeCategoryFixedContent,
					CategoryPercentContent: testPaymentFeeCategoryPercentContent,
					FixedTag:               testPaymentFeeFixedTag,
					PercentTag:             testPaymentFeePercentTag,
				},
			}

			got := uc.StaticToDynamic(&test.args.qrisStatic, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, test.args.paymentFeeCategory, test.args.paymentFee)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "StaticToDynamic()", test.want, got)
			}
		})
	}
}

func TestQRISDynamicToDynamicString(t *testing.T) {
	type args struct {
		qrisDynamic *entities.QRISDynamic
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
				qrisDynamic: &entities.QRISDynamic{
					Version: testQRISStatic.Version,
					Category: entities.ExtractData{
						Tag:     testCategoryTag,
						Content: testCategoryDynamicContent,
						Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
					},
					Acquirer:             testQRISStatic.Acquirer,
					Switching:            testQRISStatic.Switching,
					MerchantCategoryCode: testQRISStatic.MerchantCategoryCode,
					CurrencyCode:         testQRISStatic.CurrencyCode,
					PaymentAmount: entities.ExtractData{
						Tag:     testPaymentAmountTag,
						Content: fmt.Sprintf("%d", 1337),
						Data:    testPaymentAmountTag + fmt.Sprintf("%02d", len(fmt.Sprintf("%d", 1337))) + fmt.Sprintf("%d", 1337),
					},
					PaymentFeeCategory: entities.ExtractData{
						Tag:     "",
						Content: "",
						Data:    "",
					},
					PaymentFee: entities.ExtractData{
						Tag:     "",
						Content: "",
						Data:    "",
					},
					CountryCode:           testQRISStatic.CountryCode,
					MerchantName:          testQRISStatic.MerchantName,
					MerchantCity:          testQRISStatic.MerchantCity,
					MerchantPostalCode:    testQRISStatic.MerchantPostalCode,
					AdditionalInformation: testQRISStatic.AdditionalInformation,
					CRCCode: entities.ExtractData{
						Tag:     testCRCCodeTag,
						Content: "AZ15",
						Data:    testCRCCodeTag + fmt.Sprintf("%02d", len("AZ15")) + "AZ15",
					},
				},
			},
			want: "00020101021226620016COM.MEMBASUH.WWW011893600915302259148102090225914810303UMI51440014ID.CO.QRIS.WWW0215ID10200176114730303UMI520448295303360540413375802ID5912Sintas Store6015Kota Yogyakarta6105550006202016304AZ15",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				acquirerUsecase:   test.fields.acquirerUsecase,
				switchingUsecase:  test.fields.switchingUsecase,
				dataUsecase:       test.fields.dataUsecase,
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
					StaticContent:  testCategoryStaticContent,
					DynamicContent: testCategoryDynamicContent,
				},
				qrisDynamicPaymentFee: QRISDynamicPaymentFee{
					CategoryFixedContent:   testPaymentFeeCategoryFixedContent,
					CategoryPercentContent: testPaymentFeeCategoryPercentContent,
					FixedTag:               testPaymentFeeFixedTag,
					PercentTag:             testPaymentFeePercentTag,
				},
			}

			got := uc.DynamicToDynamicString(test.args.qrisDynamic)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "DynamicToDynamicString()", test.want, got)
			}
		})
	}
}
