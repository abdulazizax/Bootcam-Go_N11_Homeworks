package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetDB() (*sqlx.DB, error) {

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"pass",
		"n11",
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	return db, err
}
