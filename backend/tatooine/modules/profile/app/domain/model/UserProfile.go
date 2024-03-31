package model

type UserProfile struct {
	Id                 int64                   `json:"id"`
	Name               string                  `json:"name"`
	LastName           string                  `json:"lastName"`
	About              string                  `json:"about"`
	AttandedActivities []Activity              `json:"attandedActivities"`
	Adress             UserProfileAdress       `json:"address"`
	Stat               UserProfileStat         `json:"stats"`
	ProfileImageUrl    string                  `json:"profileImageUrl"`
	ExternalId         string                  `json:"externalId"`
	UserName           string                  `json:"userName"`
	Email              string                  `json:"email"`
	Badges             map[int64]*ProfileBadge `json:"badges"`
	IsVerified         bool                    `json:"isVerified"`
}
