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
	stats.points, stats.followings, stats.followers,
	a.city
	FROM user_profile_addresses a
	JOIN user_profiles p ON p.id = a.profile_id
	JOIN user_profile_stats stats ON stats.profile_id = p.id
	WHERE a.city = $1
	`

	rows, err := r.pool.Query(ctx, q, address.City)
	if err != nil {
		return nil, fmt.Errorf("could not get users, %w", err)
	}

	var users []model.UserProfile
	for rows.Next() {
		var res model.UserProfile
		err := rows.Scan(&res.Id, &res.Name, &res.LastName, &res.ProfileImageUrl,
			&res.Stat.Points, &res.Stat.Following, &res.Stat.Followers,
			&res.Adress.City)
		if err != nil {
			return nil, fmt.Errorf("err getting rows %w ", err)
		}
		users = append(users, res)
	}
	return users, nil
}
func (r *userProfileRepo) InsertUser(user *model.UserProfile) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO user_profiles
	 (name,last_name,profile_image_url)
	 Values($1,$2,$3) RETURNING id`
	var id int64

	errQR := r.pool.QueryRow(ctx, q, user.Name, user.LastName, user.ProfileImageUrl).Scan(&id)
	if errQR != nil {
		return false, fmt.Errorf("could not create %w", errQR)
	}

	user.Id = id
	errAdress := r.insertProfileAdress(user)
	if errAdress != nil {
		return false, errAdress
	}
	errStat := r.insertProfileStat(user)
	if errStat != nil {
		return false, errStat
	}

	return true, nil
}
func (r *userProfileRepo) UpdateUser(user *model.UserProfile) error {
	return nil
}
func (r *userProfileRepo) DeleteUserById(id uint) (bool, error) {
	return false, nil
}

func (r *userProfileRepo) insertProfileAdress(user *model.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO user_profile_addresses (profile_id,city) Values($1,$2)`
	_, err := r.pool.Exec(ctx, q, user.Id, user.Adress.City)
	if err != nil {
		return fmt.Errorf("could not insert profile adress %w", err)
	}

	return nil
}

func (r *userProfileRepo) insertProfileStat(user *model.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := fmt.Sprintf(
		`INSERT INTO user_profile_stats
		 (profile_id,followers,followings,points)
		 Values(%d,%d,%d,%f)`, user.Id, user.Stat.Followers, user.Stat.Following, user.Stat.Points)
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return fmt.Errorf("could not insert profile stats %w", err)
	}

	return nil

}
