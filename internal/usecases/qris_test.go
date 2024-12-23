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
				qrisUsecases: &QRISUsecases{
					Data:                  &Data{},
					Field:                 &Field{},
					PaymentFee:            &PaymentFee{},
					AdditionalInformation: &AdditionalInformation{},
					CRC16CCITT:            &CRC16CCITT{},
				},
				qrisTags:                       &QRISTags{},
				qrisCategoryContents:           &QRISCategoryContents{},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{},
			},
			want: &QRIS{
				qrisUsecases: &QRISUsecases{
					Data:                  &Data{},
					Field:                 &Field{},
					PaymentFee:            &PaymentFee{},
					AdditionalInformation: &AdditionalInformation{},
					CRC16CCITT:            &CRC16CCITT{},
				},
				qrisTags:                       &QRISTags{},
				qrisCategoryContents:           &QRISCategoryContents{},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{},
			},
		},
		{
			name: "Success: With Field",
			fields: QRIS{
				qrisUsecases: &QRISUsecases{
					Data:                  &Data{},
					Field:                 &Field{},
					PaymentFee:            &PaymentFee{},
					AdditionalInformation: &AdditionalInformation{},
					CRC16CCITT:            &CRC16CCITT{},
				},
				qrisTags: &QRISTags{
					Version:               testVersionTag,
					Category:              testCategoryTag,
					Acquirer:              testAcquirerTag,
					Switching:             testSwitchingTag,
					MerchantCategoryCode:  testMerchantCategoryCodeTag,
					CurrencyCode:          testCurrencyCodeTag,
					PaymentAmount:         testPaymentAmountTag,
					PaymentFeeCategory:    testPaymentFeeCategoryTag,
					CountryCode:           testCountryCodeTag,
					MerchantName:          testMerchantNameTag,
					MerchantCity:          testMerchantCityTag,
					MerchantPostalCode:    testMerchantPostalCodeTag,
					AdditionalInformation: testAdditionalInformationTag,
					CRCCode:               testCRCCodeTag,
				},
				qrisCategoryContents: &QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			},
			want: &QRIS{
				qrisUsecases: &QRISUsecases{
					Data:                  &Data{},
					Field:                 &Field{},
					PaymentFee:            &PaymentFee{},
					AdditionalInformation: &AdditionalInformation{},
					CRC16CCITT:            &CRC16CCITT{},
				},
				qrisTags: &QRISTags{
					Version:               testVersionTag,
					Category:              testCategoryTag,
					Acquirer:              testAcquirerTag,
					Switching:             testSwitchingTag,
					MerchantCategoryCode:  testMerchantCategoryCodeTag,
					CurrencyCode:          testCurrencyCodeTag,
					PaymentAmount:         testPaymentAmountTag,
					PaymentFeeCategory:    testPaymentFeeCategoryTag,
					CountryCode:           testCountryCodeTag,
					MerchantName:          testMerchantNameTag,
					MerchantCity:          testMerchantCityTag,
					MerchantPostalCode:    testMerchantPostalCodeTag,
					AdditionalInformation: testAdditionalInformationTag,
					CRCCode:               testCRCCodeTag,
				},
				qrisCategoryContents: &QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRIS(test.fields.qrisUsecases, test.fields.qrisTags, test.fields.qrisCategoryContents, test.fields.qrisPaymentFeeCategoryContents)

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
				qrisUsecases: &QRISUsecases{
					Data: &mockDataUsecase{
						ParseFunc: func(codeString string) (*entities.Data, error) {
							return nil, fmt.Errorf("invalid format code")
						},
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
				qrisUsecases: &QRISUsecases{
					Data: &mockDataUsecase{
						ParseFunc: func(codeString string) (*entities.Data, error) {
							return &testQRIS.Version, nil
						},
					},
					Field: &mockFieldUsecase{
						AssignFunc: func(qris *entities.QRIS, data *entities.Data) error {
							return fmt.Errorf("invalid extract acquirer for content %s", testQRIS.Acquirer.Content)
						},
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
			name: "Error: uc.fieldUsecase.IsValid()",
			fields: QRIS{
				qrisUsecases: &QRISUsecases{
					Data: &mockDataUsecase{
						ParseFunc: func(codeString string) (*entities.Data, error) {
							return &testQRIS.Version, nil
						},
					},
					Field: &mockFieldUsecase{
						AssignFunc: func(qris *entities.QRIS, data *entities.Data) error {
							if data.Tag == testVersionTag {
								qris.Version = *data
							}
							return nil
						},
						IsValidFunc: func(qris *entities.QRIS, errs *[]string) {
							*errs = append(*errs, "Category tag is missing")
							return
						},
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
				qrisUsecases: &QRISUsecases{
					Data: &mockDataUsecase{
						ParseFunc: func(codeString string) (*entities.Data, error) {
							return &testQRIS.Version, nil
						},
					},
					Field: &mockFieldUsecase{
						AssignFunc: func(qris *entities.QRIS, data *entities.Data) error {
							if data.Tag == testVersionTag {
								qris.Version = *data
							}
							return nil
						},
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
				qrisUsecases: &QRISUsecases{
					Data:                  test.fields.qrisUsecases.Data,
					Field:                 test.fields.qrisUsecases.Field,
					PaymentFee:            test.fields.qrisUsecases.PaymentFee,
					AdditionalInformation: test.fields.qrisUsecases.AdditionalInformation,
					CRC16CCITT:            test.fields.qrisUsecases.CRC16CCITT,
				},
				qrisTags: &QRISTags{
					Version:               testVersionTag,
					Category:              testCategoryTag,
					Acquirer:              testAcquirerTag,
					AcquirerBankTransfer:  testAcquirerBankTransferTag,
					Switching:             testSwitchingTag,
					MerchantCategoryCode:  testMerchantCategoryCodeTag,
					CurrencyCode:          testCurrencyCodeTag,
					PaymentAmount:         testPaymentAmountTag,
					PaymentFeeCategory:    testPaymentFeeCategoryTag,
					CountryCode:           testCountryCodeTag,
					MerchantName:          testMerchantNameTag,
					MerchantCity:          testMerchantCityTag,
					MerchantPostalCode:    testMerchantPostalCodeTag,
					AdditionalInformation: testAdditionalInformationTag,
					CRCCode:               testCRCCodeTag,
				},
				qrisCategoryContents: &QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
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

func TestQRISModify(t *testing.T) {
	var (
		testMerchantCityContent       = "New Merchant City"
		testMerchantPostalCodeContent = "55181"
		testPaymentAmountValue        = uint32(1337)
		testTerminalLabel             = "Awesome Terminal Label"
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
		terminalLabel      string
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   *entities.QRIS
	}{
		{
			name: "Success",
			fields: QRIS{
				qrisUsecases: &QRISUsecases{
					PaymentFee: &mockPaymentFeeUsecase{
						ModifyFunc: func(qris *entities.QRIS, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
							qris.PaymentFeeCategory = entities.Data{
								Tag:     testPaymentFeeCategoryTag,
								Content: testPaymentFeeCategoryFixedContent,
								Data:    testPaymentFeeCategoryTag + fmt.Sprintf("%02d", len(testPaymentFeeCategoryFixedContent)) + testPaymentFeeCategoryFixedContent,
							}
							qris.PaymentFee = entities.Data{
								Tag:     testPaymentFeeFixedTag,
								Content: "666",
								Data:    testPaymentFeeFixedTag + fmt.Sprintf("%02d", len("666")) + "666",
							}

							return qris
						},
					},
					AdditionalInformation: &mockAdditionalInformationUsecase{
						ToStringFunc: func(additionalInformation *entities.AdditionalInformationDetail) string {
							return testAdditionalInformationDetailTerminalLabelTag + fmt.Sprintf("%02d", len(testTerminalLabel)) + testTerminalLabel
						},
					},
					CRC16CCITT: &mockCRC16CCITTUsecase{
						GenerateCodeFunc: func(code string) string {
							return "AZ15"
						},
					},
					Data: &mockDataUsecase{
						ModifyContentFunc: func(extractData *entities.Data, content string) *entities.Data {
							var tag string
							switch extractData.Tag {
							case testQRIS.PaymentAmount.Tag:
								tag = testQRIS.PaymentAmount.Tag
							case testQRIS.MerchantCity.Tag:
								tag = testQRIS.MerchantCity.Tag
							case testQRIS.MerchantPostalCode.Tag:
								tag = testQRIS.MerchantPostalCode.Tag
							case testQRIS.AdditionalInformation.Detail.TerminalLabel.Tag:
								tag = testQRIS.AdditionalInformation.Detail.TerminalLabel.Tag
							case testQRIS.CRCCode.Tag:
								tag = testQRIS.CRCCode.Tag
							default:
								return &entities.Data{}
							}

							return &entities.Data{
								Tag:     tag,
								Content: content,
								Data:    tag + fmt.Sprintf("%02d", len(content)) + content,
							}
						},
					},
				},
			},
			args: args{
				qris:               testQRIS,
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
				terminalLabel:      testTerminalLabel,
			},
			want: &entities.QRIS{
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
				CountryCode:        testQRIS.CountryCode,
				MerchantName:       testQRIS.MerchantName,
				MerchantCity:       testMerchantCity,
				MerchantPostalCode: testMerchantPostalCode,
				AdditionalInformation: entities.AdditionalInformation{
					Tag:     testAdditionalInformationTag,
					Content: testAdditionalInformationDetailTerminalLabelTag + fmt.Sprintf("%02d", len(testTerminalLabel)) + testTerminalLabel,
					Data:    testAdditionalInformationTag + "26" + testAdditionalInformationDetailTerminalLabelTag + fmt.Sprintf("%02d", len(testTerminalLabel)) + testTerminalLabel,
					Detail: entities.AdditionalInformationDetail{
						TerminalLabel: entities.Data{
							Tag:     testAdditionalInformationDetailTerminalLabelTag,
							Content: testTerminalLabel,
							Data:    testAdditionalInformationDetailTerminalLabelTag + fmt.Sprintf("%02d", len(testTerminalLabel)) + testTerminalLabel,
						},
					},
				},
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
				qrisUsecases: &QRISUsecases{
					Data:                  test.fields.qrisUsecases.Data,
					Field:                 test.fields.qrisUsecases.Field,
					PaymentFee:            test.fields.qrisUsecases.PaymentFee,
					AdditionalInformation: test.fields.qrisUsecases.AdditionalInformation,
					CRC16CCITT:            test.fields.qrisUsecases.CRC16CCITT,
				},
				qrisTags: &QRISTags{
					Version:               testVersionTag,
					Category:              testCategoryTag,
					Acquirer:              testAcquirerTag,
					AcquirerBankTransfer:  testAcquirerBankTransferTag,
					Switching:             testSwitchingTag,
					MerchantCategoryCode:  testMerchantCategoryCodeTag,
					CurrencyCode:          testCurrencyCodeTag,
					PaymentAmount:         testPaymentAmountTag,
					PaymentFeeCategory:    testPaymentFeeCategoryTag,
					PaymentFeeFixed:       testPaymentFeeFixedTag,
					PaymentFeePercent:     testPaymentFeePercentTag,
					CountryCode:           testCountryCodeTag,
					MerchantName:          testMerchantNameTag,
					MerchantCity:          testMerchantCityTag,
					MerchantPostalCode:    testMerchantPostalCodeTag,
					AdditionalInformation: testAdditionalInformationTag,
					CRCCode:               testCRCCodeTag,
				},
				qrisCategoryContents: &QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.Modify(&test.args.qris, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, test.args.paymentFeeCategory, test.args.paymentFee, test.args.terminalLabel)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Modify()", test.want, got)
			}
		})
	}
}

func TestQRISToString(t *testing.T) {
	type args struct {
		qris *entities.QRIS
	}

	testQRIS := &entities.QRIS{
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
			name: "Success",
			fields: QRIS{
				qrisUsecases: &QRISUsecases{
					Data:                  &Data{},
					Field:                 &Field{},
					PaymentFee:            &PaymentFee{},
					AdditionalInformation: &AdditionalInformation{},
					CRC16CCITT:            &CRC16CCITT{},
				},
			},
			args: args{
				qris: testQRIS,
			},
			want: testQRIS.Version.Data +
				testQRIS.Category.Data +
				testQRIS.Acquirer.Data +
				testQRIS.Switching.Data +
				testQRIS.MerchantCategoryCode.Data +
				testQRIS.CurrencyCode.Data +
				testQRIS.PaymentAmount.Data +
				testQRIS.PaymentFeeCategory.Data +
				testQRIS.PaymentFee.Data +
				testQRIS.CountryCode.Data +
				testQRIS.MerchantName.Data +
				testQRIS.MerchantCity.Data +
				testQRIS.MerchantPostalCode.Data +
				testQRIS.AdditionalInformation.Data +
				testQRIS.CRCCode.Data,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				qrisUsecases: &QRISUsecases{
					Data:                  test.fields.qrisUsecases.Data,
					Field:                 test.fields.qrisUsecases.Field,
					PaymentFee:            test.fields.qrisUsecases.PaymentFee,
					AdditionalInformation: test.fields.qrisUsecases.AdditionalInformation,
					CRC16CCITT:            test.fields.qrisUsecases.CRC16CCITT,
				},
				qrisTags: &QRISTags{
					Version:               testVersionTag,
					Category:              testCategoryTag,
					Acquirer:              testAcquirerTag,
					Switching:             testSwitchingTag,
					MerchantCategoryCode:  testMerchantCategoryCodeTag,
					CurrencyCode:          testCurrencyCodeTag,
					PaymentAmount:         testPaymentAmountTag,
					PaymentFeeCategory:    testPaymentFeeCategoryTag,
					CountryCode:           testCountryCodeTag,
					MerchantName:          testMerchantNameTag,
					MerchantCity:          testMerchantCityTag,
					MerchantPostalCode:    testMerchantPostalCodeTag,
					AdditionalInformation: testAdditionalInformationTag,
					CRCCode:               testCRCCodeTag,
				},
				qrisCategoryContents: &QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.ToString(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToString()", test.want, got)
			}
		})
	}
}

func TestQRISIsValid(t *testing.T) {
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
			name: "Success: True",
			fields: QRIS{
				qrisUsecases: &QRISUsecases{
					CRC16CCITT: &mockCRC16CCITTUsecase{
						GenerateCodeFunc: func(code string) string {
							return "1FA2"
						},
					},
				},
			},
			args: args{
				qris: &testQRIS,
			},
			want: true,
		},
		{
			name: "Success: False",
			fields: QRIS{
				qrisUsecases: &QRISUsecases{
					CRC16CCITT: &mockCRC16CCITTUsecase{
						GenerateCodeFunc: func(code string) string {
							return "AZ15"
						},
					},
				},
			},
			args: args{
				qris: &testQRIS,
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				qrisUsecases: &QRISUsecases{
					Data:                  test.fields.qrisUsecases.Data,
					Field:                 test.fields.qrisUsecases.Field,
					PaymentFee:            test.fields.qrisUsecases.PaymentFee,
					AdditionalInformation: test.fields.qrisUsecases.AdditionalInformation,
					CRC16CCITT:            test.fields.qrisUsecases.CRC16CCITT,
				},
				qrisTags: &QRISTags{
					Version:               testVersionTag,
					Category:              testCategoryTag,
					Acquirer:              testAcquirerTag,
					Switching:             testSwitchingTag,
					MerchantCategoryCode:  testMerchantCategoryCodeTag,
					CurrencyCode:          testCurrencyCodeTag,
					PaymentAmount:         testPaymentAmountTag,
					PaymentFeeCategory:    testPaymentFeeCategoryTag,
					CountryCode:           testCountryCodeTag,
					MerchantName:          testMerchantNameTag,
					MerchantCity:          testMerchantCityTag,
					MerchantPostalCode:    testMerchantPostalCodeTag,
					AdditionalInformation: testAdditionalInformationTag,
					CRCCode:               testCRCCodeTag,
				},
				qrisCategoryContents: &QRISCategoryContents{
					Static:  testCategoryStaticContent,
					Dynamic: testCategoryDynamicContent,
				},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.IsValid(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "IsValid()", test.want, got)
			}
		})
	}
}
