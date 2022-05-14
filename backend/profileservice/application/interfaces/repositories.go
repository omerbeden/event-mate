package interfaces

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
)

type UserRepository interface {
	GetUsers() []core.UserProfile
	GetUserById(id uint)
	InsertUser(user *core.UserProfile)
	UpdateUser(user *core.UserProfile)
	DeleteUserById(id uint)
}

type EventRepository interface {
	GetUserEvent(userId int) ([]core.Event, error)
}

type UserStatRepository interface {
	GetUserStat(userId int) (core.UserProfileStat, error)
}
