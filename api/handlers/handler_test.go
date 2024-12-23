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
	ParseFunc   func(qrisString string) (*entities.QRIS, error, *[]string)
	ConvertFunc func(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) (string, string, error, *[]string)
	IsValidFunc func(qrisString string) (error, *[]string)
}

func (m *mockQRISController) Parse(qrisString string) (*entities.QRIS, error, *[]string) {
	if m.ParseFunc != nil {
		return m.ParseFunc(qrisString)
	}
	return nil, nil, nil
}

func (m *mockQRISController) Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) (string, string, error, *[]string) {
	if m.ConvertFunc != nil {
		return m.ConvertFunc(qrisString, merchantCityValue, merchantPostalCodeValue, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue, terminalLabelValue)
	}
	return "", "", nil, nil
}

func (m *mockQRISController) IsValid(qrisString string) (error, *[]string) {
	if m.IsValidFunc != nil {
		return m.IsValidFunc(qrisString)
	}
	return nil, nil
}
