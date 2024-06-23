package storage

import (
	"user/storage/postgres"
	"user/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	User() repo.UsersStorageI
}

type storagePg struct {
	db         *sqlx.DB
	userRepo   repo.UsersStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		userRepo:   postgres.NewUserrepo(db),
	}
}


func (s *storagePg) User() repo.UsersStorageI {
	return s.userRepo
}
