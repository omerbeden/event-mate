package postgresadapter

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type userProfileRepo struct {
	pool db.Executor
}

const errlogprefix = "repo:userProfile"

func NewUserProfileRepo(pool db.Executor) *userProfileRepo {
	return &userProfileRepo{
		pool: pool,
	}
}

func (r *userProfileRepo) Insert(ctx context.Context, tx db.Tx, user *model.UserProfile) (*model.UserProfile, error) {

	q := `INSERT INTO user_profiles
	 (name,last_name,profile_image_url,about,external_id,user_name,email,is_verified)
	 Values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`
	var id int64

	errQR := tx.QueryRow(ctx, q, user.Header.Name, user.Header.LastName, user.Header.ProfileImageUrl, user.About, user.ExternalId, user.Header.UserName, user.Email, user.IsVerified).Scan(&id)
	if errQR != nil {
		return nil, fmt.Errorf("%s could not create %w", errlogprefix, errQR)
	}

	user.Id = id

	return user, nil
}

func (r *userProfileRepo) GetUsersByAddress(ctx context.Context, address model.UserProfileAdress) ([]model.UserProfile, error) {

	q := `
	Select p.id, p.name, p.last_name, p.profile_image_url, email,
	(SELECT ROUND(AVG(points),1) FROM user_points WHERE receiver_id = p.external_id) as point,
	a.city
	FROM user_profile_addresses a
	JOIN user_profiles p ON p.id = a.profile_id
	JOIN user_profile_stats stats ON stats.profile_id = p.id
	WHERE a.city = $1
	`

	rows, err := r.pool.Query(ctx, q, address.City)
	if err != nil {
		return nil, fmt.Errorf("%s could not get users, %w", errlogprefix, err)
	}

	var users []model.UserProfile
	for rows.Next() {
		var res model.UserProfile
		err := rows.Scan(&res.Id, &res.Header.Name, &res.Header.LastName, &res.Header.ProfileImageUrl, &res.Email,
			&res.Stat.Point,
			&res.Adress.City)
		if err != nil {
			return nil, fmt.Errorf("%s err getting rows %w ", errlogprefix, err)
		}
		users = append(users, res)
	}
	return users, nil
}

func (r *userProfileRepo) UpdateProfileImage(ctx context.Context, externalId string, imageUrl string) error {

	q := `UPDATE user_profiles 
		SET profile_image_url = $1
		WHERE external_id = $2`

	_, err := r.pool.Exec(ctx, q, imageUrl, externalId)
	if err != nil {
		return fmt.Errorf("%s could not update user %s , %w", errlogprefix, externalId, err)
	}

	return err
}
func (r *userProfileRepo) DeleteUser(ctx context.Context, externalId string) error {

	q := `DELETE FROM user_profiles  WHERE external_id = $1`
	_, err := r.pool.Exec(ctx, q, externalId)
	if err != nil {
		return fmt.Errorf("%s could not delete user, id: %s %w", errlogprefix, externalId, err)
	}

	return nil
}

func (r *userProfileRepo) GetAttandedActivities(ctx context.Context, userId int64) ([]model.Activity, error) {

	q := `SELECT a.id , a.title, a.category, a.content , a.start_at, a.gender_composition,a.quota,a.participant_count,
	p.name, p.last_name, p.user_name, p.profile_image_url, ups.average_point,
	loc.city
	FROM participants attended
	JOIN user_profiles p ON p.id = attended.user_id
	JOIN activities a ON a.id = attended.activity_id
	JOIN activity_locations loc ON loc.activity_id = a.id
	JOIN user_profiles created ON created.id = a.created_by
	JOIN user_profile_stats ups ON ups.profile_id = p.id
	WHERE attended.user_id = $1
	`

	rows, err := r.pool.Query(ctx, q, userId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activities for user: %d %w", errlogprefix, userId, err)
	}

	var activities []model.Activity
	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity.ID, &activity.Title, &activity.Category, &activity.Content, &activity.StartAt, &activity.GenderComposition, &activity.Quota, &activity.ParticipantCount,
			&activity.CreatedBy.Name, &activity.CreatedBy.LastName, &activity.CreatedBy.UserName, &activity.CreatedBy.ProfileImageUrl, &activity.CreatedBy.Points,
			&activity.Location.City)
		if err != nil {
			return nil, fmt.Errorf("%s error getting rows for user : %d %w", errlogprefix, userId, err)
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

func (r *userProfileRepo) GetCreatedActivities(ctx context.Context, userId int64) ([]model.Activity, error) {
	q := `SELECT a.id , a.title, a.category, a.content , a.start_at, a.gender_composition,a.quota,a.participant_count,
	p.name, p.last_name, p.user_name, p.profile_image_url, ups.average_point,
	loc.city
	FROM activities a
	JOIN user_profiles p ON p.id = a.created_by
	JOIN activity_locations loc ON loc.activity_id = a.id
	JOIN user_profile_stats ups ON ups.profile_id = p.id
	WHERE a.created_by = $1`

	rows, err := r.pool.Query(ctx, q, userId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activities for user: %d %w", errlogprefix, userId, err)
	}

	var activities []model.Activity
	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity.ID, &activity.Title, &activity.Category, &activity.Content, &activity.StartAt, &activity.GenderComposition, &activity.Quota, &activity.ParticipantCount,
			&activity.CreatedBy.Name, &activity.CreatedBy.LastName, &activity.CreatedBy.UserName, &activity.CreatedBy.ProfileImageUrl, &activity.CreatedBy.Points,
			&activity.Location.City)
		if err != nil {
			return nil, fmt.Errorf("%s error getting rows for user : %d %w", errlogprefix, userId, err)
		}
		activities = append(activities, activity)
	}

	return activities, nil

}

func (r *userProfileRepo) GetCurrentUserProfile(ctx context.Context, externalId string) (*model.UserProfile, error) {

	q := `SELECT up.id, up.name, up.last_name, up.about, up.profile_image_url, up.external_id, up.user_name,email,up.is_verified,
    upa.city,
    ups.attanded_activities, 
	ups.average_point as point
	FROM user_profiles up
	JOIN user_profile_stats ups ON ups.profile_id = up.id
	JOIN user_profile_addresses upa ON upa.profile_id = up.id
	LEFT JOIN user_points ON up.external_id = receiver_id
	WHERE up.external_id = $1
	`

	var user model.UserProfile
	err := r.pool.QueryRow(ctx, q, externalId).Scan(&user.Id, &user.Header.Name, &user.Header.LastName, &user.About, &user.Header.ProfileImageUrl, &user.ExternalId, &user.Header.UserName, &user.Email, &user.IsVerified,
		&user.Adress.City,
		&user.Stat.AttandedActivities, &user.Stat.Point)
	if err != nil {
		if err == pgx.ErrNoRows {
			//fmt.Printf("user not found with externalId: %s", externalId)
			return nil, customerrors.ERR_NOT_FOUND
		}
		return nil, fmt.Errorf("%s could not get user profile for : %s %w", errlogprefix, externalId, err)

	}

	user.AttandedActivities, err = r.GetAttandedActivities(ctx, user.Id)
	if err != nil {
		return nil, err

	}

	return &user, nil
}
func (r *userProfileRepo) GetUserProfile(ctx context.Context, username string) (*model.UserProfile, error) {

	q := `SELECT up.id, up.name, up.last_name, up.about, up.profile_image_url,email,
    upa.city,
    ups.attanded_activities, 
	ups.average_point as point
	FROM user_profiles up
	JOIN user_profile_stats ups ON ups.profile_id = up.id
	JOIN user_profile_addresses upa ON upa.profile_id = up.id
	WHERE up.user_name = $1;`

	var user model.UserProfile
	err := r.pool.QueryRow(ctx, q, username).Scan(&user.Id, &user.Header.Name, &user.Header.LastName, &user.About, &user.Header.ProfileImageUrl, &user.Email,
		&user.Adress.City,
		&user.Stat.AttandedActivities, &user.Stat.Point)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile for : %s %w", errlogprefix, username, err)

	}

	user.AttandedActivities, err = r.GetAttandedActivities(ctx, user.Id)
	if err != nil {
		return nil, fmt.Errorf("%s could not get attanded activities for profile: %d %w", errlogprefix, user.Id, err)

	}

	return &user, nil
}

func (r *userProfileRepo) GetUserProfileById(ctx context.Context, id int64) (*model.UserProfile, error) {

	q := `SELECT up.id, up.name, up.last_name, up.about, up.profile_image_url,email,up.external_id,up.user_name,
    upa.city,
    ups.attanded_activities, 
	ups.average_point as point	FROM user_profiles up
	JOIN user_profile_stats ups ON ups.profile_id = up.id
	JOIN user_profile_addresses upa ON upa.profile_id = up.id
	WHERE up.id = $1;`

	var user model.UserProfile
	err := r.pool.QueryRow(ctx, q, id).Scan(&user.Id, &user.Header.Name, &user.Header.LastName, &user.About, &user.Header.ProfileImageUrl, &user.Email, &user.ExternalId, &user.Header.UserName,
		&user.Adress.City,
		&user.Stat.AttandedActivities, &user.Stat.Point)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile for : %d %w", errlogprefix, id, err)

	}

	user.AttandedActivities, err = r.GetAttandedActivities(ctx, user.Id)
	if err != nil {
		return nil, fmt.Errorf("%s could not get attanded activities for profile: %d %w", errlogprefix, user.Id, err)

	}

	return &user, nil
}

func (r *userProfileRepo) UpdateVerification(ctx context.Context, externalId string, isVerified bool) error {

	q := `UPDATE user_profiles 
		SET is_verified = $1
		WHERE external_id = $2`

	_, err := r.pool.Exec(ctx, q, isVerified, externalId)
	if err != nil {
		return fmt.Errorf("%s could not update is_verified  %s , %w", errlogprefix, externalId, err)
	}

	return nil
}

func (r *userProfileRepo) GetId(ctx context.Context, externalId string) (int64, error) {
	q := `SELECT id FROM user_profiles WHERE external_id = $1`

	var id int64
	err := r.pool.QueryRow(ctx, q, externalId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s could not get id , externalId:%s , %w", errlogprefix, externalId, err)
	}

	return id, nil

}
