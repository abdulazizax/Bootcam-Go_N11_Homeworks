package api

import (
	h "atto/api/handler"
	"atto/storage"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Storage storage.IStorage
}

func New(option Option) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := h.New(&h.HandlerConfig{
		Storage: option.Storage,
	})

	crud := router.Group("")
	{
		crud.POST("/card/post", handler.CreateCard)
		crud.POST("/transaction/debit/post", handler.CreateDebitTransaction)
		crud.POST("/transaction/credit/post", handler.CreateCreditTransaction)
		crud.POST("/station/post", handler.CreateStation)

		crud.GET("/card/get", handler.GetCards)
		crud.GET("/transaction/get", handler.GetTransactions)
		crud.GET("/station/get", handler.GetStations)

		crud.GET("/card/get/:id", handler.GetCardByID)
		crud.GET("/transaction/get/:id", handler.GetTransactionByID)
		crud.GET("/station/get/:id", handler.GetStationByID)

		crud.GET("/station/get/name/:name", handler.GetStationByName)

		crud.PUT("/card/put/:id", handler.UpdateCardByID)
		crud.PUT("/transaction/put/:id", handler.UpdateTransactionByID)
		crud.PUT("/station/put/:id", handler.UpdateStationByID)

		crud.DELETE("/card/:id", handler.DeleteCardByID)
		crud.DELETE("/transaction/:id", handler.DeleteTransactionByID)
		crud.DELETE("/station/:id", handler.DeleteStationByID)
	}

	return router
}
