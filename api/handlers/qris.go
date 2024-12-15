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
	ExtractStatic(c *gin.Context)
	StaticToDynamic(c *gin.Context)
}

type ExtractRequest struct {
	QRString string `json:"qr_string"`
}

type ConverterRequest struct {
	QRString           string `json:"qr_string"`
	MerchantName       string `json:"merchant_name"`
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

func (h *QRIS) ExtractStatic(c *gin.Context) {
	var req ExtractRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := h.qrisController.ExtractStatic(req.QRString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "QRIS extracted successfully",
		Data:    data,
	})
}

func (h *QRIS) StaticToDynamic(c *gin.Context) {
	var req ConverterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	qrString, qrCode, err := h.qrisController.StaticToDynamic(req.QRString, req.MerchantName, req.MerchantCity, req.MerchantPostalCode, req.PaymentAmount, req.PaymentFeeCategory, req.PaymentFee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Dynamic QRIS converted successfully",
		Data: struct {
			QRString string `json:"qr_string"`
			QRCode   string `json:"qr_code"`
		}{
			QRString: qrString,
			QRCode:   qrCode,
		},
	})
}
