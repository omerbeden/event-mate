package core

import "gorm.io/gorm"

type UserAdress struct {
	gorm.Model
	City   string
	UserID uint
}
