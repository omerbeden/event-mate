package interfaces

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
)

type UserRepository interface {
	//TODO: do implement in infra
	GetUsers() []core.UserProfile
	GetUserById(id uint)
	InsertUser(user *core.UserProfile)
	UpdateUser(user *core.UserProfile)
	DeleteUserById(id uint)
}
