package interfaces

type UserRepository interface {
	//TODO: do implement in infra
	GetUser()
	InsertUser()
	UpdateUser()
	DeleteUser()
}
