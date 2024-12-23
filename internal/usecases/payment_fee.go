package usecases

import (
	"fmt"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type PaymentFee struct {
	qrisTags                       *QRISTags
	qrisPaymentFeeCategoryContents *QRISPaymentFeeCategoryContents
}

type PaymentFeeInterface interface {
	Modify(qris *entities.QRIS, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS
}

func NewPaymentFee(qrisTags *QRISTags, qrisPaymentFeeCategoryContents *QRISPaymentFeeCategoryContents) PaymentFeeInterface {
	return &PaymentFee{
		qrisTags:                       qrisTags,
		qrisPaymentFeeCategoryContents: qrisPaymentFeeCategoryContents,
	}
}

func (uc *PaymentFee) Modify(qris *entities.QRIS, paymentFeeCategoryValue string, paymentFeeValue uint32) *entities.QRIS {
	paymentFeeCategoryTag := ""
	paymentFeeCategoryContent := ""
	paymentFeeCategoryContentLength := ""
	paymentFeeTag := ""
	if paymentFeeCategoryValue == "FIXED" {
		paymentFeeCategoryTag = uc.qrisTags.PaymentFeeCategory
		paymentFeeCategoryContent = uc.qrisPaymentFeeCategoryContents.Fixed
		paymentFeeCategoryContentLength = fmt.Sprintf("%02d", len(paymentFeeCategoryContent))
		paymentFeeTag = uc.qrisTags.PaymentFeeFixed
	} else if paymentFeeCategoryValue == "PERCENT" {
		paymentFeeCategoryTag = uc.qrisTags.PaymentFeeCategory
		paymentFeeCategoryContent = uc.qrisPaymentFeeCategoryContents.Percent
		paymentFeeCategoryContentLength = fmt.Sprintf("%02d", len(paymentFeeCategoryContent))
		paymentFeeTag = uc.qrisTags.PaymentFeePercent
	}

	qris.PaymentFeeCategory = entities.Data{
		Tag:     paymentFeeCategoryTag,
		Content: paymentFeeCategoryContent,
		Data:    paymentFeeCategoryTag + paymentFeeCategoryContentLength + paymentFeeCategoryContent,
	}
	if qris.PaymentFeeCategory.Tag != "" {
		content := fmt.Sprintf("%d", paymentFeeValue)
		qris.PaymentFee = entities.Data{
			Tag:     paymentFeeTag,
			Content: content,
			Data:    paymentFeeTag + fmt.Sprintf("%02d", len(content)) + content,
		}
	}

	return qris
}
