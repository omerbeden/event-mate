package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/domain/model"
)

type UserProfileRepo struct {
	pool *pgxpool.Pool
}

func (r *UserProfileRepo) GetUsers() ([]model.UserProfile, error) {
	return nil, nil
}
func (r *UserProfileRepo) GetUserById(id uint) (*model.UserProfile, error) {
	return nil, nil
}
func (r *UserProfileRepo) InsertUser(user *model.UserProfile) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO profile (name,lastName,about,photo) Values($1,$2,$3,$4) RETURNING id`
	var id int64
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	errQR := tx.QueryRow(ctx, q, user.Name, user.LastName, user.About, user.Photo).Scan(&id)
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
func (r *UserProfileRepo) UpdateUser(user *model.UserProfile) error {
	return nil
}
func (r *UserProfileRepo) DeleteUserById(id uint) (bool, error) {
	return false, nil
}

func (r *UserProfileRepo) insertProfileAdress(user *model.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO userProfileAdress (city,userProfileId) Values($1,$2)`
	_, err := r.pool.Exec(ctx, q, user.Adress.City, user.Id)
	if err != nil {
		return fmt.Errorf("could not insert profile adress")
	}

	return nil
}

func (r *UserProfileRepo) insertProfileStat(user *model.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := fmt.Sprintf(
		`INSERT INTO userProfileStat
		 (userprofileid,followers,following,points)
		 Values(%d,%d,%d,%f)`, user.Id, user.Stat.Followers, user.Stat.Following, user.Stat.Points)
	_, err := r.pool.Exec(ctx, q)
	if err != nil {
		return fmt.Errorf("could not insert profile adress")
	}

	return nil

}
