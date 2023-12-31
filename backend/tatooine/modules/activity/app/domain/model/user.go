package model

type User struct {
	ID              int64  `json:"-"`
	Name            string `json:"name"`
	LastName        string `json:"lastName"`
	ProfileImageUrl string `json:"profileImage"`
	ProfilePoint    int    `json:"profilePoint"`
}
