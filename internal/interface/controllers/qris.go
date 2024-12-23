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
	Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) (string, string, error, *[]string)
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

func (c *QRIS) Convert(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) (string, string, error, *[]string) {
	errs := &[]string{}
	merchantCityValue = c.inputUtil.Sanitize(merchantCityValue)
	if len(merchantCityValue) > 15 {
		*errs = append(*errs, "merchant city exceeds 15 characters")
	}
	merchantPostalCodeValue = c.inputUtil.Sanitize(merchantPostalCodeValue)
	if len(merchantPostalCodeValue) > 10 {
		*errs = append(*errs, "merchant postal code exceeds 10 characters")
	}
	terminalLabelValue = c.inputUtil.Sanitize(terminalLabelValue)
	if len(terminalLabelValue) > 99 {
		*errs = append(*errs, "terminal label exceeds 99 characters")
	}
	if len(*errs) > 0 {
		return "", "", fmt.Errorf("input length exceeds the maximum permitted characters"), errs
	}

	qrisString = c.inputUtil.Sanitize(qrisString)
	qris, err, errs := c.qrisUsecase.Parse(qrisString)
	if err != nil {
		return "", "", err, errs
	}

	paymentFeeCategoryValue = strings.ToUpper(c.inputUtil.Sanitize(paymentFeeCategoryValue))
	qris = c.qrisUsecase.Modify(qris, merchantCityValue, merchantPostalCodeValue, paymentAmountValue, paymentFeeCategoryValue, paymentFeeValue, terminalLabelValue)
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
