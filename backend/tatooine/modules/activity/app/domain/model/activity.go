package model

import "time"

type Activity struct {
	ID                 int64     `json:"id"`
	Title              string    `json:"title"`
	Category           string    `json:"category"`
	CreatedBy          User      `json:"createdBy"`
	Location           Location  `json:"location"`
	Participants       []User    `json:"participants"`
	BackgroundImageUrl string    `json:"backgroundImage"`
	StartAt            time.Time `json:"startAt"`
	EndAt              time.Time `json:"endAt"`
	Content            string    `json:"content"`
	Rules              []string  `json:"rules"`
	Flow               []string  `json:"flow"`
	Quota              int       `json:"quota"`
}

type Location struct {
	ActivityId  int64   `json:"-"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	Description string  `json:"description"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}
