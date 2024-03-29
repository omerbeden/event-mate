package presenter

const APIVersion = "1.0"

type BaseResponse struct {
	APIVersion string `json:"apiVersion"`
	Data       any    `json:"data"`
	Error      string `json:"error"`
}

type CreateActivityRequest struct {
	Title               string         `json:"title"`
	Category            string         `json:"category"`
	CreatedById         int64          `json:"createdById"`
	CreatedByExternalId string         `json:"createdByExternalId"`
	Location            LocationDetail `json:"location"`
	StartAt             string         `json:"startAt"`
	EndAt               string         `json:"endAt"`
	Content             string         `json:"content"`
	Rules               []string       `json:"rules"`
	Flow                []string       `json:"flow"`
	Quota               int            `json:"quota"`
	GenderComposition   string         `json:"genderComposition"`
}

type Location struct {
	City     string `json:"city"`
	District string `json:"district"`
}
type LocationDetail struct {
	Location    Location `json:"location"`
	Description string   `json:"description"`
	Latitude    float32  `json:"latitude"`
	Longitude   float32  `json:"longitude"`
}
