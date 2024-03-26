package presenter

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type ProfileImageUpdateRequest struct {
	ProfileImageUrl string `json:"profileImageUrl"`
}
type ProfileVerificationUpdateRequest struct {
	IsVerified bool `json:"isVerified"`
}

type GetUserProfileResponse struct {
	Id                 int64                   `json:"-"`
	Name               string                  `json:"name"`
	LastName           string                  `json:"lastName"`
	About              string                  `json:"about"`
	AttandedActivities []model.Activity        `json:"attandedActivities"`
	Adress             model.UserProfileAdress `json:"address"`
	Stat               model.UserProfileStat   `json:"stats"`
	ProfileImageUrl    string                  `json:"profileImageUrl"`
	ExternalId         string                  `json:"externalId"`
	UserName           string                  `json:"userName"`
	Email              string                  `json:"email"`
	Badges             []model.ProfileBadge    `json:"badges"`
	IsVerified         bool                    `json:"isVerified"`
}

func ProfileToGetUserResponse(profile model.UserProfile) *GetUserProfileResponse {

	badgesSlice := ProfileBadgeMapToSlice(profile.Badges)

	return &GetUserProfileResponse{
		Id:                 profile.Id,
		Name:               profile.Name,
		LastName:           profile.LastName,
		About:              profile.About,
		AttandedActivities: profile.AttandedActivities,
		Adress:             profile.Adress,
		Stat:               profile.Stat,
		ProfileImageUrl:    profile.ProfileImageUrl,
		ExternalId:         profile.ExternalId,
		UserName:           profile.UserName,
		Email:              profile.Email,
		Badges:             badgesSlice,
		IsVerified:         profile.IsVerified,
	}
}

func ProfileBadgeMapToSlice(badges map[int64]*model.ProfileBadge) []model.ProfileBadge {
	var badgesSlice []model.ProfileBadge

	for _, value := range badges {
		badgesSlice = append(badgesSlice, *value)
	}

	return badgesSlice
}
