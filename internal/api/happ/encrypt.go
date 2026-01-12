package happ

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var encryptionURL = "https://crypto.happ.su/api.php"

// Structure for sending a request to the Happ API (string to encrypt)
type Request struct {
	URL string `json:"url"`
}

// Structure for receiving a response from the Happ API (encrypted string)
type Response struct {
	EncryptedLink string `json:"encrypted_link"`
}

// Encrypts a string using RSA-4096
func Encrypt(plaintext string) (string, error) {
	reqBody, err := json.Marshal(Request{URL: plaintext})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(encryptionURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", err
	}

	return apiResponse.EncryptedLink, nil
}
