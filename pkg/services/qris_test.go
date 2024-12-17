package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/config"
	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/models"
)

func TestNewQRIS(t *testing.T) {
	qrisTags := &usecases.QRISTags{
		VersionTag:               config.VersionTag,
		CategoryTag:              config.CategoryTag,
		AcquirerTag:              config.AcquirerTag,
		AcquirerBankTransferTag:  config.AcquirerBankTransferTag,
		SwitchingTag:             config.SwitchingTag,
		MerchantCategoryCodeTag:  config.MerchantCategoryCodeTag,
		CurrencyCodeTag:          config.CurrencyCodeTag,
		PaymentAmountTag:         config.PaymentAmountTag,
		PaymentFeeCategoryTag:    config.PaymentFeeCategoryTag,
		PaymentFeeFixedTag:       config.PaymentFeeFixedTag,
		PaymentFeePercentTag:     config.PaymentFeePercentTag,
		CountryCodeTag:           config.CountryCodeTag,
		MerchantNameTag:          config.MerchantNameTag,
		MerchantCityTag:          config.MerchantCityTag,
		MerchantPostalCodeTag:    config.MerchantPostalCodeTag,
		AdditionalInformationTag: config.AdditionalInformationTag,
		CRCCodeTag:               config.CRCCodeTag,
	}
	qrisCategoryContents := &usecases.QRISCategoryContents{
		Static:  config.CategoryStaticContent,
		Dynamic: config.CategoryDynamicContent,
	}
	qrisPaymentFeeCategoryContents := &usecases.QRISPaymentFeeCategoryContents{
		Fixed:   config.PaymentFeeCategoryFixedContent,
		Percent: config.PaymentFeeCategoryPercentContent,
	}

	dataUsecase := usecases.NewData()
	acquirerUsecase := usecases.NewAcquirer(dataUsecase, config.AcquirerDetailSiteTag, config.AcquirerDetailMPANTag, config.AcquirerDetailTerminalIDTag, config.AcquirerDetailCategoryTag)
	switchingUsecase := usecases.NewSwitching(dataUsecase, config.SwitchingDetailSiteTag, config.SwitchingDetailNMIDTag, config.SwitchingDetailCategoryTag)
	fieldUsecase := usecases.NewField(acquirerUsecase, switchingUsecase, qrisTags, qrisCategoryContents)
	crc16CCITTUsecase := usecases.NewCRC16CCITT()
	qrisUsecase := usecases.NewQRIS(
		dataUsecase,
		fieldUsecase,
		crc16CCITTUsecase,
		qrisTags,
		qrisCategoryContents,
		qrisPaymentFeeCategoryContents,
	)

	tests := []struct {
		name string
		want QRISInterface
	}{
		{
			name: "Success",
			want: &QRIS{
				crc16CCITTUsecase: crc16CCITTUsecase,
				qrisUsecase:       qrisUsecase,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewQRIS()

			if s == nil {
				t.Errorf(expectedReturnNonNil, "NewQRIS", "QRISInterface")
			}

			got, ok := s.(*QRIS)
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
				qrString: testQRISEntityString,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid format code"),
		},
		{
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &testQRISEntity, nil, nil
					},
				},
			},
			args: args{
				qrString: testQRISEntityString,
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

func TestQRISIsValid(t *testing.T) {
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
					IsValidFunc: func(qris *entities.QRIS) bool {
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

			got := uc.IsValid(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "IsValid()", test.want, got)
			}
		})
	}
}

func TestQRISModify(t *testing.T) {
	type args struct {
		qris *models.QRIS
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   *models.QRIS
	}{
		{
			name: "Success: Fixed Payment Fee",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
						return &testQRISEntityModified
					},
				},
			},
			args: args{
				qris: &testQRISModel,
			},
			want: &testQRISModelModified,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got := uc.Modify(test.args.qris, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, testPaymentFeeCategoryFixedContent, testPaymentFeeValue)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToDynamic()", test.want, got)
			}
		})
	}
}

func TestQRISToString(t *testing.T) {
	type args struct {
		qris *models.QRIS
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
				crc16CCITTUsecase: &mockCRC16CCITTUsecase{
					GenerateCodeFunc: func(code string) string {
						return "AZ15"
					},
				},
			},
			args: args{
				qris: &testQRISModel,
			},
			want: testQRISModel.Version.Data +
				testQRISModel.Category.Data +
				testQRISModel.Acquirer.Data +
				testQRISModel.Switching.Data +
				testQRISModel.MerchantCategoryCode.Data +
				testQRISModel.CurrencyCode.Data +
				testQRISModel.PaymentAmount.Data +
				testQRISModel.PaymentFeeCategory.Data +
				testQRISModel.PaymentFee.Data +
				testQRISModel.CountryCode.Data +
				testQRISModel.MerchantName.Data +
				testQRISModel.MerchantCity.Data +
				testQRISModel.MerchantPostalCode.Data +
				testQRISModel.AdditionalInformation.Data +
				testQRISModel.CRCCode.Tag + "04AZ15",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got := uc.ToString(test.args.qris)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToString()", test.want, got)
			}
		})
	}
}

func TestQRISConvert(t *testing.T) {
	type args struct {
		qrString string
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
						return nil, fmt.Errorf("invalid extract acquirer for content %s", testQRISEntity.Acquirer.Content), nil
					},
				},
			},
			args: args{
				qrString: testQRISEntityString,
			},
			want:      "",
			wantError: fmt.Errorf("invalid extract acquirer for content %s", testQRISEntity.Acquirer.Content),
		},
		{
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &testQRISEntity, nil, nil
					},
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
						return &testQRISEntityModified
					},
					ToStringFunc: func(qris *entities.QRIS) string {
						return testQRISEntityModifiedString
					},
				},
			},
			args: args{
				qrString: testQRISModelModifiedString,
			},
			want:      testQRISModelModifiedString,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
			}

			got, err, _ := uc.Convert(test.args.qrString, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, testPaymentFeeCategoryFixedContent, testPaymentFeeValue)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Convert()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Convert()", test.want, got)
			}
		})
	}
}
