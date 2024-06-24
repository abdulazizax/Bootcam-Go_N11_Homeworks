package postgres

import (
	m "atto/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type TransactionRepo struct {
	DB *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{
		DB: db,
	}
}

func (t *TransactionRepo) CreateTransaction(ctx context.Context, transaction m.TransactionRequest) (m.TransactionResponse, error) {
	return m.TransactionResponse{}, nil
}

func (t *TransactionRepo) GetTransactionById(ctx context.Context, id string) (m.TransactionResponse, error) {
	return m.TransactionResponse{}, nil
}

func (t *TransactionRepo) GetTransactions(ctx context.Context) ([]m.TransactionResponse, error) {
	return []m.TransactionResponse{}, nil
}

func (t *TransactionRepo) UpdateTransactionById(ctx context.Context, id string, updatedTransaction m.TransactionRequest) (m.TransactionResponse, error) {
	return m.TransactionResponse{}, nil
}

func (t *TransactionRepo) DeleteTransactionByID(ctx context.Context, id string) error {
	return nil
}
