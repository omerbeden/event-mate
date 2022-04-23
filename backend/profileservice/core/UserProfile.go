package core

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	Name     string
	LastName string
	About    string
	Adress   UserProfileAdress
}
