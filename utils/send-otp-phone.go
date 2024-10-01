package utils

import (
	"fmt"
	"log"
	"pangan-segar/config"

	"github.com/go-resty/resty/v2"
)

func SendPhoneOtp(phone string, kode uint64) (map[string]interface{}, int, error) {
	// Split the phone number if it starts with '08'
	if len(phone) > 2 && phone[:2] == "08" {
		phone = "628" + phone[2:]
	}

	// Prepare the parameters
	params := map[string]interface{}{
		"from": config.ID_SEND_OTP,
		"to":   phone,
		"text": fmt.Sprintf("KODE OTP MBG: %d berlaku selama 5 menit. RAHASIAKAN KODE OTP MBG Anda! Jangan beritahukan kepada SIAPAPUN!", kode),
	}

	// Create a new Resty client
	client := resty.New()

	// Create a variable to hold the response result
	var result map[string]interface{}

	// Make the POST request to the OTP service
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", config.API_KEY_SEND_OTP).
		SetBody(params).
		SetResult(&result). // Automatically unmarshals response into the 'result' map
		Post(config.URL_SEND_PHONE_OTP)

	// Handle any error during the request
	if err != nil {
		log.Println("Error sending OTP:", err)
		return nil, 0, err
	}

	// Return the parsed result and HTTP status code
	return result, resp.StatusCode(), nil
}
