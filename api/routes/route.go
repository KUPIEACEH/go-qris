package routes

import (
	"net/http"

	"github.com/fyvri/go-qris/api/handlers"
	"github.com/fyvri/go-qris/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, ginEngine *gin.Engine) {
	publicRouter := ginEngine.Group("")

	ginEngine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, handlers.Response{
			Success: true,
			Message: "Made with love by Alvriyanto Azis",
			Data: [2]any{
				map[string]any{
					"1. Name":   "Extract Static QRIS",
					"2. Method": "POST",
					"3. Path":   "/extract-static",
					"4. Body": map[string]any{
						"qr_string": "000201010211y0ur4w3soMEsT47icQr15STriN6",
					},
					"5. Code Snippet": `curl --location 'https://api.qris.membasuh.com/extract-static' --header 'Content-Type: application/json' --data '{"qr_string": "000201010211y0ur4w3soMEsT47icQr15STriN6"}'`,
					"6. Example Response": map[string]any{
						"success": map[string]any{
							"success": true,
							"message": "QRIS extracted successfully",
							"data": map[string]any{
								"version": map[string]any{
									"tag":     "00",
									"content": "01",
									"data":    "000201",
								},
								"category": map[string]any{
									"tag":     "01",
									"content": "11",
									"data":    "010211",
								},
								"acquirer": map[string]any{
									"tag":     "26",
									"content": "0016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI",
									"data":    "26630016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI",
									"detail": map[string]any{
										"site": map[string]any{
											"tag":     "00",
											"content": "COM.MEMBASUH.WWW",
											"data":    "0016COM.MEMBASUH.WWW",
										},
										"mpan": map[string]any{
											"tag":     "01",
											"content": "936000091100004515",
											"data":    "0118936000091100004515",
										},
										"terminal_id": map[string]any{
											"tag":     "02",
											"content": "0489371081",
											"data":    "02100489371081",
										},
										"category": map[string]any{
											"tag":     "03",
											"content": "UMI",
											"data":    "0303UMI",
										},
									},
								},
								"switching": map[string]any{
									"tag":     "51",
									"content": "0014ID.CO.QRIS.WWW0215ID20200340731930303UKE",
									"data":    "51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE",
									"detail": map[string]any{
										"site": map[string]any{
											"tag":     "00",
											"content": "ID.CO.QRIS.WWW",
											"data":    "0014ID.CO.QRIS.WWW",
										},
										"nmid": map[string]any{
											"tag":     "02",
											"content": "ID2020034073193",
											"data":    "0215ID2020034073193",
										},
										"category": map[string]any{
											"tag":     "03",
											"content": "UKE",
											"data":    "0303UKE",
										},
									},
								},
								"merchant_category_code": map[string]any{
									"tag":     "52",
									"content": "4829",
									"data":    "52044829",
								},
								"currency_code": map[string]any{
									"tag":     "53",
									"content": "360",
									"data":    "5303360",
								},
								"country_code": map[string]any{
									"tag":     "58",
									"content": "ID",
									"data":    "5802ID",
								},
								"merchant_name": map[string]any{
									"tag":     "59",
									"content": "Sintas Store",
									"data":    "5912Sintas Store",
								},
								"merchant_city": map[string]any{
									"tag":     "60",
									"content": "Kota Yogyakarta",
									"data":    "6015Kota Yogyakarta",
								},
								"merchant_postal_code": map[string]any{
									"tag":     "61",
									"content": "55000",
									"data":    "610555000",
								},
								"additional_information": map[string]any{
									"tag":     "",
									"content": "",
									"data":    "",
								},
								"crc_code": map[string]any{
									"tag":     "63",
									"content": "1FA2",
									"data":    "63041FA2",
								},
							},
						},
						"error": map[string]any{
							"success": false,
							"message": "not static QRIS content detected",
							"data":    nil,
						},
					},
				},
				map[string]any{
					"1. Name":   "Convert Static QRIS Into Dynamic",
					"2. Method": "POST",
					"3. Path":   "/static-to-dynamic",
					"4. Body": map[string]any{
						"qr_string":            "000201010211y0ur4w3soMEsT47icQr15STriN6",
						"merchant_city":        "Kota Yogyakarta",
						"merchant_postal_code": "55000",
						"payment_amount":       1337,
						"payment_fee_category": "FIXED",
						"payment_fee":          666,
					},
					"5. Code Snippet": `curl --location 'https://api.qris.membasuh.com/static-to-dynamic' --header 'Content-Type: application/json' --data '{ "qr_string": "000201010211y0ur4w3soMEsT47icQr15STriN6", "merchant_city": "Kota Yogyakarta", "merchant_postal_code": "55000", "payment_amount": 1337, "payment_fee_category": "FIXED", "payment_fee": 666 }'`,
					"6. Example Response": map[string]any{
						"success": map[string]any{
							"success": true,
							"message": "Dynamic QRIS converted successfully",
							"data": map[string]any{
								"qr_string": "00020101021226630016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE5204482953033605404133755020256036665802ID5912Sintas Store6015Kota Yogyakarta61055500063040377",
								"qr_code":   "data:image/png;base64,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",
							},
						},
						"error": map[string]any{
							"success": false,
							"message": "invalid extract acquirer for content 0016COM.MEMBASUH.WWW0118936000091100004515021004893710810303U",
							"data":    nil,
						},
					},
				},
			},
		})
	})

	NewQRISRouter(env, publicRouter)
}
