package model

type UserProfileAdress struct {
	ProfileId int64  `json:"-"`
	City      string `json:"city"`
	//UserProfileId int64
}
