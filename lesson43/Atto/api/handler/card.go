package handler

import (
	m "atto/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCard(c *gin.Context) {
	var newCard m.CardRequest
	if err := c.ShouldBindJSON(&newCard); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.storage.Card().CreateCard(c, newCard)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, res)
}

func (h *handler) GetCardByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.storage.Card().GetCardById(c, id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) GetCards(c *gin.Context) {
	res, err := h.storage.Card().GetCards(c)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) UpdateCardByID(c *gin.Context) {
	id := c.Param("id")

	var newCard m.CardRequest

	if err := c.ShouldBindJSON(&newCard); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.storage.Card().UpdateCardById(c, id, newCard)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, res)
}

func (h *handler) DeleteCardByID(c *gin.Context) {
	id := c.Param("id")
	err := h.storage.Card().DeleteCardByID(c, id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Card deleted"})
}
