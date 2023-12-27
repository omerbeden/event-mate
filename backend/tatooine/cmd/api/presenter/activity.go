package presenter

const APIVersion = "1.0"

type BaseResponse struct {
	APIVersion string `json:"apiVersion"`
	Data       any    `json:"data"`
	Error      string `json:"error"`
}
