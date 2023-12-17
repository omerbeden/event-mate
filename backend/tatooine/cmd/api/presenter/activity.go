package presenter

const APIVersion = "1.0"

type BaseResponse struct {
	APIVersion string
	Data       any
	Error      any
}
