package handler

import (
	m "atto/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateDebitTransaction(c *gin.Context) {
	var newTransaction m.TransactionRequest
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.storage.Transaction().DeleteTransactionByID(c, newTransaction)
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
}

func (h *handler) CreateCreditTransaction(c *gin.Context) {

}

func (h *handler) GetTransactions(c *gin.Context) {

}

func (h *handler) GetTransactionByID(c *gin.Context) {

}

func (h *handler) UpdateTransactionByID(c *gin.Context) {

}

func (h *handler) DeleteTransactionByID(c *gin.Context) {

}

func (h *handler) GetStationByName(c *gin.Context) {

}
