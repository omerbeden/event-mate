package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnPG() *gorm.DB {
	dsn := "postgres://postgres:password@localhost:5432/test" //test
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func InitMigration(models ...interface{}) {
	db := NewConnPG()
	db.AutoMigrate(models...)
}
