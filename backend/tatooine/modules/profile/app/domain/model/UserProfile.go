package model

type UserProfile struct {
	Id                 int64             `json:"-"`
	Name               string            `json:"name"`
	LastName           string            `json:"lastName"`
	About              string            `json:"about"`
	AttandedActivities []Activity        `json:"attandedActivities"`
	Adress             UserProfileAdress `json:"address"`
	Stat               UserProfileStat   `json:"stats"`
	ProfileImageUrl    string            `json:"profileImageUrl"`
	ExternalId         string            `json:"externalId"`
}
