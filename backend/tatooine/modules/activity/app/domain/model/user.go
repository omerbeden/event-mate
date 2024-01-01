package model

type User struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	LastName        string  `json:"lastName"`
	ProfileImageUrl string  `json:"profileImage"`
	ProfilePoint    float64 `json:"profilePoint"`
}
