package model

import "time"

type UserEvaluation struct {
	GiverId           int64   `json:"giverId"`
	ReceiverId        int64   `json:"receiverId"`
	Points            float32 `json:"points"`
	Comment           string  `json:"comment"`
	RelatedActivityId int64   `json:"relatedActivityId"`
}

type GetUserEvaluations struct {
	GiverUserName        string    `json:"giverUserName"`
	GiverProfileImageUrl string    `json:"giverProfileImageUrl"`
	GivenAt              time.Time `json:"givenAt"`
	Comment              string    `json:"comment"`
	GivenPoint           float32   `json:"givenPoint"`
}
