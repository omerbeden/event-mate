package presenter

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

type ProfileImageUpdateRequest struct {
	ProfileImageUrl string `json:"profileImageUrl"`
}
type ProfileVerificationUpdateRequest struct {
	IsVerified bool `json:"isVerified"`
}

type GetUserProfileResponse struct {
	Id                 int64                   `json:"id"`
	ProfileHeader      model.UserProfileHeader `json:"profileHeader"`
	About              string                  `json:"about"`
	AttandedActivities []model.Activity        `json:"attandedActivities"`
	Adress             model.UserProfileAdress `json:"address"`
	Stat               model.UserProfileStat   `json:"stats"`
	ExternalId         string                  `json:"externalId"`
	Email              string                  `json:"email"`
	Badges             []model.ProfileBadge    `json:"badges"`
	IsVerified         bool                    `json:"isVerified"`
}

func ProfileToGetUserResponse(profile model.UserProfile) *GetUserProfileResponse {

	badgesSlice := ProfileBadgeMapToSlice(profile.Badges)

	return &GetUserProfileResponse{
		Id: profile.Id,
		ProfileHeader: model.UserProfileHeader{
			Name:            profile.Header.Name,
			LastName:        profile.Header.LastName,
			ProfileImageUrl: profile.Header.ProfileImageUrl,
			UserName:        profile.Header.UserName,
		},
		About:              profile.About,
		AttandedActivities: profile.AttandedActivities,
		Adress:             profile.Adress,
		Stat:               profile.Stat,
		ExternalId:         profile.ExternalId,
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
