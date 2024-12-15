package usecases

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewData(t *testing.T) {
	tests := []struct {
		name string
		want DataInterface
	}{
		{
			name: "Success",
			want: &Data{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewData()

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewData", "DataInterface")
			}

			got, ok := uc.(*Data)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*Data")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*Data", test.want, got)
			}
		})
	}
}
func TestDataExtract(t *testing.T) {
	type args struct {
		codeString string
	}

	tests := []struct {
		name      string
		fields    Data
		args      args
		want      *entities.ExtractData
		wantError error
	}{
		{
			name:   "Error: Content Length < 5",
			fields: Data{},
			args: args{
				codeString: "1102",
			},
			want:      nil,
			wantError: errors.New("invalid format code"),
		},
		{
			name:   "Error: Content Length Format",
			fields: Data{},
			args: args{
				codeString: testQRISStatic.Version.Tag + "X211",
			},
			want:      nil,
			wantError: fmt.Errorf("invalid length format for tag %s: strconv.Atoi: parsing \"X2\": invalid syntax", testQRISStatic.Version.Tag),
		},
		{
			name:   "Error: Content Length Not Match With Field",
			fields: Data{},
			args: args{
				codeString: testQRISStatic.Version.Tag + "021",
			},
			want:      nil,
			wantError: fmt.Errorf("invalid length for tag %s", testQRISStatic.Version.Tag),
		},
		{
			name:   "Success",
			fields: Data{},
			args: args{
				codeString: testQRISStatic.Version.Data,
			},
			want:      &testQRISStatic.Version,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := test.fields

			got, err := uc.Extract(test.args.codeString)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Extract()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Extract()", test.want, got)
			}
		})
	}
}

func TestDataModifyContent(t *testing.T) {
	type args struct {
		extractData entities.ExtractData
		content     string
	}

	tests := []struct {
		name string
		args args
		want entities.ExtractData
	}{
		{
			name: "No Content",
			args: args{
				extractData: entities.ExtractData{},
				content:     "",
			},
			want: entities.ExtractData{
				Tag:     "",
				Content: "",
				Data:    "00",
			},
		},
		{
			name: "With Content",
			args: args{
				extractData: entities.ExtractData{
					Tag:     "13",
					Content: "Old Content",
					Data:    "1311Old Content",
				},
				content: "New Content",
			},
			want: entities.ExtractData{
				Tag:     "13",
				Content: "New Content",
				Data:    "1311New Content",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Data{}
			got := uc.ModifyContent(test.args.extractData, test.args.content)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ModifyContent()", test.want, got)
			}
		})
	}
}
