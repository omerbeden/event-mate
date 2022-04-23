package core

import "gorm.io/gorm"

type UserProfileAdress struct {
	gorm.Model
	City          string
	UserProfileID uint
}
