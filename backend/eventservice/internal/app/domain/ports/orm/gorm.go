package database

import "gorm.io/gorm"

type Database interface {
	NewConnection() *gorm.DB
}
