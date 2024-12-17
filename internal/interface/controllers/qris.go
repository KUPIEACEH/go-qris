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
	Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) (string, string, error, *[]string)
	IsValid(qrisString string) (error, *[]string)
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

func (c *QRIS) Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32) (string, string, error, *[]string) {
	qrisString = c.inputUtil.Sanitize(qrisString)
	qris, err, errs := c.qrisUsecase.Parse(qrisString)
	if err != nil {
		return "", "", err, errs
	}

	merchantCityValue = c.inputUtil.Sanitize(merchantCityValue)
	merchantPostalCodeValue = c.inputUtil.Sanitize(merchantPostalCodeValue)
	paymentFeeCategoryValue = strings.ToUpper(c.inputUtil.Sanitize(paymentFeeCategoryValue))
	qris = c.qrisUsecase.Modify(qris, merchantCityValue, merchantPostalCodeValue, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue)
	qrisString = c.qrisUsecase.ToString(qris)

	qrCode, err := c.qrCodeUtil.StringToImageBase64(qrisString, c.qrCodeSize)
	if err != nil {
		return qrisString, "", err, nil
	}

	return qrisString, qrCode, nil, nil
}

func (c *QRIS) IsValid(qrisString string) (error, *[]string) {
	qrisString = c.inputUtil.Sanitize(qrisString)
	qris, err, errs := c.qrisUsecase.Parse(qrisString)
	if err != nil {
		return err, errs
	}

	isValid := c.qrisUsecase.IsValid(qris)
	if !isValid {
		return fmt.Errorf("invalid CRC16-CCITT code"), nil
	}

	return nil, nil
}
