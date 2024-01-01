package model

type UserProfile struct {
	Id                 int64
	Name               string
	LastName           string
	About              string
	AttandedActivities []Activity
	Adress             UserProfileAdress
	Stat               UserProfileStat
	ProfileImageUrl    string
}
