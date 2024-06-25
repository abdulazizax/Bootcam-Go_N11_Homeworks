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
	query := `
		INSERT INTO transactions (card_id, amount, terminal_id, transaction_type)
		VALUES ($1, $2, $3, $4)
		RETURNING
			transaction_id
			card_id
			amount
			terminal_id
			transaction_type
			transaction_time
	`

	var ResTransaction m.TransactionResponse

	row := t.DB.QueryRowxContext(ctx, query, transaction.Card_ID, transaction.Amount, transaction.Terminal_ID, transaction.Transaction_Type)
	err := row.Scan(
		&ResTransaction.Transaction_ID,
		&ResTransaction.Card_ID,
		&ResTransaction.Amount,
		&ResTransaction.Terminal_ID,
		&ResTransaction.Transaction_Type,
		&ResTransaction.Transaction_Time,
	)
	if err != nil {
		return m.TransactionResponse{}, err
	}

	return ResTransaction, nil
}

func (t *TransactionRepo) GetTransactionById(ctx context.Context, id string) (m.TransactionResponse, error) {
	query := `
		SELECT 
			transaction_id, 
			card_id, 
			amount, 
			terminal_id, 
			transaction_type, 
			transaction_time 
		FROM transactions 
		WHERE transaction_id = $1
	`
	var ResTransaction m.TransactionResponse

	row := t.DB.QueryRowxContext(ctx, query, id)
	err := row.Scan(
		&ResTransaction.Transaction_ID,
		&ResTransaction.Card_ID,
		&ResTransaction.Amount,
		&ResTransaction.Terminal_ID,
		&ResTransaction.Transaction_Type,
		&ResTransaction.Transaction_Time,
	)
	if err != nil {
		return m.TransactionResponse{}, err
	}

	return ResTransaction, nil
}

func (t *TransactionRepo) GetTransactionByCardId(ctx context.Context, id string) ([]m.TransactionResponse, error) {
	query := `
		SELECT 
			transaction_id, 
			card_id, 
			amount, 
			terminal_id, 
			transaction_type, 
			transaction_time 
		FROM transactions 
		WHERE card_id = $1
	`
	var transactions []m.TransactionResponse

	rows, err := t.DB.QueryxContext(ctx, query)
	if err != nil {
		return []m.TransactionResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var ResTransaction m.TransactionResponse
		if err := rows.Scan(
			&ResTransaction.Transaction_ID,
			&ResTransaction.Card_ID,
			&ResTransaction.Amount,
			&ResTransaction.Terminal_ID,
			&ResTransaction.Transaction_Type,
			&ResTransaction.Transaction_Time,
		); err != nil {
			return []m.TransactionResponse{}, err
		}
		transactions = append(transactions, ResTransaction)
	}

	if err := rows.Err(); err != nil {
		return []m.TransactionResponse{}, err
	}

	return transactions, nil
}

func (t *TransactionRepo) GetTransactions(ctx context.Context) ([]m.TransactionResponse, error) {
	query := `
		SELECT 
			transaction_id, 
			card_id, 
			amount, 
			terminal_id, 
			transaction_type, 
			transaction_time 
		FROM transactions
	`
	var transactions []m.TransactionResponse

	rows, err := t.DB.QueryxContext(ctx, query)
	if err != nil {
		return []m.TransactionResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var ResTransaction m.TransactionResponse
		if err := rows.Scan(
			&ResTransaction.Transaction_ID,
			&ResTransaction.Card_ID,
			&ResTransaction.Amount,
			&ResTransaction.Terminal_ID,
			&ResTransaction.Transaction_Type,
			&ResTransaction.Transaction_Time,
		); err != nil {
			return []m.TransactionResponse{}, err
		}
		transactions = append(transactions, ResTransaction)
	}

	if err := rows.Err(); err != nil {
		return []m.TransactionResponse{}, err
	}

	return transactions, nil
}

func (t *TransactionRepo) UpdateTransactionById(ctx context.Context, id string, updatedTransaction m.TransactionRequest) (m.TransactionResponse, error) {
	query := `
		UPDATE transactions
		SET 
			card_id = $1, 
			amount = $2, 
			terminal_id = $3, 
			transaction_type = $4
		WHERE transaction_id = $5
		RETURNING 
			transaction_id, 
			card_id, 
			amount, 
			terminal_id, 
			transaction_type, 
			transaction_time
	`

	var ResTransaction m.TransactionResponse

	row := t.DB.QueryRowxContext(ctx, query, id, updatedTransaction.Card_ID, updatedTransaction.Amount, updatedTransaction.Terminal_ID, updatedTransaction.Transaction_Type)
	err := row.Scan(
		&ResTransaction.Transaction_ID,
		&ResTransaction.Card_ID,
		&ResTransaction.Amount,
		&ResTransaction.Terminal_ID,
		&ResTransaction.Transaction_Type,
		&ResTransaction.Transaction_Time,
	)
	if err != nil {
		return m.TransactionResponse{}, err
	}

	return ResTransaction, nil
}

func (t *TransactionRepo) DeleteTransactionByID(ctx context.Context, id string) error {
	query := `
		DELETE FROM transactions 
		WHERE transaction_id = $1
	`
	_, err := t.DB.ExecContext(ctx, query, id)
	return err
}
