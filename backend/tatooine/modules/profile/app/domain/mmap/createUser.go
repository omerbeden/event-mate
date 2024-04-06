package mmap

import "github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"

func CreateUserRequestToProfile(from *model.CreateUserProfileRequest) *model.UserProfile {

	badges := make(map[int64]*model.ProfileBadge)

	if from.IsVerified {
		badges[model.VerifiedBadgeId] = model.VerifiedBadge()
	}

	return &model.UserProfile{
		Header: model.UserProfileHeader{
			UserName:        from.UserName,
			Name:            from.Name,
			LastName:        from.LastName,
			ProfileImageUrl: from.ProfileImageUrl,
			Points:          0,
		},
		About:      from.About,
		Adress:     from.Adress,
		ExternalId: from.ExternalId,
		Email:      from.Email,
		Badges:     badges,
		IsVerified: from.IsVerified,
	}
}
