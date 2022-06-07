package interfaces

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
)

type UserRepository interface {
	GetUsers() ([]core.UserProfile, error)
	GetUserById(id uint) (core.UserProfile, error)
	InsertUser(user *core.UserProfile) error
	UpdateUser(user *core.UserProfile) error
	DeleteUserById(id uint)
}

type EventRepository interface {
	GetUserEvent(userId int) ([]core.Event, error)
}

type UserStatRepository interface {
	GetUserStat(userId int) (core.UserProfileStat, error)
}
