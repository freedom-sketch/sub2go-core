package linkutils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

// Encrypts a string using RSA-4096
func Encrypt(plaintext string) (string, error) {
	reqBody, err := json.Marshal(HappEncryptRequest{URL: plaintext})
	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://crypto.happ.su/api.php", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse HappEncryptResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", err
	}

	return apiResponse.EncryptedLink, nil
}

// Returns a base64 encoded string with the "base64" prefix
func Base64Encode(plaintext string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(plaintext))
	return "base64:" + encoded
}
