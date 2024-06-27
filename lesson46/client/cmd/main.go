package main

import (
	"lesson46/client/api"
	"lesson46/server/genproto/TransportService/transport"
	"lesson46/server/genproto/WeatherService/weather"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	transportClient := transport.NewTransportServiceClient(conn)
	weatherClient := weather.NewWeatherServiceClient(conn)

	t := api.NewTransportService()
	w := api.NewWeatherService()

	t.GetBusSchedule(transportClient)
	t.TrackBusLocation(transportClient)
	t.ReportTrafficJam(transportClient)

	w.GetCurrentWeather(weatherClient)
	w.GetWeatherForecast(weatherClient)
	w.ReportWeatherCondition(weatherClient)
}
