package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/pkg/models"
)

// func TestNewQRIS(t *testing.T) {
// 	type args struct {
// 		schema *Schema
// 	}

// 	tests := []struct {
// 		name   string
// 		fields QRIS
// 		args   args
// 		want   QRISInterface
// 	}{
// 		{
// 			name: "Success: No Field",
// 			fields: QRIS{
// 				crc16CCITTUsecase: &usecases.CRC16CCITT{},
// 				qrisUsecase:       &usecases.QRIS{},
// 			},
// 			args: args{
// 				schema: nil,
// 			},
// 			want: &QRIS{
// 				crc16CCITTUsecase: &usecases.CRC16CCITT{},
// 				qrisUsecase:       &usecases.QRIS{},
// 			},
// 		},
// 		{
// 			name: "Success: With Field",
// 			fields: QRIS{
// 				crc16CCITTUsecase: &usecases.CRC16CCITT{},
// 				qrisUsecase:       &usecases.QRIS{},
// 			},
// 			args: args{
// 				schema: &Schema{
// 					VersionTag:               testVersionTag,
// 					CategoryTag:              testCategoryTag,
// 					AcquirerTag:              testAcquirerTag,
// 					AcquirerBankTransferTag:  testAcquirerBankTransferTag,
// 					SwitchingTag:             testSwitchingTag,
// 					MerchantCategoryCodeTag:  testMerchantCategoryCodeTag,
// 					CurrencyCodeTag:          testCurrencyCodeTag,
// 					PaymentAmountTag:         testPaymentAmountTag,
// 					PaymentFeeCategoryTag:    testPaymentFeeCategoryTag,
// 					PaymentFeeFixedTag:       testPaymentFeeFixedTag,
// 					PaymentFeePercentTag:     testPaymentFeePercentTag,
// 					CountryCodeTag:           testCountryCodeTag,
// 					MerchantNameTag:          testMerchantNameTag,
// 					MerchantCityTag:          testMerchantCityTag,
// 					MerchantPostalCodeTag:    testMerchantPostalCodeTag,
// 					AdditionalInformationTag: testAdditionalInformationTag,
// 					CRCCodeTag:               testCRCCodeTag,

// 					CategoryStaticContent:  testCategoryStaticContent,
// 					CategoryDynamicContent: testCategoryDynamicContent,

// 					AcquirerDetailSiteTag:       testAcquirerDetailSiteTag,
// 					AcquirerDetailMPANTag:       testAcquirerDetailMPANTag,
// 					AcquirerDetailTerminalIDTag: testAcquirerDetailTerminalIDTag,
// 					AcquirerDetailCategoryTag:   testAcquirerDetailCategoryTag,

// 					SwitchingDetailSiteTag:     testSwitchingDetailSiteTag,
// 					SwitchingDetailNMIDTag:     testSwitchingDetailNMIDTag,
// 					SwitchingDetailCategoryTag: testSwitchingDetailCategoryTag,

// 					PaymentFeeCategoryFixedContent:   testPaymentFeeCategoryFixedContent,
// 					PaymentFeeCategoryPercentContent: testPaymentFeeCategoryPercentContent,
// 				},
// 			},
// 			want: &QRIS{
// 				crc16CCITTUsecase: &usecases.CRC16CCITT{},
// 				qrisUsecase:       &usecases.QRIS{},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			s := NewQRIS(test.args.schema)

// 			if s == nil {
// 				t.Errorf(expectedReturnNonNil, "NewQRIS", "QRISInterface")
// 			}

// 			got, ok := s.(*QRIS)
// 			if !ok {
// 				t.Errorf(expectedTypeAssertionErrorMessage, "*QRIS")
// 			}

// 			if !reflect.DeepEqual(test.want, got) {
// 				t.Logf("Want: Type: %T, Value: %s", test.want, test.want)
// 				t.Logf("Got : Type: %T, Value: %s", got, *got)
// 				t.Errorf(expectedButGotMessage, "*QRIS", test.want, got)
// 			}
// 		})
// 	}
// }

func TestQRISParse(t *testing.T) {
	type args struct {
		qrString string
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		want      *models.QRIS
		wantError error
	}{
		{
			name: "Error: s.qrisUsecase.Parse()",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return nil, fmt.Errorf("invalid format code"), nil
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
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &testQRIS, nil, nil
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			want:      &testQRISModel,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
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

func TestQRISValidate(t *testing.T) {
	type args struct {
		qris *models.QRIS
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   bool
	}{
		{
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ValidateFunc: func(qris *entities.QRIS) bool {
						return true
					},
				},
			},
			args: args{
				qris: &testQRISModel,
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got := uc.Validate(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Validate()", test.want, got)
			}
		})
	}
}

func TestQRISToDynamic(t *testing.T) {
	type args struct {
		qris               models.QRIS
		paymentFeeCategory string
		paymentFee         uint32
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   *models.QRISDynamic
	}{
		{
			name: "Success: Fixed Payment Fee",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ToDynamicFunc: func(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic {
						return &testQRISDynamic
					},
				},
			},
			args: args{
				qris:               *mapQRISEntityToModel(&testQRIS),
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
			},
			want: mapQRISDynamicEntityToModel(&testQRISDynamic),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got := uc.ToDynamic(&test.args.qris, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, test.args.paymentFeeCategory, test.args.paymentFee)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToDynamic()", test.want, got)
			}
		})
	}
}

func TestQRISToString(t *testing.T) {
	type args struct {
		qris models.QRIS
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   string
	}{
		{
			name: "Success: Fixed Payment Fee",
			fields: QRIS{
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "AZ15"
					},
				},
			},
			args: args{
				qris: *mapQRISEntityToModel(&testQRIS),
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
				testQRIS.CRCCode.Tag + "04AZ15",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got := uc.ToString(&test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToString()", test.want, got)
			}
		})
	}
}

func TestQRISConvert(t *testing.T) {
	type args struct {
		qrString           string
		paymentFeeCategory string
		paymentFee         uint32
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		want      string
		wantError error
	}{
		{
			name: "Error: s.qrisUsecase.Parse()",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return nil, fmt.Errorf("invalid extract acquirer for content %s", testQRIS.Acquirer.Content), nil
					},
				},
			},
			args: args{
				qrString:           testQRISString,
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
			},
			want:      "",
			wantError: fmt.Errorf("invalid extract acquirer for content %s", testQRIS.Acquirer.Content),
		},
		{
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &testQRIS, nil, nil
					},
					ToDynamicFunc: func(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic {
						return &testQRISDynamic
					},
					DynamicToStringFunc: func(qrisDynamic *entities.QRISDynamic) string {
						return testQRISDynamicString
					},
				},
			},
			args: args{
				qrString:           testQRISString,
				paymentFeeCategory: "FIXED",
				paymentFee:         666,
			},
			want:      testQRISDynamicString,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got, err, _ := uc.Convert(test.args.qrString, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, test.args.paymentFeeCategory, test.args.paymentFee)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Convert()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Convert()", test.want, got)
			}
		})
	}
}
