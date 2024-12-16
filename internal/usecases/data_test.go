package usecases

import (
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
func TestDataParse(t *testing.T) {
	type args struct {
		codeString string
	}

	tests := []struct {
		name      string
		fields    Data
		args      args
		want      *entities.Data
		wantError error
	}{
		{
			name:   "Error: Content Length < 5",
			fields: Data{},
			args: args{
				codeString: "1102",
			},
			want:      nil,
			wantError: fmt.Errorf("invalid format code"),
		},
		{
			name:   "Error: Content Length Format",
			fields: Data{},
			args: args{
				codeString: testQRIS.Version.Tag + "X211",
			},
			want:      nil,
			wantError: fmt.Errorf("invalid length format for tag %s: strconv.Atoi: parsing \"X2\": invalid syntax", testQRIS.Version.Tag),
		},
		{
			name:   "Error: Content Length Not Match With Field",
			fields: Data{},
			args: args{
				codeString: testQRIS.Version.Tag + "021",
			},
			want:      nil,
			wantError: fmt.Errorf("invalid length for tag %s", testQRIS.Version.Tag),
		},
		{
			name:   "Success",
			fields: Data{},
			args: args{
				codeString: testQRIS.Version.Data,
			},
			want:      &testQRIS.Version,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := test.fields

			got, err := uc.Parse(test.args.codeString)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Parse()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Parse()", test.want, got)
			}
		})
	}
}

func TestDataModifyContent(t *testing.T) {
	type args struct {
		data    *entities.Data
		content string
	}

	tests := []struct {
		name string
		args args
		want *entities.Data
	}{
		{
			name: "Success: No Content",
			args: args{
				data:    &entities.Data{},
				content: "",
			},
			want: &entities.Data{
				Tag:     "",
				Content: "",
				Data:    "",
			},
		},
		{
			name: "Success: With Content",
			args: args{
				data: &entities.Data{
					Tag:     "13",
					Content: "Old Content",
					Data:    "1311Old Content",
				},
				content: "New Content",
			},
			want: &entities.Data{
				Tag:     "13",
				Content: "New Content",
				Data:    "1311New Content",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Data{}
			got := uc.ModifyContent(test.args.data, test.args.content)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ModifyContent()", test.want, got)
			}
		})
	}
}
