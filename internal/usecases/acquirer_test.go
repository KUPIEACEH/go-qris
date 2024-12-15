package usecases

import (
	"errors"
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
			name: "Success",
			fields: Acquirer{
				dataUsecase:   &Data{},
				siteTag:       "",
				mpanTag:       "",
				terminalIDTag: "",
				categoryTag:   "",
			},
			want: &Acquirer{
				dataUsecase:   &Data{},
				siteTag:       "",
				mpanTag:       "",
				terminalIDTag: "",
				categoryTag:   "",
			},
		},
		{
			name: "Success With Fields",
			fields: Acquirer{
				dataUsecase:   &Data{},
				siteTag:       testAcquirerDetailSiteTag,
				mpanTag:       testAcquirerDetailMPANTag,
				terminalIDTag: testAcquirerDetailTerminalIDTag,
				categoryTag:   testAcquirerDetailCategoryTag,
			},
			want: &Acquirer{
				dataUsecase:   &Data{},
				siteTag:       testAcquirerDetailSiteTag,
				mpanTag:       testAcquirerDetailMPANTag,
				terminalIDTag: testAcquirerDetailTerminalIDTag,
				categoryTag:   testAcquirerDetailCategoryTag,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewAcquirer(test.fields.dataUsecase, test.fields.siteTag, test.fields.mpanTag, test.fields.terminalIDTag, test.fields.categoryTag)

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

func TestAcquirerExtract(t *testing.T) {
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
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
				siteTag:       testAcquirerDetailSiteTag,
				mpanTag:       testAcquirerDetailMPANTag,
				terminalIDTag: testAcquirerDetailTerminalIDTag,
				categoryTag:   testAcquirerDetailCategoryTag,
			},
			args: args{
				content: testQRISStatic.Acquirer.Content,
			},
			want:      nil,
			wantError: errors.New("invalid format code"),
		},
		{
			name: "Success",
			fields: Acquirer{
				dataUsecase: &mockDataUsecase{
					ExtractFunc: func(codeString string) (*entities.ExtractData, error) {
						switch codeString[:2] {
						case testQRISStatic.Acquirer.Detail.Site.Tag:
							return &testQRISStatic.Acquirer.Detail.Site, nil
						case testQRISStatic.Acquirer.Detail.MPAN.Tag:
							return &testQRISStatic.Acquirer.Detail.MPAN, nil
						case testQRISStatic.Acquirer.Detail.TerminalID.Tag:
							return &testQRISStatic.Acquirer.Detail.TerminalID, nil
						case testQRISStatic.Acquirer.Detail.Category.Tag:
							return &testQRISStatic.Acquirer.Detail.Category, nil
						default:
							return nil, nil
						}
					},
				},
				siteTag:       testAcquirerDetailSiteTag,
				mpanTag:       testAcquirerDetailMPANTag,
				terminalIDTag: testAcquirerDetailTerminalIDTag,
				categoryTag:   testAcquirerDetailCategoryTag,
			},
			args: args{
				content: testQRISStatic.Acquirer.Content,
			},
			want:      &testQRISStatic.Acquirer.Detail,
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &Acquirer{
				dataUsecase:   test.fields.dataUsecase,
				siteTag:       test.fields.siteTag,
				mpanTag:       test.fields.mpanTag,
				terminalIDTag: test.fields.terminalIDTag,
				categoryTag:   test.fields.categoryTag,
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
