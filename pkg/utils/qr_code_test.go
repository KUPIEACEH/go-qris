package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewQRCode(t *testing.T) {
	tests := []struct {
		name   string
		fields QRCode
		want   QRCodeInterface
	}{
		{
			name:   "Success",
			fields: QRCode{},
			want:   &QRCode{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := NewQRCode()

			if u == nil {
				t.Errorf(expectedReturnNonNil, "NewQRCode", "QRCodeInterface")
			}

			got, ok := u.(*QRCode)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*QRCode")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*QRCode", test.want, got)
			}
		})
	}
}

func TestQRCODEStringToImageBase64(t *testing.T) {
	type args struct {
		qrString   string
		qrCodeSize int
	}

	tests := []struct {
		name      string
		fields    QRCode
		args      args
		want      string
		wantError error
	}{
		{
			name:   "Error: QR Code Scale",
			fields: QRCode{},
			args: args{
				qrString:   testQRString,
				qrCodeSize: -1,
			},
			want:      "",
			wantError: fmt.Errorf("can not scale barcode to an image smaller than 21x21"),
		},
		{
			name:   "Success",
			fields: QRCode{},
			args: args{
				qrString:   testQRString,
				qrCodeSize: 125,
			},
			want:      testQRCode,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := test.fields

			got, err := u.StringToImageBase64(test.args.qrString, test.args.qrCodeSize)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "StringToImageBase64()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "StringToImageBase64()", test.want, got)
			}
		})
	}
}
