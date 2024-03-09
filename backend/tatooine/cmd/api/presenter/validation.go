package presenter

type MernisRequest struct {
	NationalId string `json:"nationalId"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	BirthYear  int    `json:"birthYear"`
}
