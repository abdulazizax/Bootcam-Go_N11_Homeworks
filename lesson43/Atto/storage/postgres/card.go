package postgres

import (
	m "atto/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type CardRepo struct {
	DB *sqlx.DB
}

func NewCardRepo(db *sqlx.DB) *CardRepo {
	return &CardRepo{
		DB: db,
	}
}

func (c *CardRepo) CreateCard(ctx context.Context, card m.CardRequest) (m.CardResponse, error) {
	query := `
		INSERT INTO cards (card_number, user_id)
		VALUES ($1, $2)
		RETURNING
			card_id,
			card_number,
			user_id,
			created_at,
			updated_at
	`

	var ResCard m.CardResponse

	if err := c.DB.QueryRowxContext(ctx, query, card.Card_Number, card.User_id).Scan(
		&ResCard.Card_ID,
		&ResCard.Card_Number,
		&ResCard.User_id,
		&ResCard.Created_at,
		&ResCard.Updated_at,
	); err != nil {
		return m.CardResponse{}, err
	}

	return ResCard, nil
}

func (c *CardRepo) GetCardById(ctx context.Context, id string) (m.CardResponse, error) {
	query := `
		SELECT
			card_id,
			card_number,
			user_id,
			created_at,
			updated_at
		FROM cards
		WHERE card_id = $1
	`

	var ResCard m.CardResponse

	if err := c.DB.QueryRowxContext(ctx, query, id).Scan(
		&ResCard.Card_ID,
		&ResCard.Card_Number,
		&ResCard.User_id,
		&ResCard.Created_at,
		&ResCard.Updated_at,
	); err != nil {
		return m.CardResponse{}, err
	}

	return ResCard, nil
}

func (c *CardRepo) GetCards(ctx context.Context) ([]m.CardResponse, error) {
	query := `
		SELECT
			card_id,
			card_number,
			user_id,
			created_at,
			updated_at
		FROM cards
	`

	var cards []m.CardResponse

	rows, err := c.DB.QueryxContext(ctx, query)
	if err != nil {
		return []m.CardResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var ResCard m.CardResponse
		if err := rows.Scan(
			&ResCard.Card_ID,
			&ResCard.Card_Number,
			&ResCard.User_id,
			&ResCard.Created_at,
			&ResCard.Updated_at,
		); err != nil {
			return []m.CardResponse{}, err
		}
		cards = append(cards, ResCard)
	}

	if err = rows.Err(); err != nil {
		return []m.CardResponse{}, err
	}

	return cards, nil
}

func (c *CardRepo) UpdateCardById(ctx context.Context, id string, updatedCard m.CardRequest) (m.CardResponse, error) {
	query := `
		UPDATE cards
		SET
			card_number = $1,
			user_id = $2,
			updated_at = NOW()
		WHERE card_id = $3
		RETURNING
			card_id,
			card_number,
			user_id,
			created_at,
			updated_at
	`

	var ResCard m.CardResponse

	err := c.DB.QueryRowxContext(ctx, query, updatedCard.Card_Number, updatedCard.User_id, id).Scan(
		&ResCard.Card_ID,
		&ResCard.Card_Number,
		&ResCard.User_id,
		&ResCard.Created_at,
		&ResCard.Updated_at,
	)
	if err != nil {
		return m.CardResponse{}, err
	}

	return ResCard, nil
}

func (c *CardRepo) DeleteCardByID(ctx context.Context, id string) error {
	query := `
		DELETE FROM cards
		WHERE card_id = $1
	`

	_, err := c.DB.ExecContext(ctx, query, id)
	return err
}
