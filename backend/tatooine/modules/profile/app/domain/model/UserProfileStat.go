package model

type UserProfileStat struct {
	ProfileId          int64   `json:"-"`
	AttandedActivities int     `json:"attandedActivities"`
	Point              float32 `json:"point"`
}
