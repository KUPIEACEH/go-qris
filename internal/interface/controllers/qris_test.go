package controllers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/utils"
)

func TestNewQRIS(t *testing.T) {
	tests := []struct {
		name   string
		fields QRIS
		want   QRISInterface
	}{
		{
			name:   "Success: No Field",
			fields: QRIS{},
			want:   &QRIS{},
		},
		{
			name: "Success: With Field",
			fields: QRIS{
				inputUtil:   &utils.Input{},
				qrCodeUtil:  &utils.QRCode{},
				qrisUsecase: &usecases.QRIS{},
				qrCodeSize:  testQRCodeSize,
			},
			want: &QRIS{
				inputUtil:   &utils.Input{},
				qrCodeUtil:  &utils.QRCode{},
				qrisUsecase: &usecases.QRIS{},
				qrCodeSize:  testQRCodeSize,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRIS(test.fields.inputUtil, test.fields.qrCodeUtil, test.fields.qrisUsecase, test.fields.qrCodeSize)

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
			name: testNameErrorParse,
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISString
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return nil, fmt.Errorf(testErrMessageInvalidFormatCode), nil
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			want:      nil,
			wantError: fmt.Errorf(testErrMessageInvalidFormatCode),
		},
		{
			name: "Success",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISString
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}, nil, nil
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			want: &entities.QRIS{
				Version: entities.Data{
					Tag:     "",
					Content: "01",
					Data:    "000201",
				},
			},
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &QRIS{
				inputUtil:   test.fields.inputUtil,
				qrCodeUtil:  test.fields.qrCodeUtil,
				qrisUsecase: test.fields.qrisUsecase,
				qrCodeSize:  test.fields.qrCodeSize,
			}

			got, err, _ := c.Parse(test.args.qrString)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Parse()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Parse()", test.want, got)
			}
		})
	}
}

func TestQRISConvert(t *testing.T) {
	type args struct {
		qrString           string
		merchantCity       string
		merchantPostalCode string
		paymentAmount      uint32
		paymentFeeCategory string
		paymentFee         uint32
		terminalLabel      string
	}

	type want struct {
		qrString string
		qrCode   string
	}

	testMerchantCity := "Kota Yogyakarta"
	testMerchantPostalCode := "55000"
	testPaymentAmount := uint32(1337)
	testPaymentFeeCategory := "FIXED"
	testPaymentFee := uint32(666)
	testTerminalLabel := "A01"

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	testOverInputLength := ""
	for i := 0; i < 100; i++ {
		testOverInputLength += string(alphabet[i%len(alphabet)])
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		want      want
		wantError error
	}{
		{
			name: "Error: Input Length",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testOverInputLength
					},
				},
			},
			args: args{
				qrString:           testQRISString,
				merchantCity:       testOverInputLength,
				merchantPostalCode: testOverInputLength,
				paymentAmount:      testPaymentAmount,
				paymentFeeCategory: testPaymentFeeCategory,
				paymentFee:         testPaymentFee,
				terminalLabel:      testOverInputLength,
			},
			want: want{
				qrString: "",
				qrCode:   "",
			},
			wantError: fmt.Errorf("input length exceeds the maximum permitted characters"),
		},
		{
			name: testNameErrorParse,
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISString
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return nil, fmt.Errorf("invalid QRIS format"), nil
					},
				},
			},
			args: args{
				qrString:           testQRISString,
				merchantCity:       testMerchantCity,
				merchantPostalCode: testMerchantPostalCode,
				paymentAmount:      testPaymentAmount,
				paymentFeeCategory: testPaymentFeeCategory,
				paymentFee:         testPaymentFee,
			},
			want: want{
				qrString: "",
				qrCode:   "",
			},
			wantError: fmt.Errorf("invalid QRIS format"),
		},
		{
			name: "Error: c.qrCodeUtil.StringToImageBase64()",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return input
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}, nil, nil
					},
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}
					},
					ToStringFunc: func(qris *entities.QRIS) string {
						return testQRISModifiedString
					},
				},
				qrCodeUtil: &mockQRCodeUtil{
					StringToImageBase64Func: func(qrString string, qrCodeSize int) (string, error) {
						return "", fmt.Errorf("unsupported QR code format")
					},
				},
			},
			args: args{
				qrString:           testQRISString,
				merchantCity:       testMerchantCity,
				merchantPostalCode: testMerchantPostalCode,
				paymentAmount:      testPaymentAmount,
				paymentFeeCategory: testPaymentFeeCategory,
				paymentFee:         testPaymentFee,
			},
			want: want{
				qrString: testQRISModifiedString,
				qrCode:   "",
			},
			wantError: fmt.Errorf("unsupported QR code format"),
		},
		{
			name: "Success",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						switch input {
						case testQRISString:
							return testQRISString
						case testMerchantCity:
							return testMerchantCity
						case testMerchantPostalCode:
							return testMerchantPostalCode
						case testPaymentFeeCategory:
							return testPaymentFeeCategory
						default:
							return ""
						}
					},
				},
				qrCodeUtil: &mockQRCodeUtil{
					StringToImageBase64Func: func(qrString string, qrCodeSize int) (string, error) {
						return "data:image/png;base64,QRIS Modified Code Image Base64", nil
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}, nil, nil
					},
					ModifyFunc: func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) *entities.QRIS {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}
					},
					ToStringFunc: func(qris *entities.QRIS) string {
						return testQRISModifiedString
					},
				},
				qrCodeSize: 125,
			},
			args: args{
				qrString:           testQRISString,
				merchantCity:       testMerchantCity,
				merchantPostalCode: testMerchantPostalCode,
				paymentAmount:      testPaymentAmount,
				paymentFeeCategory: testPaymentFeeCategory,
				paymentFee:         testPaymentFee,
				terminalLabel:      testTerminalLabel,
			},
			want: want{
				qrString: testQRISModifiedString,
				qrCode:   "data:image/png;base64,QRIS Modified Code Image Base64",
			},
			wantError: nil,
		},
	}

	funcName := "Convert()"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &QRIS{
				inputUtil:   test.fields.inputUtil,
				qrCodeUtil:  test.fields.qrCodeUtil,
				qrisUsecase: test.fields.qrisUsecase,
				qrCodeSize:  test.fields.qrCodeSize,
			}

			got1, got2, err, _ := c.Convert(test.args.qrString, test.args.merchantCity, test.args.merchantPostalCode, test.args.paymentAmount, test.args.paymentFeeCategory, test.args.paymentFee, test.args.terminalLabel)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, funcName, test.wantError, err)
			}
			if !reflect.DeepEqual(got1, test.want.qrString) {
				t.Errorf(expectedButGotMessage, funcName, test.want, got1)
			}
			if !reflect.DeepEqual(got2, test.want.qrCode) {
				t.Errorf(expectedButGotMessage, funcName, test.want, got2)
			}
		})
	}
}

func TestQRISIsValid(t *testing.T) {
	type args struct {
		qrString string
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		wantError error
	}{
		{
			name: testNameErrorParse,
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISString
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return nil, fmt.Errorf(testErrMessageInvalidFormatCode), nil
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			wantError: fmt.Errorf(testErrMessageInvalidFormatCode),
		},
		{
			name: "Success: c.qrisUsecase.IsValid() Is False",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISString
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}, nil, nil
					},
					IsValidFunc: func(qris *entities.QRIS) bool {
						return false
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			wantError: fmt.Errorf("invalid CRC16-CCITT code"),
		},
		{
			name: "Success",
			fields: QRIS{
				inputUtil: &mockInputUtil{
					SanitizeFunc: func(input string) string {
						return testQRISString
					},
				},
				qrisUsecase: &mockQRISUsecase{
					ParseFunc: func(qrString string) (*entities.QRIS, error, *[]string) {
						return &entities.QRIS{
							Version: entities.Data{
								Tag:     "",
								Content: "01",
								Data:    "000201",
							},
						}, nil, nil
					},
					IsValidFunc: func(qris *entities.QRIS) bool {
						return true
					},
				},
			},
			args: args{
				qrString: testQRISString,
			},
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &QRIS{
				inputUtil:   test.fields.inputUtil,
				qrCodeUtil:  test.fields.qrCodeUtil,
				qrisUsecase: test.fields.qrisUsecase,
				qrCodeSize:  test.fields.qrCodeSize,
			}

			err, _ := c.IsValid(test.args.qrString)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "IsValid()", test.wantError, err)
			}
		})
	}
}
