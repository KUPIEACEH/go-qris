# Go QRIS

Go QRIS is a Go-based project designed to convert static QRIS codes into dynamic ones. QRIS (Quick Response Code Indonesian Standard) is widely used for payments, but static QR codes have limitations in flexibility. This tool enhances QRIS transactions by enabling dynamic data like payment amounts, merchant details, and fees, making payments more adaptable and efficient. Go QRIS simplifies the process of generating dynamic QR codes, improving payment flexibility for businesses and providing a seamless experience for customers.

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
└── LICENSE.md      # Project license
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

## 🎉 Running the Application

1.  Run the application locally:

    ```bash
    go run cmd/main.go
    ```

2.  Run the application using Docker:

    ```bash
    docker build -f ./deployments/Dockerfile -t go-qris .
    docker run --name go-qris -e APP_ENV=development -e QR_CODE_SIZE=200 -p 8080:1337 go-qris
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

## 🚀 Deployment

Use the configuration files in the `deployments` folder to set up deployment in a production environment.

## 🔥 API Endpoints

1.  **Extract Static QRIS**

    - Endpoint: `POST /extract-static`
    - URL: `https://api.qris.membasuh.com/extract-static`
    - Method: `POST`
    - Content-Type: `application/json`
    - Request Body:

      ```json
      {
        "qr_string": "000201010211y0ur4w3soMEsT47icQr15STriN6"
      }
      ```

    - Example Response:

      `Success`

      ```json
      {
        "success": true,
        "message": "QRIS extracted successfully",
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
            "tag": "",
            "content": "",
            "data": ""
          },
          "crc_code": {
            "tag": "63",
            "content": "1FA2",
            "data": "63041FA2"
          }
        }
      }
      ```

      `Error`

      ```json
      {
        "success": false,
        "message": "not static QRIS content detected",
        "data": null
      }
      ```

2.  **Convert Static QRIS Into Dynamic**

    - Endpoint: `POST /static-to-dynamic`
    - URL: `https://api.qris.membasuh.com/static-to-dynamic`
    - Method: `POST`
    - Content-Type: `application/json`
    - Request Body:

      ```json
      {
        "qr_string": "000201010211y0ur4w3soMEsT47icQr15STriN6",
        "merchant_city": "Kota Yogyakarta", // optional
        "merchant_postal_code": "55000", // optional
        "payment_amount": 1337,
        "payment_fee_category": "FIXED", // optional, value: FIXED or PERCENT
        "payment_fee": 666 // optional based on payment fee category
      }
      ```

      - Example Response:

        `Success`

        ```json
        {
          "success": true,
          "message": "Dynamic QRIS converted successfully",
          "data": {
            "qr_string": "00020101021226630016COM.MEMBASUH.WWW0118936000091100004515021004893710810303UMI51440014ID.CO.QRIS.WWW0215ID20200340731930303UKE5204482953033605404133755020256036665802ID5912Sintas Store6015Kota Yogyakarta61055500063040377",
            "qr_code": "data:image/png;base64,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"
          }
        }
        ```

        `Error`

        ```json
        {
          "success": false,
          "message": "invalid extract acquirer for content 0016COM.MEMBASUH.WWW0118936000091100004515021004893710810303U",
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

This project is licensed under [MIT License](LICENSE.md). Feel free to use and modify it as needed.
