package handler

import (
	"api-gateway/genproto/WeatherService/weather"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetCurrentWeather(c *gin.Context) {
	var req weather.WeatherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.weather.GetCurrentWeather(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) GetWeatherForecast(c *gin.Context) {
	var req weather.ForecastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.weather.GetWeatherForecast(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) ReportWeatherCondition(c *gin.Context) {
	var req weather.WeatherCondition
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.weather.ReportWeatherCondition(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}