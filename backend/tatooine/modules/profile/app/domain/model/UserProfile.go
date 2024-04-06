package model

type UserProfile struct {
	Id                 int64                   `json:"id"`
	Header             UserProfileHeader       `json:"header"`
	About              string                  `json:"about"`
	AttandedActivities []Activity              `json:"attandedActivities"`
	Adress             UserProfileAdress       `json:"address"`
	Stat               UserProfileStat         `json:"stats"`
	ExternalId         string                  `json:"externalId"`
	Email              string                  `json:"email"`
	Badges             map[int64]*ProfileBadge `json:"badges"`
	IsVerified         bool                    `json:"isVerified"`
}

type UserProfileHeader struct {
	UserName        string `json:"username"`
	Name            string `json:"name"`
	LastName        string `json:"lastName"`
	ProfileImageUrl string `json:"profileImageUrl"`
	Points          int    `json:"points"`
}
