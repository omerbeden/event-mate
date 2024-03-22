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
	VerifiedBadgeId    = int64(1)
	TrustworthyBadgeId = int64(2)
	ActiveBadgeId      = int64(3)
)

func VerifiedBadge() *ProfileBadge {
	return &ProfileBadge{
		BadgeId:  VerifiedBadgeId,
		ImageUrl: "https://i.ibb.co/z6z4z4z/trust-badge.png",
		Text:     "Verified",
	}
}

func TrustworthyBadge() *ProfileBadge {
	return &ProfileBadge{
		BadgeId:  TrustworthyBadgeId,
		ImageUrl: "https://i.ibb.co/z6z4z4z/trust-badge.png",
		Text:     "Trustworthy",
	}
}

func ActiveBadge() *ProfileBadge {
	return &ProfileBadge{
		BadgeId:  ActiveBadgeId,
		ImageUrl: "https://i.ibb.co/z6z4z4z/trust-badge.png",
		Text:     "Active",
	}
}
