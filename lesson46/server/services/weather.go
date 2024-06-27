package serveces

import (
	"context"
	pb "lesson46/server/genproto/WeatherService/weather"
	"lesson46/server/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type WeatherRepo struct {
	db *postgres.WeatherRepo
	pb.UnimplementedWeatherServiceServer
}

func NewWeatherRepo(db *sqlx.DB) *WeatherRepo {
	return &WeatherRepo{
		db: postgres.NewWeatherRepo(db),
	}
}

func (r *WeatherRepo) GetCurrentWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	res, err := r.db.GetCurrentWeather(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r *WeatherRepo) GetWeatherForecast(ctx context.Context, in *pb.ForecastRequest) (*pb.ForecastResponse, error) {
	res, err := r.db.GetWeatherForecast(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r *WeatherRepo) ReportWeatherCondition(ctx context.Context, in *pb.WeatherCondition) (*pb.ReportResponse, error) {
	res, err := r.db.ReportWeatherCondition(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, err
}
