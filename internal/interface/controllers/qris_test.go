package controllers

import (
	"errors"
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
				qrisUsecase: &usecases.QRIS{},
				qrCodeUtil:  &utils.QRCode{},
				qrCodeSize:  testQRCodeSize,
			},
			want: &QRIS{
				qrisUsecase: &usecases.QRIS{},
				qrCodeUtil:  &utils.QRCode{},
				qrCodeSize:  testQRCodeSize,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRIS(test.fields.qrisUsecase, test.fields.qrCodeUtil, test.fields.qrCodeSize)

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

func TestQRISSanitizeInput(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success: Trim Space",
			args: args{
				input: "  hello world  ",
			},
			want: "hello world",
		},
		{
			name: "Success: With Word",
			args: args{
				input: "\nhello\nworld\r",
			},
			want: "helloworld",
		},
		{
			name: "Success: No Word",
			args: args{
				input: "   \n\r  ",
			},
			want: "",
		},
		{
			name: "Success: No Replace",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sanitizeInput(test.args.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "sanitizeInput()", test.want, got)
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
			name: "Error: No String",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{},
			},
			args: args{
				qrString: "",
			},
			want:      nil,
			wantError: fmt.Errorf("QRIS not found"),
		},
		{
			name: "Error: ExtractStatic",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ExtractStaticFunc: func(qrString string) (*entities.QRISStatic, error) {
						return nil, errors.New("invalid format code")
					},
				},
			},
			args: args{
				qrString: testQRISStatic.Acquirer.Content,
			},
			want:      nil,
			wantError: errors.New("invalid format code"),
		},
		{
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ExtractStaticFunc: func(qrString string) (*entities.QRISStatic, error) {
						return &testQRISStatic, nil
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
			c := &QRIS{
				qrisUsecase: test.fields.qrisUsecase,
				qrCodeUtil:  test.fields.qrCodeUtil,
				qrCodeSize:  test.fields.qrCodeSize,
			}

			got, err := c.ExtractStatic(test.args.qrString)
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
	type args struct {
		qrString           string
		merchantName       string
		merchantCity       string
		merchantPostalCode string
		paymentAmount      uint32
		paymentFeeCategory string
		paymentFee         uint32
	}

	type want struct {
		qrString string
		qrCode   string
	}

	tests := []struct {
		name      string
		fields    QRIS
		args      args
		want      want
		wantError error
	}{
		{
			name: "Error: No QR String",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{},
			},
			args: args{
				qrString: "",
			},
			want: want{
				qrString: "",
				qrCode:   "",
			},
			wantError: fmt.Errorf("QRIS not found"),
		},
		{
			name: "Error: ExtractStatic",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ExtractStaticFunc: func(qrString string) (*entities.QRISStatic, error) {
						return nil, fmt.Errorf("not static QRIS format detected")
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want: want{
				qrString: "",
				qrCode:   "",
			},
			wantError: fmt.Errorf("not static QRIS format detected"),
		},
		{
			name: "Error: StringToImageBase64",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ExtractStaticFunc: func(qrString string) (*entities.QRISStatic, error) {
						return &testQRISStatic, nil
					},
					StaticToDynamicFunc: func(qrisStatic *entities.QRISStatic, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) *entities.QRISDynamic {
						return &entities.QRISDynamic{}
					},
					DynamicToDynamicStringFunc: func(QRISDynamic *entities.QRISDynamic) string {
						return testQRISStaticString
					},
				},
				qrCodeUtil: &mockQRCodeUtil{
					StringToImageBase64Func: func(qrString string, qrCodeSize int) (string, error) {
						return "", fmt.Errorf("unsupported qr code format")
					},
				},
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want: want{
				qrString: testQRISStaticString,
				qrCode:   "",
			},
			wantError: errors.New("unsupported qr code format"),
		},
		{
			name: "Success",
			fields: QRIS{
				qrisUsecase: &mockQRISUsecase{
					ExtractStaticFunc: func(qrString string) (*entities.QRISStatic, error) {
						return &testQRISStatic, nil
					},
					StaticToDynamicFunc: func(qrisStatic *entities.QRISStatic, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) *entities.QRISDynamic {
						return &entities.QRISDynamic{}
					},
					DynamicToDynamicStringFunc: func(QRISDynamic *entities.QRISDynamic) string {
						return testQRISStaticString
					},
				},
				qrCodeUtil: &mockQRCodeUtil{
					StringToImageBase64Func: func(qrString string, qrCodeSize int) (string, error) {
						return "data:image/png;base64,aZ15aLVr1y4nt0", nil
					},
				},
				qrCodeSize: 125,
			},
			args: args{
				qrString: testQRISStaticString,
			},
			want: want{
				qrString: testQRISStaticString,
				qrCode:   "data:image/png;base64,aZ15aLVr1y4nt0",
			},
			wantError: nil,
		},
	}

	funcName := "StaticToDynamic"
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &QRIS{
				qrisUsecase: test.fields.qrisUsecase,
				qrCodeUtil:  test.fields.qrCodeUtil,
				qrCodeSize:  test.fields.qrCodeSize,
			}

			got1, got2, err := c.StaticToDynamic(test.args.qrString, test.args.merchantName, test.args.merchantCity, test.args.merchantPostalCode, test.args.paymentAmount, test.args.paymentFeeCategory, test.args.paymentFee)
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
