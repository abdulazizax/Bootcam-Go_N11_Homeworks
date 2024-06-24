package postgres

import (
	m "atto/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type StationRepo struct {
	DB *sqlx.DB
}

func NewStationRepo(db *sqlx.DB) *StationRepo {
	return &StationRepo{
		DB: db,
	}
}

func (s *StationRepo) CreateStation(ctx context.Context, station m.StationRequest) (m.StationResponse, error) {
	query := `
		INSERT INTO station (station_name)
		VALUES ($1)
		RETURNING
			station_id,
			station_name,
			created_at,
			updated_at
	`

	var ResStat m.StationResponse

	row := s.DB.QueryRowxContext(ctx, query, station.Station_Name)
	err := row.Scan(&ResStat.Station_ID, &ResStat.Station_Name, &ResStat.Created_at, &ResStat.Updated_at)
	if err != nil {
		return m.StationResponse{}, err
	}

	return ResStat, nil
}

func (s *StationRepo) GetStationById(ctx context.Context, id string) (m.StationResponse, error) {
	query := `
		SELECT
			station_id,
			station_name,
			created_at,
			updated_at
		FROM station
		WHERE station_id = $1
	`

	var ResStat m.StationResponse

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&ResStat.Station_ID,
		&ResStat.Station_Name,
		&ResStat.Created_at,
		&ResStat.Updated_at,
	)
	if err != nil {
		return m.StationResponse{}, err
	}

	return ResStat, nil
}

func (s *StationRepo) GetStations(ctx context.Context) ([]m.StationResponse, error) {
	query := `
		SELECT
			station_id,
			station_name,
			created_at,
			updated_at
		FROM station
	`

	var stations []m.StationResponse

	rows, err := s.DB.QueryxContext(ctx, query)
	if err != nil {
		return []m.StationResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var ResStat m.StationResponse
		if err := rows.Scan(
			&ResStat.Station_ID,
			&ResStat.Station_Name,
			&ResStat.Created_at,
			&ResStat.Updated_at,
		); err != nil {
			return []m.StationResponse{}, err
		}
		stations = append(stations, ResStat)
	}

	if err := rows.Err(); err != nil {
		return []m.StationResponse{}, err
	}

	return stations, nil
}

func (s *StationRepo) UpdateStationById(ctx context.Context, id string, updatedStation m.StationRequest) (m.StationResponse, error) {
	query := `
		UPDATE station
		SET
			station_name = $1,
			updated_at = NOW()
		WHERE station_id = $2
		RETURNING
			station_id,
			station_name,
			created_at,
			updated_at
	`

	var ResStat m.StationResponse

	err := s.DB.QueryRowxContext(ctx, query, updatedStation.Station_Name, id).Scan(
		&ResStat.Station_ID,
		&ResStat.Station_Name,
		&ResStat.Created_at,
		&ResStat.Updated_at,
	)
	if err != nil {
		return m.StationResponse{}, err
	}

	return ResStat, nil
}

func (s *StationRepo) DeleteStationByID(ctx context.Context, id string) error {
	query := `
		DELETE FROM station
		WHERE station_id = $1
	`

	_, err := s.DB.ExecContext(ctx, query, id)
	return err
}

func (s *StationRepo) GetStationByName(ctx context.Context, name string) (m.StationResponse, error) {
	query := `
		SELECT
			station_id,
			station_name,
			created_at,
			updated_at
		FROM station
		WHERE station_name = $1
	`

	var ResStat m.StationResponse

	if err := s.DB.QueryRowxContext(ctx, query, name).Scan(
		&ResStat.Station_ID,
		&ResStat.Station_Name,
		&ResStat.Created_at,
		&ResStat.Updated_at,
	); err != nil {
		return m.StationResponse{}, err
	}

	return ResStat, nil
}
