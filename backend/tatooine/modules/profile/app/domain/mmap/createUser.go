package mmap

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

func CreateUserRequestToProfile(from *model.CreateUserProfileRequest) *model.UserProfile {

	badges := make(map[int64]*model.ProfileBadge)

	if from.IsVerified {
		badges[model.VerifiedBadgeId] = model.VerifiedBadge()
	}

	return &model.UserProfile{
		Name:            from.Name,
		LastName:        from.LastName,
		About:           from.About,
		Adress:          from.Adress,
		ProfileImageUrl: from.ProfileImageUrl,
		ExternalId:      from.ExternalId,
		UserName:        from.UserName,
		Email:           from.Email,
		Badges:          badges,
		IsVerified:      from.IsVerified,
	}
}
