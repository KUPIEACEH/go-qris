# Go-QRIS

Go-QRIS is a Go-based project designed to convert QRIS code into dynamic ones. QRIS (Quick Response Code Indonesian Standard) is widely used for payments, but QR code has limitations in flexibility. This tool enhances QRIS transactions by enabling dynamic data like payment amounts, merchant details, and fees, making payments more adaptable and efficient. Go-QRIS simplifies the process of generating dynamic QRIS code, improving payment flexibility for businesses and providing a seamless experience for customers.

## 📝 Directory Structure

```
├── .github         # CI/CD workflows
├── api             # API endpoints
│   ├── handlers    # Request handlers for API endpoints
│   └── routes      # Route definitions for QRIS APIs
├── bootstrap       # Application initialization
├── cmd             # Application entry point
├── deployments     # Deployment configurations
├── internal        # Core application logic
│   ├── config      # Configuration management for internal modules
│   ├── domain      # Business domain entities
│   ├── interface   # Contains related logic to use case
│   └── usecases    # Application use cases
├── pkg             # Independent libraries
├── .dockerignore   # Docker ignore file
├── .env.example    # Example environment configuration
├── .gitignore      # Git ignore file
├── go.mod          # Go module configuration
├── go.sum          # Go dependencies
├── LICENSE         # Project license
└── README.md       # Project documentation
```

## 💌 Prerequisites

- **Go**: Ensure the latest version of Go is installed.
- **Docker** (optional): To run the application in an isolated environment.

## 🛠️ Installation

1.  Clone this repository:

    ```bash
    git clone git@github.com:fyvri/go-qris.git && cd go-qris
    ```

2.  Copy the `.env.example` file to `.env` and adjust the configuration:

    ```bash
    cp .env.example .env
    ```

3.  Install dependencies:

    ```bash
    go mod tidy
    ```

## ⚙️ Running the Application

1.  Run the application locally:

    ```bash
    go run ./cmd/main.go
    ```

2.  Run the application using Docker:

    ```bash
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-qris -trimpath ./cmd/main.go
    docker build -f ./deployments/Dockerfile -t go-qris .
    docker run --name go-qris -e APP_ENV=development -e QR_CODE_SIZE=256 -p 8080:1337 go-qris
    ```

    Alternatively, open the following url in your browser: [https://github.com/fyvri/go-qris/pkgs/container/go-qris](https://github.com/fyvri/go-qris/pkgs/container/go-qris)

3.  Implement into your own awesome project:

    ```go
    package main

    import (
        "fmt"

        "github.com/fyvri/go-qris/pkg/services"
    )

    func main() {
        qrisString := "000201010211y0ur4w3soMEQr15STriN6"
        merchantCity := "Kota Yogyakarta"                    // optional
        merchantPostalCode := "55000"                        // optional
        paymentAmount := 1337                                // mandatory
        paymentFeeCategory := "FIXED"                        // optional, value: FIXED or PERCENT
        paymentFee := 666                                    // optional, based on paymentFeeCategory value
        terminalLabel := "Made with love by Alvriyanto Azis" // optional, it works if terminal label exists in qrisString

        qrisService := services.NewQRIS()
        qrisString, err, errs := qrisService.Convert(qrisString, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee, terminalLabel)
        if err != nil {
            fmt.Println("[ FAILURE ]", err)
            if errs != nil {
                for _, err := range *errs {
                    fmt.Println("            -", err)
                }
            }
            return
        }
        fmt.Println("[ SUCCESS ]", qrisString)
    }
    ```

    Here are additional functions you can use to interact:

    - **Parse QRIS**

      `Parse(qrisString string) (*models.QRIS, error, *[]string)`

      ```go
      qris, err, errs := qrisService.Parse(qrisString)
      ```

    - **Validate QRIS**

      `IsValid(qris *models.QRIS) bool`

      ```go
      isValid := qrisService.IsValid(qris)
      ```

    - **Modify QRIS**

      `Modify(qris *models.QRIS, merchantCityValue string, merchantPostalCodeValue string, paymentAmountValue int, paymentFeeCategoryValue string, paymentFeeValue int, terminalLabelValue string) (*models.QRIS, error, *[]string)`

      ```go
      qris, err, errs = qrisService.Modify(qris, merchantCity, merchantPostalCode, paymentAmount, paymentFeeCategory, paymentFee, terminalLabel)
      ```

    - **Convert QRIS to String**

      `ToString(qris *models.QRIS) string`

      ```go
      qrisString = qrisService.ToString(qris)
      ```

## 🧪 Testing

1.  Run all unit tests:

    ```bash
    go test ./...
    ```

2.  Check test coverage:

    ```bash
    go test ./... -cover
    ```

## 🔥 API Endpoints

To learn more about the available endpoints, you can refer to [Postman Documentation](https://documenter.getpostman.com/view/6937269/2sAYJ1jMc7) 🦸

1.  **Parse QRIS**

    - Endpoint: `POST /parse`
    - Content-Type: `application/json`
    - Request Body:

      ```json
      {
        "qr_string": "000201010211y0ur4w3soMEQr15STriN6"
      }
      ```

    - Example Response:

      `Success`

      ```json
      {
        "success": true,
        "message": "QRIS parsed successfully",
        "errors": null,
        "data": {
          "version": {
            "tag": "00",
            "content": "01",
            "data": "000201"
          },
          "category": {
            "tag": "01",
            "content": "11",
            "data": "010211"
          },
          "acquirer": {
            "tag": "26",
            "content": "0016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI",
            "data": "26630016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI",
            "detail": {
              "site": {
                "tag": "00",
                "content": "COM.MEMBASUH.WWW",
                "data": "0016COM.MEMBASUH.WWW"
              },
              "mpan": {
                "tag": "01",
                "content": "936000091100004515",
                "data": "0118936000091100004515"
              },
              "terminal_id": {
                "tag": "02",
                "content": "0489371081",
                "data": "02100489371081"
              },
              "category": {
                "tag": "03",
                "content": "UMI",
                "data": "0303UMI"
              }
            }
          },
          "switching": {
            "tag": "51",
            "content": "0014ID.CO.QRIS.WWW0215ID20200340731930303UKE",
            "data": "51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE",
            "detail": {
              "site": {
                "tag": "00",
                "content": "ID.CO.QRIS.WWW",
                "data": "0014ID.CO.QRIS.WWW"
              },
              "nmid": {
                "tag": "02",
                "content": "ID2020034073193",
                "data": "0215ID2020034073193"
              },
              "category": {
                "tag": "03",
                "content": "UKE",
                "data": "0303UKE"
              }
            }
          },
          "merchant_category_code": {
            "tag": "52",
            "content": "4829",
            "data": "52044829"
          },
          "currency_code": {
            "tag": "53",
            "content": "360",
            "data": "5303360"
          },
          "payment_amount": {
            "tag": "",
            "content": "",
            "data": ""
          },
          "payment_fee_category": {
            "tag": "",
            "content": "",
            "data": ""
          },
          "payment_fee": {
            "tag": "",
            "content": "",
            "data": ""
          },
          "country_code": {
            "tag": "58",
            "content": "ID",
            "data": "5802ID"
          },
          "merchant_name": {
            "tag": "59",
            "content": "Sintas Store",
            "data": "5912Sintas Store"
          },
          "merchant_city": {
            "tag": "60",
            "content": "Kota Yogyakarta",
            "data": "6015Kota Yogyakarta"
          },
          "merchant_postal_code": {
            "tag": "61",
            "content": "55000",
            "data": "610555000"
          },
          "additional_information": {
            "tag": "62",
            "content": "0703A01",
            "data": "62070703A01",
            "detail": {
              "bill_number": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "mobile_number": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "store_label": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "loyalty_number": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "reference_label": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "customer_label": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "terminal_label": {
                "tag": "07",
                "content": "A01",
                "data": "0703A01"
              },
              "purpose_of_transaction": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "additional_consumer_data_request": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "merchant_tax_id": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "merchant_channel": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "rfu": {
                "tag": "",
                "content": "",
                "data": ""
              },
              "payment_system_specific": {
                "tag": "",
                "content": "",
                "data": ""
              }
            }
          },
          "crc_code": {
            "tag": "63",
            "content": "9FB7",
            "data": "63049FB7"
          }
        }
      }
      ```

      `Error`

      ```json
      {
        "success": false,
        "message": "invalid QRIS format",
        "errors": [
          "Acquirer tag is missing",
          "Country code tag is missing",
          "CRC code tag is missing"
        ],
        "data": null
      }
      ```

2.  **Convert QRIS into a Dynamic Version**

    - Endpoint: `POST /convert`
    - Content-Type: `application/json`
    - Request Body:

      ```json
      {
        "qr_string": "000201010211y0ur4w3soMEQr15STriN6",
        "merchant_city": "Kota Yogyakarta", // optional
        "merchant_postal_code": "55000", // optional
        "payment_amount": 1337, // mandatory
        "payment_fee_category": "FIXED", // optional, value: FIXED or PERCENT
        "payment_fee": 666, // optional, based on payment fee category
        "terminalLabel": "Made with love by Alvriyanto Azis" // optional, it works if terminal label exists in qr string
      }
      ```

    - Example Response:

      `Success`

      ```json
      {
        "success": true,
        "message": "Dynamic QRIS converted successfully",
        "errors": null,
        "data": {
          "qr_string": "00020101021226630016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE5204482953033605404133755020256036665802ID5912Sintas Store6015Kota Yogyakarta61055500062070703A016304F98B",
          "qr_code": "data:image/png;base64,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"
        }
      }
      ```

      `Error`

      ```json
      {
        "success": false,
        "message": "invalid parse acquirer for content 0016COM.MEMBASUH.WWW0118936000091100004515021004893710810303",
        "errors": null,
        "data": null
      }
      ```

3.  **Validate QRIS**

    - Endpoint: `POST /is-valid`
    - Content-Type: `application/json`
    - Request Body:

      ```json
      {
        "qr_string": "000201010211y0ur4w3soMEQr15STriN6"
      }
      ```

    - Example Response:

      `Success`

      ```json
      {
        "success": true,
        "message": "CRC16-CCITT code is valid",
        "errors": null,
        "data": null
      }
      ```

      `Error`

      ```json
      {
        "success": false,
        "message": "invalid CRC16-CCITT code",
        "errors": null,
        "data": null
      }
      ```

## 👥 Contribution

If you have any ideas, [open an issue](https://github.com/fyvri/go-qris/issues/new) and tell me what you think.

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

> [!IMPORTANT]
> If you have a suggestion that would make this better, please fork the repo and create a pull request. Don't forget to give the project a star 🌟 I can't stop saying thank you!
>
> 1. Fork this project
> 2. Create your feature branch (`git checkout -b feature/awesome-feature`)
> 3. Commit your changes (`git commit -m "feat: add awesome feature"`)
> 4. Push to the branch (`git push origin feature/awesome-feature`)
> 5. Open a pull request

## 📜 License

This project is licensed under [MIT License](LICENSE). Feel free to use and modify it as needed.
