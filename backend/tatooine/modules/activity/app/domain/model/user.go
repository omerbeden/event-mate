package model

type User struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	LastName        string  `json:"lastName"`
	Username        string  `json:"username"`
	ProfileImageUrl string  `json:"profileImageUrl"`
	ProfilePoint    float64 `json:"points"`
	ExternalId      string  `json:"externalId,omitempty"`
	About           string  `json:"about,omitempty"`
}
