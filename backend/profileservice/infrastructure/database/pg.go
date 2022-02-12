package database

import (
	"database/sql"
	"time"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func NewConn() {
	dsn := "postgres://postgres:@localhost:5432/test" //test
	db, err := sql.Open("pg", dsn)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 10) // test
}
