package handler

import (
	m "atto/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateStation(c *gin.Context) {
	var newStation m.StationRequest
	if err := c.ShouldBindJSON(&newStation); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.storage.Station().CreateStation(c, newStation)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) GetStations(c *gin.Context) {
	res, err := h.storage.Station().GetStations(c)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(200, res)
}

func (h *handler) GetStationByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.storage.Station().GetStationById(c, id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) UpdateStationByID(c *gin.Context) {
	id := c.Param("id")

	var newStation m.StationRequest
	if err := c.ShouldBindJSON(&newStation); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.storage.Station().UpdateStationById(c, id, newStation)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) DeleteStationByID(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Station().DeleteStationByID(c, id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Station deleted"})
}
