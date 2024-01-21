package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
)

type userProfileRepo struct {
	pool *pgxpool.Pool
}

const errlogprefix = "repo:userProfile"

func NewUserProfileRepo(pool *pgxpool.Pool) *userProfileRepo {
	return &userProfileRepo{
		pool: pool,
	}
}

func (r *userProfileRepo) GetUsersByAddress(address model.UserProfileAdress) ([]model.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `
	Select p.id, p.name, p.last_name, p.profile_image_url, 
	(stats.point/stats.point_giving.count) as point,
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
		err := rows.Scan(&res.Id, &res.Name, &res.LastName, &res.ProfileImageUrl,
			&res.Stat.Point,
			&res.Adress.City)
		if err != nil {
			return nil, fmt.Errorf("%s err getting rows %w ", errlogprefix, err)
		}
		users = append(users, res)
	}
	return users, nil
}
func (r *userProfileRepo) InsertUser(user *model.UserProfile) (*model.UserProfile, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO user_profiles
	 (name,last_name,profile_image_url,about,external_id)
	 Values($1,$2,$3,$4,$5) RETURNING id`
	var id int64

	errQR := r.pool.QueryRow(ctx, q, user.Name, user.LastName, user.ProfileImageUrl, user.About, user.ExternalId).Scan(&id)
	if errQR != nil {
		return nil, fmt.Errorf("%s could not create %w", errlogprefix, errQR)
	}

	user.Id = id
	errAdress := r.insertProfileAdress(user)
	if errAdress != nil {
		return nil, errAdress
	}
	errStat := r.insertProfileStat(user)
	if errStat != nil {
		return nil, errStat
	}

	user.Stat.ProfileId = user.Id
	user.Adress.ProfileId = user.Id
	return user, nil
}
func (r *userProfileRepo) UpdateProfileImage(externalId string, imageUrl string) (*model.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE user_profiles 
		SET profile_image_url = $1
		WHERE external_id = $2`

	_, err := r.pool.Exec(ctx, q, imageUrl, externalId)
	if err != nil {
		return nil, fmt.Errorf("%s could not update user %s , %w", errlogprefix, externalId, err)
	}

	updatedUser, err := r.GetUserProfile(externalId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get updated user %d , %w", errlogprefix, updatedUser.Id, err)
	}

	updatedUser.AttandedActivities, err = r.GetAttandedActivities(updatedUser.Id)
	if err != nil {
		return nil, fmt.Errorf("%s could not get updated user's attanded activities %d , %w", errlogprefix, updatedUser.Id, err)
	}
	fmt.Printf("updated user :%+v\n", updatedUser)

	updatedUser.Stat.ProfileId = updatedUser.Id
	updatedUser.Adress.ProfileId = updatedUser.Id
	return updatedUser, nil
}
func (r *userProfileRepo) DeleteUserById(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `DELETE FROM user_profiles  WHERE id = $1`
	_, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s could not delete user, id: %d %w", errlogprefix, id, err)
	}

	return nil
}

func (r *userProfileRepo) GetAttandedActivities(userId int64) ([]model.Activity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT a.id , a.title, a.category, a.background_image_url, a.content , a.start_at,
	loc.city
	FROM participants attended
	JOIN user_profiles p ON p.id = attended.user_id
	JOIN activities a ON a.id = attended.activity_id
	JOIN activity_locations loc ON loc.activity_id = a.id
	JOIN user_profiles created ON created.id = a.created_by
	WHERE attended.user_id = $1
	`

	rows, err := r.pool.Query(ctx, q, userId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get activities for user: %d", errlogprefix, userId)
	}

	var activities []model.Activity
	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity.ID, &activity.Title, &activity.Category, &activity.BackgroundImageUrl, &activity.Content, &activity.StartAt,
			&activity.Location.City)
		if err != nil {
			return nil, fmt.Errorf("%s error getting rows for user : %d", errlogprefix, userId)
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

func (r *userProfileRepo) insertProfileAdress(user *model.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO user_profile_addresses (profile_id,city) Values($1,$2)`
	_, err := r.pool.Exec(ctx, q, user.Id, user.Adress.City)
	if err != nil {
		return fmt.Errorf("%s could not insert profile adress %w", errlogprefix, err)
	}

	return nil
}

func (r *userProfileRepo) insertProfileStat(user *model.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := fmt.Sprintf(
		`INSERT INTO user_profile_stats
		 (profile_id, point, attanded_activities)
		 Values(%d,%f,%d)`, user.Id, user.Stat.Point, user.Stat.AttandedActivities)
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return fmt.Errorf("%s could not insert profile stats %w", errlogprefix, err)
	}

	return nil

}

// dont need to anymore
func (r *userProfileRepo) GetUserProfileStats(userId int64) (*model.UserProfileStat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT profile_id, (point/point_giving_count) as point , attanded_activities FROM user_profile_stats
	WHERE profile_id = $1`

	var stat model.UserProfileStat
	err := r.pool.QueryRow(ctx, q, userId).Scan(&stat.ProfileId, &stat.Point, &stat.AttandedActivities)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile stats for : %d %w", errlogprefix, userId, err)

	}

	return &stat, nil
}

func (r *userProfileRepo) GetUserProfile(externalId string) (*model.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `SELECT up.id, up.name, up.last_name, up.about, up.profile_image_url,
    upa.city,
    ups.attanded_activities, 
		CASE 
			WHEN ups.point_giving_count = 0 THEN ups.point
			ELSE (ups.point / ups.point_giving_count) 
		END as point
	FROM user_profiles up
	JOIN user_profile_stats ups ON ups.profile_id = up.id
	JOIN user_profile_addresses upa ON upa.profile_id = up.id
	WHERE up.external_id = $1;`

	var user model.UserProfile
	err := r.pool.QueryRow(ctx, q, externalId).Scan(&user.Id, &user.Name, &user.LastName, &user.About, &user.ProfileImageUrl,
		&user.Adress.City,
		&user.Stat.AttandedActivities, &user.Stat.Point)
	if err != nil {
		return nil, fmt.Errorf("%s could not get user profile for : %s %w", errlogprefix, externalId, err)

	}

	user.AttandedActivities, err = r.GetAttandedActivities(user.Id)
	if err != nil {
		return nil, fmt.Errorf("%s could not get attanded activities for profile: %d %w", errlogprefix, user.Id, err)

	}

	return &user, nil
}

func (r *userProfileRepo) UpdateProfilePoints(userId int64, point float32) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE user_profile_stats
		SET point = point + $1,
		point_giving_count = point_giving_count + 1
		WHERE profile_id = $2`

	_, err := r.pool.Exec(ctx, q, point, userId)
	if err != nil {
		return fmt.Errorf("%s could not update user %d , %w", errlogprefix, userId, err)
	}

	return nil
}
