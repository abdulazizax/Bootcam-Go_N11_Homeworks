package repo

import (
	m "atto/models"
	"context"
)

type StationStorageI interface {
	CreateStation(ctx context.Context, station m.StationRequest) (m.StationResponse, error)
	GetStationById(ctx context.Context, id string) (m.StationResponse, error)
	GetStations(ctx context.Context) ([]m.StationResponse, error)
	UpdateStationById(ctx context.Context, id string, updatedStation m.StationRequest) (m.StationResponse, error)
	DeleteStationByID(ctx context.Context, id string) error
	GetStationByName(ctx context.Context, name string) (m.StationResponse, error)
}
