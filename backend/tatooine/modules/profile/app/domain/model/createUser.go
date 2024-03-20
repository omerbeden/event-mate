package model

type CreateUserProfileRequest struct {
	UserName        string            `json:"username"`
	Name            string            `json:"name"`
	LastName        string            `json:"lastname"`
	About           string            `json:"about"`
	Email           string            `json:"email"`
	ProfileImageUrl string            `json:"profileImageUrl"`
	ExternalId      string            `json:"externalId"`
	IsVerified      bool              `json:"isVerified"`
	Adress          UserProfileAdress `json:"address"`
}
