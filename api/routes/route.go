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
			Data: map[string]any{
				"1_Title":       "Go QRIS",
				"2_Description": "Go QRIS is a Go-based project designed to convert QRIS codes into dynamic ones. QRIS (Quick Response Code Indonesian Standard) is widely used for payments, but QR codes have limitations in flexibility. This tool enhances QRIS transactions by enabling dynamic data like payment amounts, merchant details, and fees, making payments more adaptable and efficient. Go QRIS simplifies the process of generating dynamic QR codes, improving payment flexibility for businesses and providing a seamless experience for customers.",
				"3_Resources": [2]any{
					"https://github.com/fyvri/go-qris",
					"https://hub.docker.com/r/azisalvriyanto/go-qris",
				},
				"4_API_Endpoints": [3]any{
					map[string]any{
						"1_Name":   "Parse QRIS",
						"2_Method": "POST",
						"3_Path":   "/parse",
						"4_Body": map[string]any{
							"qr_string": "000201010211y0ur4w3soMEQr15STriN6",
						},
						"5_Code_Snippet": `curl --location 'https://api.qris.membasuh.com/parse' --header 'Content-Type: application/json' --data '{"qr_string": "000201010211y0ur4w3soMEQr15STriN6"}'`,
					},
					map[string]any{
						"1_Name":   "Convert QRIS Into Dynamic",
						"2_Method": "POST",
						"3_Path":   "/to-dynamic",
						"4_Body": map[string]any{
							"qr_string":            "000201010211y0ur4w3soMEQr15STriN6",
							"merchant_city":        "Kota Yogyakarta",
							"merchant_postal_code": "55000",
							"payment_amount":       1337,
							"payment_fee_category": "FIXED",
							"payment_fee":          666,
						},
						"5_Code_Snippet": `curl --location 'https://api.qris.membasuh.com/to-dynamic' --header 'Content-Type: application/json' --data '{ "qr_string": "000201010211y0ur4w3soMEQr15STriN6", "merchant_city": "Kota Yogyakarta", "merchant_postal_code": "55000", "payment_amount": 1337, "payment_fee_category": "FIXED", "payment_fee": 666 }'`,
					},
					map[string]any{
						"1_Name":   "Validate QRIS",
						"2_Method": "POST",
						"3_Path":   "/validate",
						"4_Body": map[string]any{
							"qr_string": "000201010211y0ur4w3soMEQr15STriN6",
						},
						"5_Code_Snippet": `curl --location 'https://api.qris.membasuh.com/validate' --header 'Content-Type: application/json' --data '{"qr_string": "000201010211y0ur4w3soMEQr15STriN6"}'`,
					},
				},
			},
		})
	})

	NewQRISRouter(env, publicRouter)
}
