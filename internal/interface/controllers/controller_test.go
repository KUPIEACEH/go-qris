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

	testQRISString        = "QR String"
	testQRISDynamicString = "QRIS Dynamic String"
	testQRCodeSize        = 125
)

type mockQRISUsecase struct {
	ParseFunc           func(qrString string) (*entities.QRIS, error, *[]string)
	ToDynamicFunc       func(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic
	DynamicToStringFunc func(qrisDynamic *entities.QRISDynamic) string
	ValidateFunc        func(qris *entities.QRIS) bool
}

func (m *mockQRISUsecase) Parse(qrString string) (*entities.QRIS, error, *[]string) {
	if m.ParseFunc != nil {
		return m.ParseFunc(qrString)
	}
	return nil, nil, nil
}

func (m *mockQRISUsecase) ToDynamic(qris *entities.QRIS, merchantCity string, merchantPostalCode string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRISDynamic {
	if m.ToDynamicFunc != nil {
		return m.ToDynamicFunc(qris, merchantCity, merchantPostalCode, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)
	}
	return nil
}

func (m *mockQRISUsecase) DynamicToString(qrisDynamic *entities.QRISDynamic) string {
	if m.DynamicToStringFunc != nil {
		return m.DynamicToStringFunc(qrisDynamic)
	}
	return ""
}

func (m *mockQRISUsecase) Validate(qris *entities.QRIS) bool {
	if m.ValidateFunc != nil {
		return m.ValidateFunc(qris)
	}
	return false
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
