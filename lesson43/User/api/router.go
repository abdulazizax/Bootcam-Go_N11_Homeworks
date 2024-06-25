package api

import (
	h "user/api/handler"
	"user/storage"

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
	crud.POST("/user/post", handler.CreateUser)

	crud.GET("/user/get", handler.GetUsers)

	crud.GET("/user/get/:id", handler.GetUserByID)

	crud.GET("/user_name/get/:name", handler.GetUserByName)

	crud.PUT("/user/put/:id", handler.UpdateUserByID)

	crud.DELETE("/user/delete/:id", handler.DeleteUserByID)

	return router
}
