package model

import "time"

type Activity struct {
	ID                int64             `json:"id"`
	Title             string            `json:"title"`
	Category          string            `json:"category"`
	Location          Location          `json:"location"`
	CreatedBy         UserProfileHeader `json:"createdBy"`
	Content           string            `json:"content"`
	StartAt           time.Time         `json:"startAt"`
	EndAt             time.Time         `json:"endAt"`
	GenderComposition string            `json:"genderComposition"`
	Quota             int               `json:"quota"`
	ParticipantCount  int               `json:"participantCount"`
}

type Location struct {
	ActivityId int64  `json:"-"`
	City       string `json:"city"`
	District   string `json:"district"`
}
