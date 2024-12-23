package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/models"
	"github.com/fyvri/go-qris/pkg/utils"
)

func TestNewQRIS(t *testing.T) {
	qrisTags := &usecases.QRISTags{
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
	}
	qrisCategoryContents := &usecases.QRISCategoryContents{
		Static:  testCategoryStaticContent,
		Dynamic: testCategoryDynamicContent,
	}
	qrisPaymentFeeCategoryContents := &usecases.QRISPaymentFeeCategoryContents{
		Fixed:   testPaymentFeeCategoryFixedContent,
		Percent: testPaymentFeeCategoryPercentContent,
	}
	acquirerDetailTags := &usecases.AcquirerDetailTags{
		Site:       testAcquirerDetailSiteTag,
		MPAN:       testAcquirerDetailMPANTag,
		TerminalID: testAcquirerDetailTerminalIDTag,
		Category:   testAcquirerDetailCategoryTag,
	}
	switchingDetailTags := &usecases.SwitchingDetailTags{
		Site:     testSwitchingDetailSiteTag,
		NMID:     testSwitchingDetailNMIDTag,
		Category: testSwitchingDetailCategoryTag,
	}
	qrisAdditionalInformationDetailTags := &usecases.AdditionalInformationDetailTags{
		BillNumber:                    testAdditionalInformationDetailBillNumber,
		MobileNumber:                  testAdditionalInformationDetailMobileNumber,
		StoreLabel:                    testAdditionalInformationDetailStoreLabel,
		LoyaltyNumber:                 testAdditionalInformationDetailLoyaltyNumber,
		ReferenceLabel:                testAdditionalInformationDetailReferenceLabel,
		CustomerLabel:                 testAdditionalInformationDetailCustomerLabel,
		TerminalLabel:                 testAdditionalInformationDetailTerminalLabel,
		PurposeOfTransaction:          testAdditionalInformationDetailPurposeOfTransaction,
		AdditionalConsumerDataRequest: testAdditionalInformationDetailAdditionalConsumerDataRequest,
		MerchantTaxID:                 testAdditionalInformationDetailMerchantTaxID,
		MerchantChannel:               testAdditionalInformationDetailMerchantChannel,
		RFUStart:                      testAdditionalInformationDetailRFUStart,
		RFUEnd:                        testAdditionalInformationDetailRFUEnd,
		PaymentSystemSpecificStart:    testAdditionalInformationDetailPaymentSystemSpecificStart,
		PaymentSystemSpecificEnd:      testAdditionalInformationDetailPaymentSystemSpecificEnd,
	}

	dataUsecase := usecases.NewData()
	acquirerUsecase := usecases.NewAcquirer(dataUsecase, acquirerDetailTags)
	switchingUsecase := usecases.NewSwitching(dataUsecase, switchingDetailTags)
	additionalInformationUsecase := usecases.NewAdditionalInformation(dataUsecase, qrisAdditionalInformationDetailTags)
	fieldUsecase := usecases.NewField(acquirerUsecase, switchingUsecase, additionalInformationUsecase, qrisTags, qrisCategoryContents)
	paymentFeeUsecase := usecases.NewPaymentFee(qrisTags, qrisPaymentFeeCategoryContents)
	crc16CCITTUsecase := usecases.NewCRC16CCITT()

	qrisUsecases := &usecases.QRISUsecases{
		Data:                  dataUsecase,
		Field:                 fieldUsecase,
		PaymentFee:            paymentFeeUsecase,
		AdditionalInformation: additionalInformationUsecase,
		CRC16CCITT:            crc16CCITTUsecase,
	}
	qrisUsecase := usecases.NewQRIS(qrisUsecases, qrisTags, qrisCategoryContents, qrisPaymentFeeCategoryContents)
	inputUtil := utils.NewInput()

	tests := []struct {
		name string
		want QRISInterface
	}{
		{
			name: "Success",
			want: &QRIS{
				crc16CCITTUsecase: crc16CCITTUsecase,
				qrisUsecase:       qrisUsecase,
				inputUtil:         inputUtil,
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
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISEntityString
					},
				},
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
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISEntityString
					},
				},
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
				inputUtil:         test.fields.inputUtil,
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
				inputUtil:         test.fields.inputUtil,
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
		name      string
		fields    QRIS
		args      args
		want      *models.QRIS
		wantError error
	}{
		{
			name: "Error: Input Length",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						alphabet := "abcdefghijklmnopqrstuvwxyz"
						result := ""
						for i := 0; i < 100; i++ {
							result += string(alphabet[i%len(alphabet)])
						}

						return result
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
						return &testQRISEntityModified
					},
				},
			},
			args: args{
				qris: &testQRISModel,
			},
			want:      nil,
			wantError: fmt.Errorf("input length exceeds the maximum permitted characters"),
		},
		{
			name: "Success",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						switch input {
						case testMerchantCityContent:
							return testMerchantCityContent
						case testMerchantPostalCodeContent:
							return testMerchantPostalCodeContent
						case testPaymentFeeCategoryFixedContent:
							return testPaymentFeeCategoryFixedContent
						case testAdditionalInformationTerminalLabelContent:
							return testAdditionalInformationTerminalLabelContent
						default:
							return ""
						}
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
						return &testQRISEntityModified
					},
				},
			},
			args: args{
				qris: &testQRISModel,
			},
			want:      &testQRISModelModified,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &QRIS{
				crc16CCITTUsecase: test.fields.crc16CCITTUsecase,
				qrisUsecase:       test.fields.qrisUsecase,
				inputUtil:         test.fields.inputUtil,
			}

			got, err, _ := uc.Modify(test.args.qris, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, testPaymentFeeCategoryFixedContent, testPaymentFeeValue, testAdditionalInformationTerminalLabelContent)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Modify()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Modify()", test.want, got)
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
				inputUtil:         test.fields.inputUtil,
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
			name: "Error: Input Length",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						alphabet := "abcdefghijklmnopqrstuvwxyz"
						result := ""
						for i := 0; i < 100; i++ {
							result += string(alphabet[i%len(alphabet)])
						}

						return result
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &testQRISEntity, nil, nil
					},
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
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
			want:      "",
			wantError: fmt.Errorf("input length exceeds the maximum permitted characters"),
		},
		{
			name: "Error: s.qrisUsecase.Parse()",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						switch input {
						case testQRISEntityString:
							return testQRISEntityString
						case testMerchantCityContent:
							return testMerchantCityContent
						case testMerchantPostalCodeContent:
							return testMerchantPostalCodeContent
						case testPaymentFeeCategoryFixedContent:
							return testPaymentFeeCategoryFixedContent
						default:
							return ""
						}
					},
				},
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
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						switch input {
						case testQRISEntityString:
							return testQRISEntityString
						case testMerchantCityContent:
							return testMerchantCityContent
						case testMerchantPostalCodeContent:
							return testMerchantPostalCodeContent
						case testPaymentFeeCategoryFixedContent:
							return testPaymentFeeCategoryFixedContent
						default:
							return ""
						}
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &testQRISEntity, nil, nil
					},
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
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
				inputUtil:         test.fields.inputUtil,
			}

			got, err, _ := uc.Convert(test.args.qrString, testMerchantCityContent, testMerchantPostalCodeContent, testPaymentAmountValue, testPaymentFeeCategoryFixedContent, testPaymentFeeValue, testAdditionalInformationTerminalLabelContent)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Convert()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Convert()", test.want, got)
			}
		})
	}
}
