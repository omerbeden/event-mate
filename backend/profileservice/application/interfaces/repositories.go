package interfaces

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
)

type UserRepository interface {
	//TODO: do implement in infra
	GetUsers() []core.User
	GetUserById(id uint)
	InsertUser(user *core.User)
	UpdateUser(user *core.User)
	DeleteUserById(id uint)
}
