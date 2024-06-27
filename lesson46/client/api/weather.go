package api

import (
	"context"
	"fmt"
	"lesson46/server/genproto/WeatherService/weather"
	"log"
)

type weatherService struct {
}

func NewWeatherService() *weatherService {
	return &weatherService{}
}

func (w *weatherService) GetCurrentWeather(client weather.WeatherServiceClient) {
	req := &weather.WeatherRequest{
		Location: "Toshkent",
	}

	res, err := client.GetCurrentWeather(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("GetCurrentWeather: ", res)
	fmt.Println()
}

func (w *weatherService) GetWeatherForecast(client weather.WeatherServiceClient) {
	req := &weather.ForecastRequest{
		Location: "Toshkent",
		Days:     3,
	}

	res, err := client.GetWeatherForecast(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("GetWeatherForecast:", res)
	fmt.Println()
}

func (w *weatherService) ReportWeatherCondition(client weather.WeatherServiceClient) {
	req := &weather.WeatherCondition{
		Location:    "Toshkent",
		Temperature: 39.0,
		Humidity:    80.0,
	}

	res, err := client.ReportWeatherCondition(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("ReportWeatherCondition: ", res)
	fmt.Println()
}
