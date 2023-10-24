package repo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/event/app/domain/model"
)

type LocationRepository struct {
	pool *pgxpool.Pool
}

// TODO: bu tur ayarları bir kere set ediyor olmam lazım daha genel biryere tasi
// daha sonra aslında repolar icin interface e gerek yok
func NewLoc(cnnStr string) *LocationRepository {
	//dbUrl := os.Getenv("Db_Conn_Str")
	config, err := pgxpool.ParseConfig(cnnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse config: %v\n", err)
		os.Exit(1)
	}
	config.MinConns = 5
	config.MaxConns = 10
	//later import db tracer

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse config: %v\n", err)
		os.Exit(1)
	}
	return &LocationRepository{
		pool: pool,
	}
}

func (r *LocationRepository) Create(location model.Location) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO locations (city) Values($1) `
	_, err := r.pool.Exec(ctx, q, location.City)
	if err != nil {
		return false, fmt.Errorf("could not create %w", err)
	}

	return true, nil
}

func (r *LocationRepository) Close() {
	r.pool.Close()
}
