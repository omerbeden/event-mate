package core

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	LastName string
	About    string
	Adress   UserAdress
}
