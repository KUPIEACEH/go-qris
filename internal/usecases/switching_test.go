package usecases

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewSwitching(t *testing.T) {
	tests := []struct {
		name   string
		fields Switching
		want   SwitchingInterface
	}{
		{
			name: "Success",
			fields: Switching{
				dataUsecase: &Data{},
				siteTag:     "",
				nmidTag:     "",
				categoryTag: "",
			},
			want: &Switching{
				dataUsecase: &Data{},
				siteTag:     "",
				nmidTag:     "",
				categoryTag: "",
			},
		},
		{
			name: "Success With Fields",
			fields: Switching{
				dataUsecase: &Data{},
				siteTag:     testSwitchingDetailSiteTag,
				nmidTag:     testSwitchingDetailNMIDTag,
				categoryTag: testSwitchingDetailCategoryTag,
			},
			want: &Switching{
				dataUsecase: &Data{},
				siteTag:     testSwitchingDetailSiteTag,
				nmidTag:     testSwitchingDetailNMIDTag,
				categoryTag: testSwitchingDetailCategoryTag,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewSwitching(test.fields.dataUsecase, test.fields.siteTag, test.fields.nmidTag, test.fields.categoryTag)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewSwitching", "SwitchingInterface")
			}

			got, ok := uc.(*Switching)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*Switching")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*Switching", test.want, got)
			}
		})
	}
}

func TestSwitchingExtract(t *testing.T) {
	type args struct {
		content string
	}

	tests := []struct {
		name      string
		fields    Switching
		args      args
		want      *entities.SwitchingDetail
		wantError error
	}{
		{
			name: "Error: Parse",
			fields: Switching{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
				siteTag:     testSwitchingDetailSiteTag,
				nmidTag:     testSwitchingDetailNMIDTag,
				categoryTag: testSwitchingDetailCategoryTag,
			},
			args: args{
				content: testQRISStatic.Switching.Content,
			},
			want:      nil,
			wantError: errors.New("invalid format code"),
		},
		{
			name: "Success",
			fields: Switching{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						switch codeString[:2] {
						case testQRISStatic.Switching.Detail.Site.Tag:
							return &testQRISStatic.Switching.Detail.Site, nil
						case testQRISStatic.Switching.Detail.NMID.Tag:
							return &testQRISStatic.Switching.Detail.NMID, nil
						case testQRISStatic.Switching.Detail.Category.Tag:
							return &testQRISStatic.Switching.Detail.Category, nil
						default:
							return nil, nil
						}
					},
				},
				siteTag:     testSwitchingDetailSiteTag,
				nmidTag:     testSwitchingDetailNMIDTag,
				categoryTag: testSwitchingDetailCategoryTag,
			},
			args: args{
				content: testQRISStatic.Switching.Content,
			},
			want:      &testQRISStatic.Switching.Detail,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Switching{
				dataUsecase: test.fields.dataUsecase,
				siteTag:     test.fields.siteTag,
				nmidTag:     test.fields.nmidTag,
				categoryTag: test.fields.categoryTag,
			}

			got, err := uc.Extract(test.args.content)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Extract()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Extract()", test.want, got)
			}
		})
	}
}
