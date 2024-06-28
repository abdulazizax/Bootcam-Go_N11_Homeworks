package api

import (
	"api-gateway/api/handler"
	"api-gateway/genproto/TransportService/transport"
	"api-gateway/genproto/WeatherService/weather"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Option struct{}

func New(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	transportClient := transport.NewTransportServiceClient(conn)
	weatherClient := weather.NewWeatherServiceClient(conn)

	handler := handler.New(transportClient, weatherClient)

	crud := router.Group("/")
	{
		crud.GET("/transport/bus/schedule", handler.GetBusSchedule)
		crud.GET("/transport/bus/location", handler.TrackBusLocation)
		crud.POST("/transport/traffic/jam", handler.ReportTrafficJam)

		crud.GET("/weather/current", handler.GetCurrentWeather)
		crud.GET("/weather/forecast", handler.GetWeatherForecast)
		crud.POST("/weather/hourly", handler.ReportWeatherCondition)
	}

	return router
}
