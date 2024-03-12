package model

type UserEvaluation struct {
	GiverId           string  `json:"giver_id"`
	ReceiverId        string  `json:"receiver_id"`
	Points            float32 `json:"points"`
	Comment           string  `json:"comment"`
	RelatedActivityId int64   `json:"related_activity_id"`
}
