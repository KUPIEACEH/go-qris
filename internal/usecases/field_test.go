package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewField(t *testing.T) {
	tests := []struct {
		name   string
		fields Field
		want   FieldInterface
	}{
		{
			name: "Success: No Field",
			fields: Field{
				acquirerUsecase:              &Acquirer{},
				switchingUsecase:             &Switching{},
				additionalInformationUsecase: &AdditionalInformation{},
				qrisTags:                     &QRISTags{},
				qrisCategoryContents:         &QRISCategoryContents{},
			},
			want: &Field{
				acquirerUsecase:              &Acquirer{},
				switchingUsecase:             &Switching{},
				additionalInformationUsecase: &AdditionalInformation{},
				qrisTags:                     &QRISTags{},
				qrisCategoryContents:         &QRISCategoryContents{},
			},
		},
		{
			name: "Success: With Field",
			fields: Field{
				acquirerUsecase:              &Acquirer{},
				switchingUsecase:             &Switching{},
				additionalInformationUsecase: &AdditionalInformation{},
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
			},
			want: &Field{
				acquirerUsecase:              &Acquirer{},
				switchingUsecase:             &Switching{},
				additionalInformationUsecase: &AdditionalInformation{},
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
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewField(test.fields.acquirerUsecase, test.fields.switchingUsecase, test.fields.additionalInformationUsecase, test.fields.qrisTags, test.fields.qrisCategoryContents)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewField", "FieldInterface")
			}

			got, ok := uc.(*Field)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*Field")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*Field", test.want, got)
			}
		})
	}
}

func TestFieldAssign(t *testing.T) {
	type args struct {
		qris *entities.QRIS
		data *entities.Data
	}

	tests := []struct {
		name      string
		fields    Field
		args      args
		wantError error
	}{
		{
			name:   "Success: Pass Version Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.Version,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Category Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.Category,
			},
			wantError: nil,
		},
		{
			name: "Error: uc.acquirerUsecase.Parse()",
			fields: Field{
				acquirerUsecase: &mockAcquirerUsecase{
					ParseFunc: func(content string) (*entities.AcquirerDetail, error) {
						return nil, fmt.Errorf("invalid parse acquirer for content %s", content)
					},
				},
			},
			args: args{
				data: &entities.Data{
					Tag:     testQRIS.Acquirer.Tag,
					Content: testQRIS.Acquirer.Content,
					Data:    testQRIS.Acquirer.Data,
				},
			},
			wantError: fmt.Errorf("invalid parse acquirer for content %s", testQRIS.Acquirer.Content),
		},
		{
			name: "Success: Pass Acquirer Tag",
			fields: Field{
				acquirerUsecase: &mockAcquirerUsecase{
					ParseFunc: func(content string) (*entities.AcquirerDetail, error) {
						return &testQRIS.Acquirer.Detail, nil
					},
				},
			},
			args: args{
				qris: &entities.QRIS{},
				data: &entities.Data{
					Tag:     testQRIS.Acquirer.Tag,
					Content: testQRIS.Acquirer.Content,
					Data:    testQRIS.Acquirer.Data,
				},
			},
			wantError: nil,
		},
		{
			name: "Error: uc.switchingUsecase.Parse()",
			fields: Field{
				switchingUsecase: &mockSwitchingUsecase{
					ParseFunc: func(content string) (*entities.SwitchingDetail, error) {
						return nil, fmt.Errorf("invalid parse switching for content %s", content)
					},
				},
			},
			args: args{
				data: &entities.Data{
					Tag:     testQRIS.Switching.Tag,
					Content: testQRIS.Switching.Content,
					Data:    testQRIS.Switching.Data,
				},
			},
			wantError: fmt.Errorf("invalid parse switching for content %s", testQRIS.Switching.Content),
		},
		{
			name: "Success: Pass Switching Tag",
			fields: Field{
				switchingUsecase: &mockSwitchingUsecase{
					ParseFunc: func(content string) (*entities.SwitchingDetail, error) {
						return &testQRIS.Switching.Detail, nil
					},
				},
			},
			args: args{
				qris: &entities.QRIS{},
				data: &entities.Data{
					Tag:     testQRIS.Switching.Tag,
					Content: testQRIS.Switching.Content,
					Data:    testQRIS.Switching.Data,
				},
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Merchant Category Code Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.MerchantCategoryCode,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Merchant Category Code Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.MerchantCategoryCode,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Currency Code Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.CurrencyCode,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Payment Amount Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.PaymentAmount,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Payment Fee Category Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.PaymentFeeCategory,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Payment Fee Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.PaymentFee,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Country Code Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.CountryCode,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Merchant Name Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.MerchantName,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Merchant City Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.MerchantCity,
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass Merchant Postal Code Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.MerchantPostalCode,
			},
			wantError: nil,
		},
		{
			name: "Error: uc.additionalInformationUsecase.Parse()",
			fields: Field{
				additionalInformationUsecase: &mockAdditionalInformationUsecase{
					ParseFunc: func(content string) (*entities.AdditionalInformationDetail, error) {
						return nil, fmt.Errorf("invalid parse additional information for content %s", content)
					},
				},
			},
			args: args{
				data: &entities.Data{
					Tag:     testQRIS.AdditionalInformation.Tag,
					Content: testQRIS.AdditionalInformation.Content,
					Data:    testQRIS.AdditionalInformation.Data,
				},
			},
			wantError: fmt.Errorf("invalid parse additional information for content %s", testQRIS.AdditionalInformation.Content),
		},
		{
			name: "Success: Pass Additional Information Tag",
			fields: Field{
				additionalInformationUsecase: &mockAdditionalInformationUsecase{
					ParseFunc: func(content string) (*entities.AdditionalInformationDetail, error) {
						return &testQRIS.AdditionalInformation.Detail, nil
					},
				},
			},
			args: args{
				qris: &entities.QRIS{},
				data: &entities.Data{
					Tag:     testQRIS.AdditionalInformation.Tag,
					Content: testQRIS.AdditionalInformation.Content,
					Data:    testQRIS.AdditionalInformation.Data,
				},
			},
			wantError: nil,
		},
		{
			name:   "Success: Pass CRC Code Tag",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{},
				data: &testQRIS.CRCCode,
			},
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Field{
				acquirerUsecase:              test.fields.acquirerUsecase,
				switchingUsecase:             test.fields.switchingUsecase,
				additionalInformationUsecase: test.fields.additionalInformationUsecase,
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
			}

			err := uc.Assign(test.args.qris, test.args.data)
			if (err != nil && test.wantError == nil) || (err == nil && test.wantError != nil) || (err != nil && err.Error() != test.wantError.Error()) {
				t.Errorf(expectedErrorButGotMessage, "Parse()", test.wantError, err)
			}
		})
	}
}

func TestFieldIsValid(t *testing.T) {
	type args struct {
		qris *entities.QRIS
	}

	tests := []struct {
		name   string
		fields Field
		args   args
		want   *[]string
	}{
		{
			name:   "Error: Category Content Undefined",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Version: testQRIS.Version,
					Category: entities.Data{
						Tag:     testQRIS.Category.Tag,
						Content: "1337",
						Data:    testQRIS.Category.Tag + "041337",
					},
					Acquirer:             testQRIS.Acquirer,
					Switching:            testQRIS.Switching,
					MerchantCategoryCode: testQRIS.MerchantCategoryCode,
					CurrencyCode:         testQRIS.CurrencyCode,
					CountryCode:          testQRIS.CountryCode,
					MerchantName:         testQRIS.MerchantName,
					MerchantCity:         testQRIS.MerchantCity,
					MerchantPostalCode:   testQRIS.MerchantPostalCode,
					CRCCode:              testQRIS.CRCCode,
				},
			},
			want: &[]string{
				"Category content undefined",
			},
		},
		{
			name:   "Error: Acquirer Tag Is Missing",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Version:              testQRIS.Version,
					Category:             testQRIS.Category,
					Acquirer:             entities.Acquirer{},
					MerchantCategoryCode: testQRIS.MerchantCategoryCode,
					CurrencyCode:         testQRIS.CurrencyCode,
					CountryCode:          testQRIS.CountryCode,
					MerchantName:         testQRIS.MerchantName,
					MerchantCity:         testQRIS.MerchantCity,
					MerchantPostalCode:   testQRIS.MerchantPostalCode,
					CRCCode:              testQRIS.CRCCode,
				},
			},
			want: &[]string{
				"Acquirer tag is missing",
			},
		},
		{
			name:   "Error: Switching Tag Is Missing",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Version:              testQRIS.Version,
					Category:             testQRIS.Category,
					Acquirer:             testQRIS.Acquirer,
					Switching:            entities.Switching{},
					MerchantCategoryCode: testQRIS.MerchantCategoryCode,
					CurrencyCode:         testQRIS.CurrencyCode,
					CountryCode:          testQRIS.CountryCode,
					MerchantName:         testQRIS.MerchantName,
					MerchantCity:         testQRIS.MerchantCity,
					MerchantPostalCode:   testQRIS.MerchantPostalCode,
					CRCCode:              testQRIS.CRCCode,
				},
			},
			want: &[]string{
				"Switching tag is missing",
			},
		},
		{
			name:   "Error: Payment Fee Category Tag Is Missing",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Version:              testQRIS.Version,
					Category:             testQRIS.Category,
					Acquirer:             testQRIS.Acquirer,
					Switching:            testQRIS.Switching,
					MerchantCategoryCode: testQRIS.MerchantCategoryCode,
					CurrencyCode:         testQRIS.CurrencyCode,
					PaymentFeeCategory:   entities.Data{},
					PaymentFee:           testQRIS.PaymentFee,
					CountryCode:          testQRIS.CountryCode,
					MerchantName:         testQRIS.MerchantName,
					MerchantCity:         testQRIS.MerchantCity,
					MerchantPostalCode:   testQRIS.MerchantPostalCode,
					CRCCode:              testQRIS.CRCCode,
				},
			},
			want: &[]string{
				"Payment fee category tag is missing",
			},
		},
		{
			name:   "Error: Payment Fee Tag Is Missing",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Version:              testQRIS.Version,
					Category:             testQRIS.Category,
					Acquirer:             testQRIS.Acquirer,
					Switching:            testQRIS.Switching,
					MerchantCategoryCode: testQRIS.MerchantCategoryCode,
					CurrencyCode:         testQRIS.CurrencyCode,
					PaymentFeeCategory:   testQRIS.PaymentFeeCategory,
					PaymentFee:           entities.Data{},
					CountryCode:          testQRIS.CountryCode,
					MerchantName:         testQRIS.MerchantName,
					MerchantCity:         testQRIS.MerchantCity,
					MerchantPostalCode:   testQRIS.MerchantPostalCode,
					CRCCode:              testQRIS.CRCCode,
				},
			},
			want: &[]string{
				"Payment fee tag is missing",
			},
		},
		{
			name:   "Error: CRC Code Tag Is Missing",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Version:              testQRIS.Version,
					Category:             testQRIS.Category,
					Acquirer:             testQRIS.Acquirer,
					Switching:            testQRIS.Switching,
					MerchantCategoryCode: testQRIS.MerchantCategoryCode,
					CurrencyCode:         testQRIS.CurrencyCode,
					CountryCode:          testQRIS.CountryCode,
					MerchantName:         testQRIS.MerchantName,
					MerchantCity:         testQRIS.MerchantCity,
					MerchantPostalCode:   testQRIS.MerchantPostalCode,
					CRCCode:              entities.Data{},
				},
			},
			want: &[]string{
				"CRC code tag is missing",
			},
		},
		{
			name:   "Error: Some Errors",
			fields: Field{},
			args: args{
				qris: &entities.QRIS{
					Acquirer: entities.Acquirer{
						Tag:     testQRIS.Acquirer.Tag,
						Content: testQRIS.Acquirer.Content,
						Data:    testQRIS.Acquirer.Data,
						Detail:  entities.AcquirerDetail{},
					},
					Switching: entities.Switching{
						Tag:     testQRIS.Switching.Tag,
						Content: testQRIS.Switching.Content,
						Data:    testQRIS.Switching.Data,
						Detail:  entities.SwitchingDetail{},
					},
				},
			},
			want: &[]string{
				"Version tag is missing",
				"Category tag is missing",
				"Category content undefined",
				"Acquirer site tag is missing",
				"Acquirer MPAN tag is missing",
				"Acquirer terminal id tag is missing",
				"Acquirer category tag is missing",
				"Switching site tag is missing",
				"Switching NMID tag is missing",
				"Switching category tag is missing",
				"Merchant category tag is missing",
				"Currency code tag is missing",
				"Country code tag is missing",
				"Merchant name tag is missing",
				"Merchant city tag is missing",
				"Merchant postal code tag is missing",
				"CRC code tag is missing",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Field{
				acquirerUsecase:  test.fields.acquirerUsecase,
				switchingUsecase: test.fields.switchingUsecase,
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
			}

			var got []string
			uc.IsValid(test.args.qris, &got)
			if !reflect.DeepEqual(&got, test.want) {
				t.Errorf(expectedButGotMessage, "IsValid()", *test.want, got)
			}
		})
	}
}
