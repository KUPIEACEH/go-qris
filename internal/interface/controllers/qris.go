package controllers

import (
	"fmt"
	"strings"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/utils"
)

type QRIS struct {
	qrisUsecase usecases.QRISInterface
	qrCodeUtil  utils.QRCodeInterface
	qrCodeSize  int
}

type QRISInterface interface {
	ExtractStatic(qrisStaticString string) (*entities.QRISStatic, error)
	StaticToDynamic(qrisStaticString string, merchantName string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error)
}

func NewQRIS(qrisUsecase usecases.QRISInterface, qrCodeUtil utils.QRCodeInterface, qrCodeSize int) QRISInterface {
	return &QRIS{
		qrisUsecase: qrisUsecase,
		qrCodeUtil:  qrCodeUtil,
		qrCodeSize:  qrCodeSize,
	}
}

func sanitizeInput(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return strings.TrimSpace(input)
}

func (c *QRIS) ExtractStatic(qrisStaticString string) (*entities.QRISStatic, error) {
	qrisStaticString = sanitizeInput(qrisStaticString)
	if qrisStaticString == "" {
		return nil, fmt.Errorf("QRIS not found")
	}

	return c.qrisUsecase.ExtractStatic(qrisStaticString)
}

func (c *QRIS) StaticToDynamic(qrisStaticString string, merchantName string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error) {
	qrisStaticString = sanitizeInput(qrisStaticString)
	if qrisStaticString == "" {
		return "", "", fmt.Errorf("QRIS not found")
	}

	merchantCity = sanitizeInput(merchantCity)
	merchantPostalCode = sanitizeInput(merchantPostalCode)
	paymentFeeCategory = strings.ToUpper(sanitizeInput(paymentFeeCategory))

	qrStatic, err := c.qrisUsecase.ExtractStatic(qrisStaticString)
	if err != nil {
		return "", "", err
	}

	qrDynamic := c.qrisUsecase.StaticToDynamic(qrStatic, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee)
	qrDynamicString := c.qrisUsecase.DynamicToDynamicString(qrDynamic)

	qrCode, err := c.qrCodeUtil.StringToImageBase64(qrDynamicString, c.qrCodeSize)
	if err != nil {
		return qrDynamicString, "", err
	}

	return qrDynamicString, qrCode, nil
}
