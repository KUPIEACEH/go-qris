package handlers

import "github.com/fyvri/go-qris/internal/domain/entities"

var (
	expectedButGotMessage             = "Expected %v = %v, but got = %v"
	expectedTypeAssertionErrorMessage = "Expected type assertion error, but got = %v"
	expectedReturnNonNil              = "Expected %v to return a non-nil %v"
	expectedStatusCode                = "Expected status code %d, got %d"
	expectedResponseToContain         = "Expected response to contain %s, but got %s"

	testNameInvalidJSON        = "Error: Invalid JSON"
	testHeaderContentType      = "Content-Type"
	testHeaderContentTypeValue = "application/json"
)

type mockQRISController struct {
	ParseFunc     func(qrisString string) (*entities.QRIS, error, *[]string)
	ToDynamicFunc func(qrisString string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error, *[]string)
	ValidateFunc  func(qrisString string) (error, *[]string)
}

func (m *mockQRISController) Parse(qrisString string) (*entities.QRIS, error, *[]string) {
	if m.ParseFunc != nil {
		return m.ParseFunc(qrisString)
	}
	return nil, nil, nil
}

func (m *mockQRISController) ToDynamic(qrisString string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error, *[]string) {
	if m.ToDynamicFunc != nil {
		return m.ToDynamicFunc(qrisString, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee)
	}
	return "", "", nil, nil
}

func (m *mockQRISController) Validate(qrisString string) (error, *[]string) {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(qrisString)
	}
	return nil, nil
}
