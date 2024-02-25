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
	Content            string    `json:"content"`
}

type Location struct {
	ActivityId  int64  `json:"-"`
	City        string `json:"city"`
	District    string `json:"district"`
	Description string `json:"description"`
}
