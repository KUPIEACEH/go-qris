package usecases

import (
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
			name: "Success: No Field",
			fields: Switching{
				dataUsecase: &Data{},
				switchingDetailTags: &SwitchingDetailTags{
					Site:     testSwitchingDetailSiteTag,
					NMID:     testSwitchingDetailNMIDTag,
					Category: testSwitchingDetailCategoryTag,
				},
			},
			want: &Switching{
				dataUsecase: &Data{},
				switchingDetailTags: &SwitchingDetailTags{
					Site:     testSwitchingDetailSiteTag,
					NMID:     testSwitchingDetailNMIDTag,
					Category: testSwitchingDetailCategoryTag,
				},
			},
		},
		{
			name: "Success: With Field",
			fields: Switching{
				dataUsecase: &Data{},
				switchingDetailTags: &SwitchingDetailTags{
					Site:     testSwitchingDetailSiteTag,
					NMID:     testSwitchingDetailNMIDTag,
					Category: testSwitchingDetailCategoryTag,
				},
			},
			want: &Switching{
				dataUsecase: &Data{},
				switchingDetailTags: &SwitchingDetailTags{
					Site:     testSwitchingDetailSiteTag,
					NMID:     testSwitchingDetailNMIDTag,
					Category: testSwitchingDetailCategoryTag,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewSwitching(test.fields.dataUsecase, test.fields.switchingDetailTags)

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

func TestSwitchingParse(t *testing.T) {
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
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
				switchingDetailTags: &SwitchingDetailTags{
					Site:     testSwitchingDetailSiteTag,
					NMID:     testSwitchingDetailNMIDTag,
					Category: testSwitchingDetailCategoryTag,
				},
			},
			args: args{
				content: testQRIS.Switching.Content,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid format code"),
		},
		{
			name: "Success",
			fields: Switching{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						switch codeString[:2] {
						case testQRIS.Switching.Detail.Site.Tag:
							return &testQRIS.Switching.Detail.Site, nil
						case testQRIS.Switching.Detail.NMID.Tag:
							return &testQRIS.Switching.Detail.NMID, nil
						case testQRIS.Switching.Detail.Category.Tag:
							return &testQRIS.Switching.Detail.Category, nil
						default:
							return nil, nil
						}
					},
				},
				switchingDetailTags: &SwitchingDetailTags{
					Site:     testSwitchingDetailSiteTag,
					NMID:     testSwitchingDetailNMIDTag,
					Category: testSwitchingDetailCategoryTag,
				},
			},
			args: args{
				content: testQRIS.Switching.Content,
			},
			want:      &testQRIS.Switching.Detail,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Switching{
				dataUsecase:         test.fields.dataUsecase,
				switchingDetailTags: test.fields.switchingDetailTags,
			}

			got, err := uc.Parse(test.args.content)
			if err != nil && err.Error() != test.wantError.Error() {
				t.Errorf(expectedErrorButGotMessage, "Parse()", test.wantError, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "Parse()", test.want, got)
			}
		})
	}
}
