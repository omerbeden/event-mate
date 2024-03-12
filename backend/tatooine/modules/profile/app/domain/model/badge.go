package model

import "time"

type ProfileBadge struct {
	BadgeId   int64     `json:"badgeId"`
	ProfileId int64     `json:"profileId"`
	ImageUrl  string    `json:"imageUrl"`
	Text      string    `json:"text"`
	GivenAt   time.Time `json:"givenOn"`
}

const (
	TrustworthyBadgeId = int64(1)
	ActiveBadgeId      = int64(2)
)

func TrustworthyBadge() *ProfileBadge {
	return &ProfileBadge{
		BadgeId:  1,
		ImageUrl: "https://i.ibb.co/z6z4z4z/trust-badge.png",
		Text:     "Trustworthy",
	}
}

func ActiveBadge() *ProfileBadge {
	return &ProfileBadge{
		BadgeId:  2,
		ImageUrl: "https://i.ibb.co/z6z4z4z/trust-badge.png",
		Text:     "Active",
	}
}
