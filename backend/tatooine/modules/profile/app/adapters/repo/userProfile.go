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
	stats.point, stats.followings, stats.followers,
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
			&res.Stat.Point, &res.Stat.Followings, &res.Stat.Followers,
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
	 (name,last_name,profile_image_url,about)
	 Values($1,$2,$3,$4) RETURNING id`
	var id int64

	errQR := r.pool.QueryRow(ctx, q, user.Name, user.LastName, user.ProfileImageUrl, user.About).Scan(&id)
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
func (r *userProfileRepo) UpdateProfileImage(userId int64, imageUrl string) (*model.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `UPDATE user_profiles 
		SET profile_image_url = $1
		WHERE id = $2`

	_, err := r.pool.Exec(ctx, q, imageUrl, userId)
	if err != nil {
		return nil, fmt.Errorf("%s could not update user %d , %w", errlogprefix, userId, err)
	}

	updatedUser, err := r.getUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("%s could not get updated user %d , %w", errlogprefix, userId, err)
	}

	updatedUser.AttandedActivities, err = r.GetAttandedActivities(updatedUser.Id)
	if err != nil {
		return nil, fmt.Errorf("%s could not get updated user's attanded activities %d , %w", errlogprefix, userId, err)
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
		 (profile_id,followers,followings,point)
		 Values(%d,%d,%d,%f)`, user.Id, user.Stat.Followers, user.Stat.Followings, user.Stat.Point)
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return fmt.Errorf("%s could not insert profile stats %w", errlogprefix, err)
	}

	return nil

}

func (r *userProfileRepo) getUserById(userId int64) (*model.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	q := `
	Select p.id, p.name, p.last_name, p.profile_image_url, p.about,
	stats.point, stats.followings, stats.followers,
	a.city
	FROM user_profile_addresses a
	JOIN user_profiles p ON p.id = a.profile_id
	JOIN user_profile_stats stats ON stats.profile_id = p.id
	WHERE p.id = $1
	`

	var user model.UserProfile
	err := r.pool.QueryRow(ctx, q, userId).Scan(&user.Id, &user.Name, &user.LastName, &user.ProfileImageUrl, &user.About,
		&user.Stat.Point, &user.Stat.Followings, &user.Stat.Followers,
		&user.Adress.City)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
