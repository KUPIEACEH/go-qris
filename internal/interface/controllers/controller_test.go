package controllers

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
)

var (
	expectedButGotMessage             = "Expected %v = %v, but got = %v"
	expectedErrorButGotMessage        = "Expected %v error = %v, but got = %v"
	expectedTypeAssertionErrorMessage = "Expected type assertion error, but got = %v"
	expectedReturnNonNil              = "Expected %v to return a non-nil %v"

	testNameErrorParse              = "Error: c.qrisUsecase.Parse()"
	testErrMessageInvalidFormatCode = "invalid format code"

	testQRISString         = "QR String"
	testQRISModifiedString = "QRIS Modified String"
	testQRCodeSize         = 125
)

type mockQRISUsecase struct {
	ParseFunc    func(qrString string) (*entities.QRIS, error, *[]string)
	IsValidFunc  func(qris *entities.QRIS) bool
	ModifyFunc   func(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS
	ToStringFunc func(qris *entities.QRIS) string
}

func (m *mockQRISUsecase) Parse(qrString string) (*entities.QRIS, error, *[]string) {
	if m.ParseFunc != nil {
		return m.ParseFunc(qrString)
	}
	return nil, nil, nil
}

func (m *mockQRISUsecase) IsValid(qris *entities.QRIS) bool {
	if m.IsValidFunc != nil {
		return m.IsValidFunc(qris)
	}
	return false
}

func (m *mockQRISUsecase) Modify(qris *entities.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
	if m.ModifyFunc != nil {
		return m.ModifyFunc(qris, merchantCityValue, merchantPostalCodeValue, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)
	}
	return nil
}

func (m *mockQRISUsecase) ToString(qris *entities.QRIS) string {
	if m.ToStringFunc != nil {
		return m.ToStringFunc(qris)
	}
	return ""
}

type mockQRCodeUtil struct {
	StringToImageBase64Func func(qrString string, qrCodeSize int) (string, error)
}

func (m *mockQRCodeUtil) StringToImageBase64(qrString string, qrCodeSize int) (string, error) {
	if m.StringToImageBase64Func != nil {
		return m.StringToImageBase64Func(qrString, qrCodeSize)
	}
	return "", nil
}

type mockInputUtil struct {
	SanitizeFunc func(input string) string
}

func (m *mockInputUtil) Sanitize(input string) string {
	if m.SanitizeFunc != nil {
		return m.SanitizeFunc(input)
	}
	return ""
}
