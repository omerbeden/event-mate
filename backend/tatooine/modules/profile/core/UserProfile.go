package core

type UserProfile struct {
	Name           string
	LastName       string
	About          string
	Photo          string
	AttandedEvents []Event
	Adress         UserProfileAdress
	UserId         uint
}
