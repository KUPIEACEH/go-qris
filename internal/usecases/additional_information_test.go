package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

func TestNewAdditionalInformation(t *testing.T) {
	tests := []struct {
		name   string
		fields AdditionalInformation
		want   AdditionalInformationInterface
	}{
		{
			name: "Success: No Field",
			fields: AdditionalInformation{
				dataUsecase:                     &Data{},
				additionalInformationDetailTags: &AdditionalInformationDetailTags{},
			},
			want: &AdditionalInformation{
				dataUsecase:                     &Data{},
				additionalInformationDetailTags: &AdditionalInformationDetailTags{},
			},
		},
		{
			name: "Success: With Field",
			fields: AdditionalInformation{
				dataUsecase: &Data{},
				additionalInformationDetailTags: &AdditionalInformationDetailTags{
					BillNumber:                    testAdditionalInformationDetailBillNumberTag,
					MobileNumber:                  testAdditionalInformationDetailMobileNumberTag,
					StoreLabel:                    testAdditionalInformationDetailStoreLabelTag,
					LoyaltyNumber:                 testAdditionalInformationDetailLoyaltyNumberTag,
					ReferenceLabel:                testAdditionalInformationDetailReferenceLabelTag,
					CustomerLabel:                 testAdditionalInformationDetailCustomerLabelTag,
					TerminalLabel:                 testAdditionalInformationDetailTerminalLabelTag,
					PurposeOfTransaction:          testAdditionalInformationDetailPurposeOfTransactionTag,
					AdditionalConsumerDataRequest: testAdditionalInformationDetailAdditionalConsumerDataRequestTag,
					MerchantTaxID:                 testAdditionalInformationDetailMerchantTaxIDTag,
					MerchantChannel:               testAdditionalInformationDetailMerchantChannelTag,
					RFUStart:                      testAdditionalInformationDetailRFUTagStart,
					RFUEnd:                        testAdditionalInformationDetailRFUTagEnd,
					PaymentSystemSpecificStart:    testAdditionalInformationDetailPaymentSystemSpecificTagStart,
					PaymentSystemSpecificEnd:      testAdditionalInformationDetailPaymentSystemSpecificTagEnd,
				},
			},
			want: &AdditionalInformation{
				dataUsecase: &Data{},
				additionalInformationDetailTags: &AdditionalInformationDetailTags{
					BillNumber:                    testAdditionalInformationDetailBillNumberTag,
					MobileNumber:                  testAdditionalInformationDetailMobileNumberTag,
					StoreLabel:                    testAdditionalInformationDetailStoreLabelTag,
					LoyaltyNumber:                 testAdditionalInformationDetailLoyaltyNumberTag,
					ReferenceLabel:                testAdditionalInformationDetailReferenceLabelTag,
					CustomerLabel:                 testAdditionalInformationDetailCustomerLabelTag,
					TerminalLabel:                 testAdditionalInformationDetailTerminalLabelTag,
					PurposeOfTransaction:          testAdditionalInformationDetailPurposeOfTransactionTag,
					AdditionalConsumerDataRequest: testAdditionalInformationDetailAdditionalConsumerDataRequestTag,
					MerchantTaxID:                 testAdditionalInformationDetailMerchantTaxIDTag,
					MerchantChannel:               testAdditionalInformationDetailMerchantChannelTag,
					RFUStart:                      testAdditionalInformationDetailRFUTagStart,
					RFUEnd:                        testAdditionalInformationDetailRFUTagEnd,
					PaymentSystemSpecificStart:    testAdditionalInformationDetailPaymentSystemSpecificTagStart,
					PaymentSystemSpecificEnd:      testAdditionalInformationDetailPaymentSystemSpecificTagEnd,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewAdditionalInformation(test.fields.dataUsecase, test.fields.additionalInformationDetailTags)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewAdditionalInformation", "AdditionalInformationInterface")
			}

			got, ok := uc.(*AdditionalInformation)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*AdditionalInformation")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*AdditionalInformation", test.want, got)
			}
		})
	}
}

func TestAdditionalInformationParse(t *testing.T) {
	type args struct {
		content string
	}

	tests := []struct {
		name      string
		fields    AdditionalInformation
		args      args
		want      *entities.AdditionalInformationDetail
		wantError error
	}{
		{
			name: "Error: Parse",
			fields: AdditionalInformation{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						return nil, fmt.Errorf("invalid format code")
					},
				},
			},
			args: args{
				content: "0101A0201B0301C0401D0501E0601F0701G0801H0901I1001J1101K1701Q9901Z",
			},
			want:      nil,
			wantError: fmt.Errorf("invalid format code"),
		},
		{
			name: "Success",
			fields: AdditionalInformation{
				dataUsecase: &mockDataUsecase{
					ParseFunc: func(codeString string) (*entities.Data, error) {
						switch codeString[:2] {
						case testAdditionalInformationDetailBillNumberTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailBillNumberTag,
								Content: "A",
								Data:    testAdditionalInformationDetailBillNumberTag + "01A",
							}, nil
						case testAdditionalInformationDetailMobileNumberTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailMobileNumberTag,
								Content: "B",
								Data:    testAdditionalInformationDetailMobileNumberTag + "01B",
							}, nil
						case testAdditionalInformationDetailStoreLabelTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailStoreLabelTag,
								Content: "C",
								Data:    testAdditionalInformationDetailStoreLabelTag + "01C",
							}, nil
						case testAdditionalInformationDetailLoyaltyNumberTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailLoyaltyNumberTag,
								Content: "D",
								Data:    testAdditionalInformationDetailLoyaltyNumberTag + "01D",
							}, nil
						case testAdditionalInformationDetailReferenceLabelTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailReferenceLabelTag,
								Content: "E",
								Data:    testAdditionalInformationDetailReferenceLabelTag + "01E",
							}, nil
						case testAdditionalInformationDetailCustomerLabelTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailCustomerLabelTag,
								Content: "F",
								Data:    testAdditionalInformationDetailCustomerLabelTag + "01F",
							}, nil
						case testAdditionalInformationDetailTerminalLabelTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailTerminalLabelTag,
								Content: "G",
								Data:    testAdditionalInformationDetailTerminalLabelTag + "01G",
							}, nil
						case testAdditionalInformationDetailPurposeOfTransactionTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailPurposeOfTransactionTag,
								Content: "H",
								Data:    testAdditionalInformationDetailPurposeOfTransactionTag + "01H",
							}, nil
						case testAdditionalInformationDetailAdditionalConsumerDataRequestTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailAdditionalConsumerDataRequestTag,
								Content: "I",
								Data:    testAdditionalInformationDetailAdditionalConsumerDataRequestTag + "01I",
							}, nil
						case testAdditionalInformationDetailMerchantTaxIDTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailMerchantTaxIDTag,
								Content: "J",
								Data:    testAdditionalInformationDetailMerchantTaxIDTag + "01J",
							}, nil
						case testAdditionalInformationDetailMerchantChannelTag:
							return &entities.Data{
								Tag:     testAdditionalInformationDetailMerchantChannelTag,
								Content: "K",
								Data:    testAdditionalInformationDetailMerchantChannelTag + "01K",
							}, nil
						case "17":
							return &entities.Data{
								Tag:     "17",
								Content: "Q",
								Data:    "17" + "01Q",
							}, nil
						case "99":
							return &entities.Data{
								Tag:     "99",
								Content: "Z",
								Data:    "99" + "01Z",
							}, nil
						default:
							return nil, nil
						}
					},
				},
			},
			args: args{
				content: "0101A0201B0301C0401D0501E0601F0701G0801H0901I1001J1101K1701Q9901Z",
			},
			want: &entities.AdditionalInformationDetail{
				BillNumber: entities.Data{
					Tag:     testAdditionalInformationDetailBillNumberTag,
					Content: "A",
					Data:    testAdditionalInformationDetailBillNumberTag + "01A",
				},
				MobileNumber: entities.Data{
					Tag:     testAdditionalInformationDetailMobileNumberTag,
					Content: "B",
					Data:    testAdditionalInformationDetailMobileNumberTag + "01B",
				},
				StoreLabel: entities.Data{
					Tag:     testAdditionalInformationDetailStoreLabelTag,
					Content: "C",
					Data:    testAdditionalInformationDetailStoreLabelTag + "01C",
				},
				LoyaltyNumber: entities.Data{
					Tag:     testAdditionalInformationDetailLoyaltyNumberTag,
					Content: "D",
					Data:    testAdditionalInformationDetailLoyaltyNumberTag + "01D",
				},
				ReferenceLabel: entities.Data{
					Tag:     testAdditionalInformationDetailReferenceLabelTag,
					Content: "E",
					Data:    testAdditionalInformationDetailReferenceLabelTag + "01E",
				},
				CustomerLabel: entities.Data{
					Tag:     testAdditionalInformationDetailCustomerLabelTag,
					Content: "F",
					Data:    testAdditionalInformationDetailCustomerLabelTag + "01F",
				},
				TerminalLabel: entities.Data{
					Tag:     testAdditionalInformationDetailTerminalLabelTag,
					Content: "G",
					Data:    testAdditionalInformationDetailTerminalLabelTag + "01G",
				},
				PurposeOfTransaction: entities.Data{
					Tag:     testAdditionalInformationDetailPurposeOfTransactionTag,
					Content: "H",
					Data:    testAdditionalInformationDetailPurposeOfTransactionTag + "01H",
				},
				AdditionalConsumerDataRequest: entities.Data{
					Tag:     testAdditionalInformationDetailAdditionalConsumerDataRequestTag,
					Content: "I",
					Data:    testAdditionalInformationDetailAdditionalConsumerDataRequestTag + "01I",
				},
				MerchantTaxID: entities.Data{
					Tag:     testAdditionalInformationDetailMerchantTaxIDTag,
					Content: "J",
					Data:    testAdditionalInformationDetailMerchantTaxIDTag + "01J",
				},
				MerchantChannel: entities.Data{
					Tag:     testAdditionalInformationDetailMerchantChannelTag,
					Content: "K",
					Data:    testAdditionalInformationDetailMerchantChannelTag + "01K",
				},
				RFU: entities.Data{
					Tag:     "17",
					Content: "Q",
					Data:    "17" + "01Q",
				},
				PaymentSystemSpecific: entities.Data{
					Tag:     "99",
					Content: "Z",
					Data:    "99" + "01Z",
				},
			},
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &AdditionalInformation{
				dataUsecase: test.fields.dataUsecase,
				additionalInformationDetailTags: &AdditionalInformationDetailTags{
					BillNumber:                    testAdditionalInformationDetailBillNumberTag,
					MobileNumber:                  testAdditionalInformationDetailMobileNumberTag,
					StoreLabel:                    testAdditionalInformationDetailStoreLabelTag,
					LoyaltyNumber:                 testAdditionalInformationDetailLoyaltyNumberTag,
					ReferenceLabel:                testAdditionalInformationDetailReferenceLabelTag,
					CustomerLabel:                 testAdditionalInformationDetailCustomerLabelTag,
					TerminalLabel:                 testAdditionalInformationDetailTerminalLabelTag,
					PurposeOfTransaction:          testAdditionalInformationDetailPurposeOfTransactionTag,
					AdditionalConsumerDataRequest: testAdditionalInformationDetailAdditionalConsumerDataRequestTag,
					MerchantTaxID:                 testAdditionalInformationDetailMerchantTaxIDTag,
					MerchantChannel:               testAdditionalInformationDetailMerchantChannelTag,
					RFUStart:                      testAdditionalInformationDetailRFUTagStart,
					RFUEnd:                        testAdditionalInformationDetailRFUTagEnd,
					PaymentSystemSpecificStart:    testAdditionalInformationDetailPaymentSystemSpecificTagStart,
					PaymentSystemSpecificEnd:      testAdditionalInformationDetailPaymentSystemSpecificTagEnd,
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

func TestAdditionalInformationToString(t *testing.T) {
	type args struct {
		additionalInformationDetail *entities.AdditionalInformationDetail
	}

	testAdditionalInformationDetail := &entities.AdditionalInformationDetail{
		BillNumber: entities.Data{
			Tag:     testAdditionalInformationDetailBillNumberTag,
			Content: "A",
			Data:    testAdditionalInformationDetailBillNumberTag + "01A",
		},
		MobileNumber: entities.Data{
			Tag:     testAdditionalInformationDetailMobileNumberTag,
			Content: "B",
			Data:    testAdditionalInformationDetailMobileNumberTag + "01B",
		},
		StoreLabel: entities.Data{
			Tag:     testAdditionalInformationDetailStoreLabelTag,
			Content: "C",
			Data:    testAdditionalInformationDetailStoreLabelTag + "01C",
		},
		LoyaltyNumber: entities.Data{
			Tag:     testAdditionalInformationDetailLoyaltyNumberTag,
			Content: "D",
			Data:    testAdditionalInformationDetailLoyaltyNumberTag + "01D",
		},
		ReferenceLabel: entities.Data{
			Tag:     testAdditionalInformationDetailReferenceLabelTag,
			Content: "E",
			Data:    testAdditionalInformationDetailReferenceLabelTag + "01E",
		},
		CustomerLabel: entities.Data{
			Tag:     testAdditionalInformationDetailCustomerLabelTag,
			Content: "F",
			Data:    testAdditionalInformationDetailCustomerLabelTag + "01F",
		},
		TerminalLabel: entities.Data{
			Tag:     testAdditionalInformationDetailTerminalLabelTag,
			Content: "G",
			Data:    testAdditionalInformationDetailTerminalLabelTag + "01G",
		},
		PurposeOfTransaction: entities.Data{
			Tag:     testAdditionalInformationDetailPurposeOfTransactionTag,
			Content: "H",
			Data:    testAdditionalInformationDetailPurposeOfTransactionTag + "01H",
		},
		AdditionalConsumerDataRequest: entities.Data{
			Tag:     testAdditionalInformationDetailAdditionalConsumerDataRequestTag,
			Content: "I",
			Data:    testAdditionalInformationDetailAdditionalConsumerDataRequestTag + "01I",
		},
		MerchantTaxID: entities.Data{
			Tag:     testAdditionalInformationDetailMerchantTaxIDTag,
			Content: "J",
			Data:    testAdditionalInformationDetailMerchantTaxIDTag + "01J",
		},
		MerchantChannel: entities.Data{
			Tag:     testAdditionalInformationDetailMerchantChannelTag,
			Content: "K",
			Data:    testAdditionalInformationDetailMerchantChannelTag + "01K",
		},
		RFU: entities.Data{
			Tag:     "17",
			Content: "Q",
			Data:    "17" + "01Q",
		},
		PaymentSystemSpecific: entities.Data{
			Tag:     "99",
			Content: "Z",
			Data:    "99" + "01Z",
		},
	}

	tests := []struct {
		name   string
		fields AdditionalInformation
		args   args
		want   string
	}{
		{
			name:   "Success",
			fields: AdditionalInformation{},
			args: args{
				additionalInformationDetail: testAdditionalInformationDetail,
			},
			want: testAdditionalInformationDetail.BillNumber.Data +
				testAdditionalInformationDetail.MobileNumber.Data +
				testAdditionalInformationDetail.StoreLabel.Data +
				testAdditionalInformationDetail.LoyaltyNumber.Data +
				testAdditionalInformationDetail.ReferenceLabel.Data +
				testAdditionalInformationDetail.CustomerLabel.Data +
				testAdditionalInformationDetail.TerminalLabel.Data +
				testAdditionalInformationDetail.PurposeOfTransaction.Data +
				testAdditionalInformationDetail.AdditionalConsumerDataRequest.Data +
				testAdditionalInformationDetail.MerchantTaxID.Data +
				testAdditionalInformationDetail.MerchantChannel.Data +
				testAdditionalInformationDetail.RFU.Data +
				testAdditionalInformationDetail.PaymentSystemSpecific.Data,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := &AdditionalInformation{}
			got := uc.ToString(test.args.additionalInformationDetail)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(expectedButGotMessage, "ToString()", test.want, got)
			}
		})
	}
}
