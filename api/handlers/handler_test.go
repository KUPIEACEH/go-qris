package handlers

import "github.com/fyvri/go-qris/internal/domain/entities"

var (
	expectedButGotMessage             = "Expected %v = %v, but got = %v"
	expectedTypeAssertionErrorMessage = "Expected type assertion error, but got = %v"
	expectedReturnNonNil              = "Expected %v to return a non-nil %v"
	expectedStatusCode                = "Expected status code %d, got %d"
	expectedResponseToContain         = "Expected response to contain %s, but got %s"
)

type mockQRISController struct {
	ExtractStaticFunc   func(qrisStaticString string) (*entities.QRISStatic, error)
	StaticToDynamicFunc func(qrisStaticString string, merchantName string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error)
}

func (m *mockQRISController) ExtractStatic(qrisStaticString string) (*entities.QRISStatic, error) {
	if m.ExtractStaticFunc != nil {
		return m.ExtractStaticFunc(qrisStaticString)
	}
	return nil, nil
}

func (m *mockQRISController) StaticToDynamic(qrisStaticString string, merchantName string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error) {
	if m.StaticToDynamicFunc != nil {
		return m.StaticToDynamicFunc(qrisStaticString, merchantName, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee)
	}
	return "", "", nil
}
