package model

type Event struct {
	ID        int64
	Title     string
	Category  string
	CreatedBy User
	Location  Location
}

type Location struct {
	ID   int64
	City string
}
