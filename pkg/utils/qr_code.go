package utils

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QRCode struct {
}

type QRCodeInterface interface {
	StringToImageBase64(qrString string, qrCodeSize int) (string, error)
}

func NewQRCode() QRCodeInterface {
	return &QRCode{}
}

func (u *QRCode) StringToImageBase64(qrString string, qrCodeSize int) (string, error) {
	qrCode, err := qr.Encode(qrString, qr.L, qr.Auto)
	if err != nil {
		return "", err
	}

	qrCode, err = barcode.Scale(qrCode, qrCodeSize, qrCodeSize)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode)
	if err != nil {
		return "", err
	}

	base64String := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64String, nil
}
