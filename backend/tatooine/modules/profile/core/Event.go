package core

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	CoverPhoto    string
	Name          string
	UserProfileID uint
}
