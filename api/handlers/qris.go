package handlers

import (
	"net/http"

	"github.com/fyvri/go-qris/internal/interface/controllers"

	"github.com/gin-gonic/gin"
)

type QRIS struct {
	qrisController controllers.QRISInterface
}

type QRISInterface interface {
	Parse(c *gin.Context)
	ToDynamic(c *gin.Context)
	Validate(c *gin.Context)
}

type ParseRequest struct {
	QRString string `json:"qr_string"`
}

type ConverterRequest struct {
	QRString           string `json:"qr_string"`
	MerchantCity       string `json:"merchant_city"`
	MerchantPostalCode string `json:"merchant_postal_code"`
	PaymentAmount      uint32 `json:"payment_amount"`
	PaymentFeeCategory string `json:"payment_fee_category"`
	PaymentFee         uint32 `json:"payment_fee"`
}

func NewQRIS(qrisController controllers.QRISInterface) QRISInterface {
	return &QRIS{
		qrisController: qrisController,
	}
}

func (h *QRIS) Parse(c *gin.Context) {
	var req ParseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Errors:  nil,
			Data:    nil,
		})
		return
	}

	data, err, errs := h.qrisController.Parse(req.QRString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Errors:  errs,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "QRIS parsed successfully",
		Errors:  nil,
		Data:    data,
	})
}

func (h *QRIS) ToDynamic(c *gin.Context) {
	var req ConverterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Errors:  nil,
			Data:    nil,
		})
		return
	}

	qrString, qrCode, err, errs := h.qrisController.ToDynamic(req.QRString, req.MerchantCity, req.MerchantPostalCode, req.PaymentAmount, req.PaymentFeeCategory, req.PaymentFee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Errors:  errs,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Dynamic QRIS converted successfully",
		Errors:  nil,
		Data: struct {
			QRString string `json:"qr_string"`
			QRCode   string `json:"qr_code"`
		}{
			QRString: qrString,
			QRCode:   qrCode,
		},
	})
}

func (h *QRIS) Validate(c *gin.Context) {
	var req ParseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Errors:  nil,
			Data:    nil,
		})
		return
	}

	err, errs := h.qrisController.Validate(req.QRString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Errors:  errs,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "CRC16-CCITT code is valid",
		Errors:  nil,
		Data:    nil,
	})
}
