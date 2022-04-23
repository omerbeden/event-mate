package core

import "gorm.io/gorm"

type UserProfileStat struct {
	gorm.Model
	Followers      int
	Following      int
	AttandedEvents int
	Points         float32
	UserProfileID  uint
}
