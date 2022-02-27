package interfaces

import (
	"github.com/omerbeden/event-mate/backend/profileservice/core"
)

type UserRepository interface {
	//TODO: do implement in infra
	GetUsers() []core.User
	InsertUser()
	UpdateUser()
	DeleteUser()
}
