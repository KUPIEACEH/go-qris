package usecases

import (
	"reflect"
	"testing"
)

func TestNewCRC16CCITT(t *testing.T) {
	tests := []struct {
		name string
		want CRC16CCITTInterface
	}{
		{
			name: "Success",
			want: &CRC16CCITT{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewCRC16CCITT()

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewCRC16CCITT", "CRC16CCITTInterface")
			}

			got, ok := uc.(*CRC16CCITT)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*CRC16CCITT")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*CRC16CCITT", test.want, got)
			}
		})
	}
}

func TestCRC16CCITTGenerateCode(t *testing.T) {
	type args struct {
		code string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				code: "123456789",
			},
			want: "29B1",
		},
		{
			args: args{
				code: "HELLO",
			},
			want: "49D6",
		},
		{
			args: args{
				code: "abcdef",
			},
			want: "34ED",
		},
		{
			args: args{
				code: "",
			},
			want: "FFFF",
		},
	}

	for _, test := range tests {
		t.Run(test.args.code, func(t *testing.T) {
			uc := &CRC16CCITT{}
			got := uc.GenerateCode(test.args.code)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "GenerateCode()", test.want, got)
			}
		})
	}
}
