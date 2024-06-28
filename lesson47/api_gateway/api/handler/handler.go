package handler

import (
	"api-gateway/genproto/TransportService/transport"
	"api-gateway/genproto/WeatherService/weather"
)

type handler struct {
	// transport.UnimplementedTransportServiceServer
	// weather.UnimplementedWeatherServiceServer
	transport transport.TransportServiceClient
	weather   weather.WeatherServiceClient
}

type HandlerConfig struct {
}

func New(transport transport.TransportServiceClient, weather weather.WeatherServiceClient) *handler {
	return &handler{transport: transport, weather: weather}
}
