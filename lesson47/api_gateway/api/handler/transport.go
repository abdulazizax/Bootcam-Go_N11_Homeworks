package handler

import (
	"api-gateway/genproto/TransportService/transport"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBusSchedule(c *gin.Context) {
	var req transport.ScheduleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.transport.GetBusSchedule(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) TrackBusLocation(c *gin.Context) {
	var req transport.LocationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.transport.TrackBusLocation(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) ReportTrafficJam(c *gin.Context) {
	var req transport.TrafficJamReport

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.transport.ReportTrafficJam(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
