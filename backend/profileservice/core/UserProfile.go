package core

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	Name           string
	LastName       string
	About          string
	Photo          string
	AttandedEvents []Event
	Adress         UserProfileAdress
	UserId         uint
}
