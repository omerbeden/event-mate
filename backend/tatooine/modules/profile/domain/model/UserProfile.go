package model

type UserProfile struct {
	Id             int64
	Name           string
	LastName       string
	About          string
	Photo          string
	AttandedEvents []Event
	Adress         UserProfileAdress
	Stat           UserProfileStat
}
