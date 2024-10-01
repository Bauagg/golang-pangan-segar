package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST         string
	DB_PORT         string
	DB_NAMA         string
	DB_USER         string
	DB_PASSWORD     string
	APP_PORT        string
	SECRETKEY_TOKEN string
	URL_HOST_SERVER string

	// KONFIGURASU OTP PHONE
	URL_SEND_PHONE_OTP string
	API_KEY_SEND_OTP   string
	ID_SEND_OTP        string
)

func IntConfigEnv() {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, proceeding with environment variables")
	}

	// Fetch values from environment variables
	appPort := os.Getenv("APP_PORT")
	if appPort != "" {
		APP_PORT = appPort
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		DB_HOST = dbHost
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort != "" {
		DB_PORT = dbPort
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		DB_NAMA = dbName
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser != "" {
		DB_USER = dbUser
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		DB_PASSWORD = dbPassword
	}

	secretKeyToken := os.Getenv("SECRETKEY_TOKEN")
	if secretKeyToken != "" {
		SECRETKEY_TOKEN = secretKeyToken
	}

	urlHostServer := os.Getenv("URL_HOST_SERVER")
	if urlHostServer != "" {
		URL_HOST_SERVER = urlHostServer
	}

	// KONFIGURASU OTP PHONE
	urlSendPhoneOTP := os.Getenv("URL_SEND_PHONE_OTP")
	if urlSendPhoneOTP != "" {
		URL_SEND_PHONE_OTP = urlSendPhoneOTP
	}

	appKeySendOTP := os.Getenv("API_KEY_SEND_OTP")
	if appKeySendOTP != "" {
		API_KEY_SEND_OTP = appKeySendOTP
	}

	idSendOTP := os.Getenv("ID_SEND_OTP")
	if idSendOTP != "" {
		ID_SEND_OTP = idSendOTP
	}
}
