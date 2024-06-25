package api

import (
	h "api-gateway/api/handler"
	"net/http"
)

func New() *http.ServeMux {
	router := http.NewServeMux()

	handler := h.New(&h.HandlerConfig{})

	router.HandleFunc("POST /user/post", handler.CreateUser)
	router.HandleFunc("POST /card/post", handler.CreateCard)
	router.HandleFunc("POST /station/post", handler.CreateStation)
	router.HandleFunc("POST /transaction/post", handler.CreateTransaction)

	router.HandleFunc("GET /user/get", handler.GetUsers)
	router.HandleFunc("GET /card/get", handler.GetCards)
	router.HandleFunc("GET /station/get", handler.GetStations)
	router.HandleFunc("GET /transaction/get", handler.GetTransactions)

	router.HandleFunc("GET /user/get/{id}", handler.GetUserById)
	router.HandleFunc("GET /card/get/{id}", handler.GetCardById)
	router.HandleFunc("GET /station/get/{id}", handler.GetStationById)
	router.HandleFunc("GET /transaction/get/{id}", handler.GetTransactionById)
	router.HandleFunc("GET /transaction/get/card_id//{id}", handler.GetTransactionByCardId)

	router.HandleFunc("GET /station/get/name/{name}", handler.GetStationByName)

	router.HandleFunc("PUT /user/put/{id}", handler.UpdateUserById)
	router.HandleFunc("PUT /card/put/{id}", handler.UpdateCardById)
	router.HandleFunc("PUT /station/put/{id}", handler.UpdateStationById)
	router.HandleFunc("PUT /transaction/put/{id}", handler.UpdateTransactionById)

	router.HandleFunc("DELETE /user/delete/{id}", handler.DeleteUserByID)
	router.HandleFunc("DELETE /card/delete/{id}", handler.DeleteCardByID)
	router.HandleFunc("DELETE /station/delete/{id}", handler.DeleteStationByID)
	router.HandleFunc("DELETE /transaction/delete/{id}", handler.DeleteTransactionByID)

	return router
}
