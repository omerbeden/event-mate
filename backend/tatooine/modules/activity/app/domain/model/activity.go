package model

import "time"

type Activity struct {
	ID                int64             `json:"id"`
	Title             string            `json:"title"`
	Category          string            `json:"category"`
	CreatedBy         User              `json:"createdBy"`
	Location          Location          `json:"location"`
	Participants      []User            `json:"participants,omitempty"`
	StartAt           time.Time         `json:"startAt"`
	EndAt             time.Time         `json:"endAt,omitempty"`
	Content           string            `json:"content"`
	Rules             []string          `json:"rules,omitempty"`
	Flow              []string          `json:"flow,omitempty"`
	Quota             int               `json:"quota"`
	GenderComposition GenderComposition `json:"genderComposition"`
	ParticipantCount  int               `json:"participantCount"`
}

type ActivityDetail struct {
	Participants []User   `json:"participants,omitempty"`
	Rules        []string `json:"rules,omitempty"`
	Flow         []string `json:"flow,omitempty"`
}

type GetActivityCommandResult struct {
	ID                int64             `json:"id"`
	Title             string            `json:"title"`
	Category          string            `json:"category"`
	CreatedBy         User              `json:"createdBy"`
	Location          Location          `json:"location"`
	StartAt           time.Time         `json:"startAt"`
	EndAt             time.Time         `json:"endAt,omitempty"`
	Content           string            `json:"content"`
	Quota             int               `json:"quota"`
	GenderComposition GenderComposition `json:"genderComposition"`
	ParticipantCount  int               `json:"participantCount"`
}
type Location struct {
	ActivityId  int64   `json:"-"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	Description string  `json:"description,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
}

type GenderComposition string

const (
	GenderCompositionOnlyWomen GenderComposition = "only women"
	GenderCompositionOnlyMen   GenderComposition = "only men"
	GenderCompositionMixed     GenderComposition = "mixed"
)
