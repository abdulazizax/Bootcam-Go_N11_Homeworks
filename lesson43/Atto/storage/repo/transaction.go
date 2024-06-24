package repo

import (
	m "atto/models"
	"context"
)

type TransactionStorageI interface {
	CreateTransaction(ctx context.Context, transaction m.TransactionRequest) (m.TransactionResponse, error)
	GetTransactionById(ctx context.Context, id string) (m.TransactionResponse, error)
	GetTransactions(ctx context.Context) ([]m.TransactionResponse, error)
	UpdateTransactionById(ctx context.Context, id string, updatedTransaction m.TransactionRequest) (m.TransactionResponse, error)
	DeleteTransactionByID(ctx context.Context, id string) error
}
