package model

type Event struct {
	ID           int64
	Title        string
	Category     string
	CreatedBy    User
	Location     Location
	Participants []User
}

type Location struct {
	ID   int64
	City string
}
