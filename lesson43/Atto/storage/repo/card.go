package repo

import (
	m "atto/models"
	"context"
)

type CardStorageI interface {
	CreateCard(ctx context.Context, station m.CardRequest) (m.CardResponse, error)
	GetCardById(ctx context.Context, id string) (m.CardResponse, error)
	GetCards(ctx context.Context) ([]m.CardResponse, error)
	UpdateCardById(ctx context.Context, id string, updatedCard m.CardRequest) (m.CardResponse, error)
	DeleteCardByID(ctx context.Context, id string) error
}
