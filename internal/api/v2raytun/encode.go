package v2raytun

import (
	"encoding/base64"
)

// Returns a base64 encoded string with the "base64" prefix.
func Base64Encode(plaintext string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(plaintext))
	return "base64:" + encoded
}
