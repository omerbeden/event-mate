package model

type VeriffAPICred struct {
	ApiKey    string
	SecretKey string
	BaseUrl   string
}
type VeriffRequest struct {
	Callback   string `json:"callback"`
	VendorData string `json:"vendorData"`
}

type VeriffResponse struct {
	Id           string `json:"id"`
	Url          string `json:"url"`
	VendorData   string `json:"vendorData"`
	Host         string `json:"host"`
	Status       string `json:"status"`
	SessionToken string `json:"sessionToken"`
}
