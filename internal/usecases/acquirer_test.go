package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewAcquirer(t *testing.T) {
	tests := []struct {
		name   string
		fields Acquirer
		want   AcquirerInterface
	}{
		{
			name: "Success: No Field",
			fields: Acquirer{
				dataUsecase:        &Data{},
				acquirerDetailTags: &AcquirerDetailTags{},
			},
			want: &Acquirer{
				dataUsecase:        &Data{},
				acquirerDetailTags: &AcquirerDetailTags{},
			},
		},
		{
			name: "Success: With Field",
			fields: Acquirer{
				dataUsecase: &Data{},
				acquirerDetailTags: &AcquirerDetailTags{
					Site:       testAcquirerDetailSiteTag,
					MPAN:       testAcquirerDetailMPANTag,
					TerminalID: testAcquirerDetailTerminalIDTag,
					Category:   testAcquirerDetailCategoryTag,
				},
			},
			want: &Acquirer{
				dataUsecase: &Data{},
				acquirerDetailTags: &AcquirerDetailTags{
					Site:       testAcquirerDetailSiteTag,
					MPAN:       testAcquirerDetailMPANTag,
					TerminalID: testAcquirerDetailTerminalIDTag,
					Category:   testAcquirerDetailCategoryTag,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewAcquirer(test.fields.dataUsecase, test.fields.acquirerDetailTags)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewAcquirer", "AcquirerInterface")
			}

			got, ok := uc.(*Acquirer)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*Acquirer")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*Acquirer", test.want, got)
			}
		})
	}
}

func TestAcquirerParse(t *testing.T) {
	type args struct {
		content string
	}

	tests := []struct {
		name      string
		fields    Acquirer
		args      args
		want      *entities.AcquirerDetail
		wantError error
	}{
		{
			name: "Error: Parse",
			fields: Acquirer{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
				acquirerDetailTags: &AcquirerDetailTags{
					Site:       testAcquirerDetailSiteTag,
					MPAN:       testAcquirerDetailMPANTag,
					TerminalID: testAcquirerDetailTerminalIDTag,
					Category:   testAcquirerDetailCategoryTag,
				},
			},
			args: args{
				content: testQRIS.Acquirer.Content,
			},
			want:      nil,
			wantError: fmt.Errorf("invalid format code"),
		},
		{
			name: "Success",
			fields: Acquirer{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						switch codeString[:2] {
						case testQRIS.Acquirer.Detail.Site.Tag:
							return &testQRIS.Acquirer.Detail.Site, nil
						case testQRIS.Acquirer.Detail.MPAN.Tag:
							return &testQRIS.Acquirer.Detail.MPAN, nil
						case testQRIS.Acquirer.Detail.TerminalID.Tag:
							return &testQRIS.Acquirer.Detail.TerminalID, nil
						case testQRIS.Acquirer.Detail.Category.Tag:
							return &testQRIS.Acquirer.Detail.Category, nil
						default:
							return nil, nil
						}
					},
				},
				acquirerDetailTags: &AcquirerDetailTags{
					Site:       testAcquirerDetailSiteTag,
					MPAN:       testAcquirerDetailMPANTag,
					TerminalID: testAcquirerDetailTerminalIDTag,
					Category:   testAcquirerDetailCategoryTag,
				},
			},
			args: args{
				content: testQRIS.Acquirer.Content,
			},
			want:      &testQRIS.Acquirer.Detail,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Acquirer{
				dataUsecase: test.fields.dataUsecase,
				acquirerDetailTags: &AcquirerDetailTags{
					Site:       testAcquirerDetailSiteTag,
					MPAN:       testAcquirerDetailMPANTag,
					TerminalID: testAcquirerDetailTerminalIDTag,
					Category:   testAcquirerDetailCategoryTag,
				},
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
