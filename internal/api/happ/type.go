package happ

// Structure for sending a request to the Happ API (string to encrypt)
type Request struct {
	URL string `json:"url"`
}

// Structure for receiving a response from the Happ API (encrypted string)
type Response struct {
	EncryptedLink string `json:"encrypted_link"`
}
