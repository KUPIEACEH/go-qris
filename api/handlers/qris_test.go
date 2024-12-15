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

func TestQRISExtractStatic(t *testing.T) {
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
			name:   "Error: Invalid JSON",
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
			name: "Error: ExtractStatic",
			fields: QRIS{
				qrisController: &mockQRISController{
					ExtractStaticFunc: func(qris string) (*entities.QRISStatic, error) {
						return nil, fmt.Errorf("invalid QR string")
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
					ExtractStaticFunc: func(qris string) (*entities.QRISStatic, error) {
						return nil, nil
					},
				},
			},
			args: args{
				requestBody: `{"qr_string": "valid"}`,
			},
			want: want{
				code:     http.StatusOK,
				response: `"QRIS extracted successfully"`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewQRIS(test.fields.qrisController)

			gin.SetMode(gin.TestMode)
			router := gin.Default()
			router.POST("/", handler.ExtractStatic)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(test.args.requestBody))
			req.Header.Set("Content-Type", "application/json")

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

func TestQRISStaticToDynamic(t *testing.T) {
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
			name:   "Error: Invalid JSON",
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
			name: "Error: StaticToDynamic",
			fields: QRIS{
				qrisController: &mockQRISController{
					StaticToDynamicFunc: func(qris string, merchantName string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error) {
						return "", "", fmt.Errorf("invalid QR string")
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
					StaticToDynamicFunc: func(qris string, merchantName string, merchantCity string, merchantPostalCode string, paymentAmount uint32, paymentFeeCategory string, paymentFee uint32) (string, string, error) {
						return "QR String", "QR Code", nil
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
			router.POST("/", handler.StaticToDynamic)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(test.args.requestBody))
			req.Header.Set("Content-Type", "application/json")

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
