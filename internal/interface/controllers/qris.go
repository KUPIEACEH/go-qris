package controllers

import (
	"fmt"
	"strings"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/usecases"
	"github.com/fyvri/go-qris/pkg/utils"
)

type QRIS struct {
	inputUtil   utils.InputInterface
	qrCodeUtil  utils.QRCodeInterface
	qrisUsecase usecases.QRISInterface
	qrCodeSize  int
}

type QRISInterface interface {
	Parse(qrisString string) (*entities.QRIS, error, *[]string)
	ToDynamic(qrisString string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error, *[]string)
	Validate(qrisString string) (error, *[]string)
}

func NewQRIS(inputUtil utils.InputInterface, qrCodeUtil utils.QRCodeInterface, qrisUsecase usecases.QRISInterface, qrCodeSize int) QRISInterface {
	return &QRIS{
		inputUtil:   inputUtil,
		qrisUsecase: qrisUsecase,
		qrCodeUtil:  qrCodeUtil,
		qrCodeSize:  qrCodeSize,
	}
}

func (c *QRIS) Parse(qrisString string) (*entities.QRIS, error, *[]string) {
	qrisString = c.inputUtil.Sanitize(qrisString)

	return c.qrisUsecase.Parse(qrisString)
}

func (c *QRIS) ToDynamic(qrisString string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error, *[]string) {
	qrisString = c.inputUtil.Sanitize(qrisString)
	qris, err, errs := c.qrisUsecase.Parse(qrisString)
	if err != nil {
		return "", "", err, errs
	}

	merchantCity = c.inputUtil.Sanitize(merchantCity)
	merchantPostalCode = c.inputUtil.Sanitize(merchantPostalCode)
	paymentFeeCategory = strings.ToUpper(c.inputUtil.Sanitize(paymentFeeCategory))
	qrDynamic := c.qrisUsecase.ToDynamic(qris, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee)
	qrDynamicString := c.qrisUsecase.DynamicToString(qrDynamic)

	qrCode, err := c.qrCodeUtil.StringToImageBase64(qrDynamicString, c.qrCodeSize)
	if err != nil {
		return qrDynamicString, "", err, nil
	}

	return qrDynamicString, qrCode, nil, nil
}

func (c *QRIS) Validate(qrisString string) (error, *[]string) {
	qrisString = c.inputUtil.Sanitize(qrisString)
	qris, err, errs := c.qrisUsecase.Parse(qrisString)
	if err != nil {
		return err, errs
	}

	isValid := c.qrisUsecase.Validate(qris)
	if !isValid {
		return fmt.Errorf("invalid CRC16-CCITT code"), nil
	}

	return nil, nil
}
