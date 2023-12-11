package model

type Activity struct {
	ID           int64
	Title        string
	Category     string
	CreatedBy    User
	Location     Location
	Participants []User
}

type Location struct {
	ActivityId int64
	City       string
}
