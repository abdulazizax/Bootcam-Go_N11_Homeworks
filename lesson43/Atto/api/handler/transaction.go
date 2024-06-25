package handler

import (
	m "atto/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTransaction(c *gin.Context) {
	var newTransaction m.TransactionRequest
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.storage.Transaction().CreateTransaction(c, newTransaction)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, res)
}

func (h *handler) GetTransactions(c *gin.Context) {
	res, err := h.storage.Transaction().GetTransactions(c)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.storage.Transaction().GetTransactionById(c, id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) UpdateTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var newTransaction m.TransactionRequest
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.storage.Transaction().UpdateTransactionById(c, id, newTransaction)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}

func (h *handler) DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Transaction().DeleteTransactionByID(c, id)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Transaction deleted"})
}

func (h *handler) GetTransactionsByUserID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.storage.Transaction().GetTransactionByCardId(c, id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, res)
}
