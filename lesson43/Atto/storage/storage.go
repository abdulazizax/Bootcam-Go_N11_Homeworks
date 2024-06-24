package storage

import (
	"atto/storage/repo"
	"atto/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Card() repo.CardStorageI
	Station() repo.StationStorageI
	Transaction() repo.TransactionStorageI
}

type storagePg struct {
	db              *sqlx.DB
	cardRepo        repo.CardStorageI
	stationRepo     repo.StationStorageI
	transactionRepo repo.TransactionStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:              db,
		cardRepo:        postgres.NewCardRepo(db),
		stationRepo:     postgres.NewStationRepo(db),
		transactionRepo: postgres.NewTransactionRepo(db),
	}
}

func (s *storagePg) Card() repo.CardStorageI {
	return s.cardRepo
}

func (s *storagePg) Station() repo.StationStorageI {
	return s.stationRepo
}

func (s *storagePg) Transaction() repo.TransactionStorageI {
	return s.transactionRepo
}
