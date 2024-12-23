package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewPaymentFee(t *testing.T) {
	tests := []struct {
		name   string
		fields PaymentFee
		want   PaymentFeeInterface
	}{
		{
			name: "Success: No Field",
			fields: PaymentFee{
				qrisTags:                       &QRISTags{},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{},
			},
			want: &PaymentFee{
				qrisTags:                       &QRISTags{},
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{},
			},
		},
		{
			name: "Success: With Field",
			fields: PaymentFee{
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
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			},
			want: &PaymentFee{
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
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewPaymentFee(test.fields.qrisTags, test.fields.qrisPaymentFeeCategoryContents)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewPaymentFee", "PaymentFeeInterface")
			}

			got, ok := uc.(*PaymentFee)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*PaymentFee")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*PaymentFee", test.want, got)
			}
		})
	}
}

func TestPaymentFeeModify(t *testing.T) {
	type args struct {
		qris               *entities.QRIS
		paymentFeeCategory string
		paymentFee         uint32
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   *entities.QRIS
	}{
		{
			name:   "Success: Fixed Payment Fee",
			fields: QRIS{},
			args: args{
				qris: &entities.QRIS{
					Version: testQRIS.Version,
					Category: entities.Data{
						Tag:     testCategoryTag,
						Content: testCategoryDynamicContent,
						Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
					},
					Acquirer:              testQRIS.Acquirer,
					Switching:             testQRIS.Switching,
					MerchantCategoryCode:  testQRIS.MerchantCategoryCode,
					CurrencyCode:          testQRIS.CurrencyCode,
					PaymentAmount:         testQRIS.PaymentAmount,
					PaymentFeeCategory:    testQRIS.PaymentFeeCategory,
					PaymentFee:            testQRIS.PaymentFee,
					CountryCode:           testQRIS.CountryCode,
					MerchantName:          testQRIS.MerchantName,
					MerchantCity:          testQRIS.MerchantCity,
					MerchantPostalCode:    testQRIS.MerchantPostalCode,
					AdditionalInformation: testQRIS.AdditionalInformation,
					CRCCode:               testQRIS.CRCCode,
				},
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
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
				PaymentAmount:        testQRIS.PaymentAmount,
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
				MerchantCity:          testQRIS.MerchantCity,
				MerchantPostalCode:    testQRIS.MerchantPostalCode,
				AdditionalInformation: testQRIS.AdditionalInformation,
				CRCCode:               testQRIS.CRCCode,
			},
		},
		{
			name:   "Success: Percent Payment Fee",
			fields: QRIS{},
			args: args{
				qris: &entities.QRIS{
					Version: testQRIS.Version,
					Category: entities.Data{
						Tag:     testCategoryTag,
						Content: testCategoryDynamicContent,
						Data:    testCategoryTag + fmt.Sprintf("%02d", len(testCategoryDynamicContent)) + testCategoryDynamicContent,
					},
					Acquirer:              testQRIS.Acquirer,
					Switching:             testQRIS.Switching,
					MerchantCategoryCode:  testQRIS.MerchantCategoryCode,
					CurrencyCode:          testQRIS.CurrencyCode,
					PaymentAmount:         testQRIS.PaymentAmount,
					PaymentFeeCategory:    testQRIS.PaymentFeeCategory,
					PaymentFee:            testQRIS.PaymentFee,
					CountryCode:           testQRIS.CountryCode,
					MerchantName:          testQRIS.MerchantName,
					MerchantCity:          testQRIS.MerchantCity,
					MerchantPostalCode:    testQRIS.MerchantPostalCode,
					AdditionalInformation: testQRIS.AdditionalInformation,
					CRCCode:               testQRIS.CRCCode,
				},
				paymentFeeCategory: "PERCENT",
				paymentFee:         25,
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
				PaymentAmount:        testQRIS.PaymentAmount,
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
				MerchantCity:          testQRIS.MerchantCity,
				MerchantPostalCode:    testQRIS.MerchantPostalCode,
				AdditionalInformation: testQRIS.AdditionalInformation,
				CRCCode:               testQRIS.CRCCode,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &PaymentFee{
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
				qrisPaymentFeeCategoryContents: &QRISPaymentFeeCategoryContents{
					Fixed:   testPaymentFeeCategoryFixedContent,
					Percent: testPaymentFeeCategoryPercentContent,
				},
			}

			got := uc.Modify(test.args.qris, test.args.paymentFeeCategory, test.args.paymentFee)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Modify()", test.want, got)
			}
		})
	}
}
