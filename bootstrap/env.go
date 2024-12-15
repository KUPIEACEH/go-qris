package bootstrap

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	Port       string `mapstructure:"PORT"`
	QRCodeSize int    `mapstructure:"QR_CODE_SIZE"`
}

func NewEnv() *Env {
	env := Env{}
	godotenv.Load()

	// Load App Env
	appEnv, exists := os.LookupEnv("APP_ENV")
	if exists {
		env.AppEnv = appEnv
	}

	// Load Port
	port, exists := os.LookupEnv("PORT")
	if exists {
		env.Port = port
	} else {
		env.Port = "1337"
	}

	// Load QR Code Size
	qrCodeSize, exists := os.LookupEnv("QR_CODE_SIZE")
	if exists {
		size, err := strconv.Atoi(qrCodeSize)
		if err != nil {
			log.Fatalf("Invalid QR_CODE_SIZE value: %s", qrCodeSize)
		}
		env.QRCodeSize = size
	}
	if env.QRCodeSize < 60 {
		env.QRCodeSize = 256
	}

	if env.AppEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("The App is running in development env")
	}

	return &env
}
