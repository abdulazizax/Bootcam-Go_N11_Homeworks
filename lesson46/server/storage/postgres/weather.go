package postgres

import (
	"context"
	"fmt"
	"time"

	pb "lesson46/server/genproto/WeatherService/weather"

	"github.com/jmoiron/sqlx"
)

type WeatherRepo struct {
	DB *sqlx.DB
}

func NewWeatherRepo(db *sqlx.DB) *WeatherRepo {
	return &WeatherRepo{
		DB: db,
	}
}

func (w *WeatherRepo) Test() bool {
	return true
}

func (w *WeatherRepo) GetCurrentWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	query := `
		SELECT 
			temperature,
			humidity,
			speed_of_wind 
		FROM
			current_weather 
		WHERE
			location = $1
`
	var weather pb.WeatherResponse

	row := w.DB.QueryRowxContext(ctx, query, in.Location)
	err := row.Scan(
		&weather.Temperature,
		&weather.Humidity,
		&weather.SpeedOfWind,
	)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (w *WeatherRepo) GetWeatherForecast(ctx context.Context, in *pb.ForecastRequest) (*pb.ForecastResponse, error) {
	query := `
		SELECT 
			date, 
			temperature_high, 
			temperature_low, 
			condition 
		FROM 
			weather_forecast 
		WHERE 
			location = $1 
		ORDER BY date DESC 
		LIMIT $2;
`
	rows, err := w.DB.QueryxContext(ctx, query, in.Location, in.Days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forecast []*pb.Forecast
	var date time.Time

	for rows.Next() {
		var forecastItem pb.Forecast
		err := rows.Scan(
			&date,
			&forecastItem.TemperatureHigh,
			&forecastItem.TemperatureLow,
			&forecastItem.Condition,
		)
		if err != nil {
			return nil, err
		}
		forecastItem.Date = date.Format("2006-01-02")
		forecast = append(forecast, &forecastItem)
	}

	return &pb.ForecastResponse{
		Forecasts: forecast,
	}, nil
}

func (w *WeatherRepo) ReportWeatherCondition(ctx context.Context, in *pb.WeatherCondition) (*pb.ReportResponse, error) {
	query := `
		INSERT INTO weather_condition_report (
			location, 
			temperature, 
			humidity, 
			condition) 
		VALUES 
			($1, $2, $3, $4);
`

	_, err := w.DB.ExecContext(ctx, query, in.Location, in.Temperature, in.Humidity, in.Condition)
	if err != nil {
		return nil, fmt.Errorf("failed to insert weather condition report: %v", err)
	}

	return &pb.ReportResponse{Status: "Report received successfully"}, nil
}
