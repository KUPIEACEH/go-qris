package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fyvri/go-qris/internal/domain/entities"
	"github.com/fyvri/go-qris/internal/interface/controllers"
	"github.com/gin-gonic/gin"
)

func TestNewQRIS(t *testing.T) {
	tests := []struct {
		name   string
		fields QRIS
		want   QRISInterface
	}{
		{
			name:   "Success: No Field",
			fields: QRIS{},
			want:   &QRIS{},
		},
		{
			name: "Success: With Field",
			fields: QRIS{
				qrisController: &controllers.QRIS{},
			},
			want: &QRIS{
				qrisController: &controllers.QRIS{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uc := NewQRIS(test.fields.qrisController)

			if uc == nil {
				t.Errorf(expectedReturnNonNil, "NewQRIS", "QRISInterface")
			}

			got, ok := uc.(*QRIS)
			if !ok {
				t.Errorf(expectedTypeAssertionErrorMessage, "*QRIS")
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf(expectedButGotMessage, "*QRIS", test.want, got)
			}
		})
	}
}

func TestQRISParse(t *testing.T) {
	type args struct {
		requestBody string
	}
	type want struct {
		code     int
		response string
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   want
	}{
		{
			name:   testNameInvalidJSON,
			fields: QRIS{},
			args: args{
				requestBody: `"{"qr_string": 1337}"`,
			},
			want: want{
				code:     http.StatusBadRequest,
				response: `cannot unmarshal string`,
			},
		},
		{
			name: "Error: h.qrisController.Parse()",
			fields: QRIS{
				qrisController: &mockQRISController{
					ParseFunc: func(qrisString string) (*entities.QRIS, error, *[]string) {
						return nil, fmt.Errorf("invalid QR string"), nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "invalid"}`,
			},
			want: want{
				code:     http.StatusInternalServerError,
				response: `"invalid QR string"`,
			},
		},
		{
			name: "Success",
			fields: QRIS{
				qrisController: &mockQRISController{
					ParseFunc: func(qrisString string) (*entities.QRIS, error, *[]string) {
						return &entities.QRIS{}, nil, nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "valid"}`,
			},
			want: want{
				code:     http.StatusOK,
				response: `"QRIS parsed successfully"`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewQRIS(test.fields.qrisController)

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.POST("/", handler.Parse)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(test.args.requestBody))
			req.Header.Set(testHeaderContentType, testHeaderContentTypeValue)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			if recorder.Code != test.want.code {
				t.Errorf(expectedStatusCode, test.want.code, recorder.Code)
			}
			if !bytes.Contains(recorder.Body.Bytes(), []byte(test.want.response)) {
				t.Errorf(expectedResponseToContain, test.want.response, recorder.Body.String())
			}
		})
	}
}

func TestQRISConvert(t *testing.T) {
	type args struct {
		requestBody string
	}
	type want struct {
		code     int
		response string
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   want
	}{
		{
			name:   testNameInvalidJSON,
			fields: QRIS{},
			args: args{
				requestBody: `"{"qr_string": 1337}"`,
			},
			want: want{
				code:     http.StatusBadRequest,
				response: `cannot unmarshal string`,
			},
		},
		{
			name: "Error: h.qrisController.Convert()",
			fields: QRIS{
				qrisController: &mockQRISController{
					ConvertFunc: func(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) (string, string, error, *[]string) {
						return "", "", fmt.Errorf("invalid QR string"), nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "invalid"}`,
			},
			want: want{
				code:     http.StatusInternalServerError,
				response: `"invalid QR string"`,
			},
		},
		{
			name: "Success",
			fields: QRIS{
				qrisController: &mockQRISController{
					ConvertFunc: func(qrisString string, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue uint32, paymentFeeCategoryValue string, paymentFeeValue uint32, terminalLabelValue string) (string, string, error, *[]string) {
						return "QR Dynamic String", "QR Dynamic Code", nil, nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "valid"}`,
			},
			want: want{
				code:     http.StatusOK,
				response: `"Dynamic QRIS converted successfully"`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewQRIS(test.fields.qrisController)

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.POST("/", handler.Convert)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(test.args.requestBody))
			req.Header.Set(testHeaderContentType, testHeaderContentTypeValue)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			if recorder.Code != test.want.code {
				t.Errorf(expectedStatusCode, test.want.code, recorder.Code)
			}
			if !bytes.Contains(recorder.Body.Bytes(), []byte(test.want.response)) {
				t.Errorf(expectedResponseToContain, test.want.response, recorder.Body.String())
			}
		})
	}
}

func TestQRISIsValid(t *testing.T) {
	type args struct {
		requestBody string
	}
	type want struct {
		code     int
		response string
	}

	tests := []struct {
		name   string
		fields QRIS
		args   args
		want   want
	}{
		{
			name:   testNameInvalidJSON,
			fields: QRIS{},
			args: args{
				requestBody: `"{"qr_string": 1337}"`,
			},
			want: want{
				code:     http.StatusBadRequest,
				response: `cannot unmarshal string`,
			},
		},
		{
			name: "Error: h.qrisController.IsValid()",
			fields: QRIS{
				qrisController: &mockQRISController{
					IsValidFunc: func(qrisString string) (error, *[]string) {
						return fmt.Errorf("invalid CRC16-CCITT code"), nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "invalid"}`,
			},
			want: want{
				code:     http.StatusInternalServerError,
				response: `"invalid CRC16-CCITT code"`,
			},
		},
		{
			name: "Success",
			fields: QRIS{
				qrisController: &mockQRISController{
					IsValidFunc: func(qrisString string) (error, *[]string) {
						return nil, nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "valid"}`,
			},
			want: want{
				code:     http.StatusOK,
				response: `"CRC16-CCITT code is valid"`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewQRIS(test.fields.qrisController)

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.POST("/", handler.IsValid)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(test.args.requestBody))
			req.Header.Set(testHeaderContentType, testHeaderContentTypeValue)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			if recorder.Code != test.want.code {
				t.Errorf(expectedStatusCode, test.want.code, recorder.Code)
			}
			if !bytes.Contains(recorder.Body.Bytes(), []byte(test.want.response)) {
				t.Errorf(expectedResponseToContain, test.want.response, recorder.Body.String())
			}
		})
	}
}
