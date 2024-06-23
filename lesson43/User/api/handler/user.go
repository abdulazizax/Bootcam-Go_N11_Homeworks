package handler

import (
	m "user/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(c *gin.Context) {
	var newUser m.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.User().CreateUser(c, newUser)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}

func (h *handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.storage.User().GetUserById(c, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h *handler) GetUsers(c *gin.Context) {
	res, err := h.storage.User().GetUsers(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, res)
}

func (h *handler) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")

	var newUser m.UserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.storage.User().UpdateUserById(c, id, newUser)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)

}

func (h *handler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	err := h.storage.User().DeleteUserByID(c, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User deleted!",
	})
}

func (h *handler) GetUserByName(c *gin.Context) {
	name := c.Param("name")

	res, err := h.storage.User().GetUserById(c, name)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, res)
}
