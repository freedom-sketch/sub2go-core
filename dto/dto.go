package dto

// Structure for sending a request to the Happ API (string to encrypt)
type HappEncryptRequest struct {
	URL string `json:"url"`
}

// Structure for receiving a response from the Happ API (encrypted string)
type HappEncryptResponse struct {
	EncryptedLink string `json:"encrypted_link"`
}

//-----------Xray-----------
type InboundTemplateData struct {
	Tag             string
	Port            int
	Target          string
	PrivateKey      string
	ServerNamesJSON string
	ShortIdsJSON    string
}
