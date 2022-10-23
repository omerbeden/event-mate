package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title     string
	Category  string
	CreatedBy User `gorm:"foreignKey:ID"`
	Location  Location
}

type Location struct {
	City   string
	County string
}
