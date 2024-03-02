package model

import "time"

type Activity struct {
	ID           int64         `json:"-"`
	Title        string        `json:"title"`
	Category     string        `json:"category"`
	CreatedBy    UserProfile   `json:"createdBy"`
	Location     Location      `json:"location"`
	Participants []UserProfile `json:"participants"`
	StartAt      time.Time     `json:"startAt"`
	Content      string        `json:"content"`
}

type Location struct {
	ActivityId int64  `json:"-"`
	City       string `json:"city"`
}
