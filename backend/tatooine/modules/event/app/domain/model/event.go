package model

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title     string
	Category  string
	CreatedBy User
	Location  Location
}

type Location struct {
	gorm.Model
	City    string
	EventID uint
}
