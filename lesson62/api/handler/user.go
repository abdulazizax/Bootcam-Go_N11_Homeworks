package handler

import (
	"csb/internal/models"
	"csb/internal/storage/redis"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *redis.RedisRepository
}

func NewHandler() *Handler {
	return &Handler{userService: redis.NewRedisRepository()}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, resp)
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetUser(c.Request.Context(), &models.GetUserRequest{ID: id})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.UpdateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userService.UpdateUser(c.Request.Context(), &user)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.userService.DeleteUser(c.Request.Context(), &models.DeleteUserRequest{ID: id})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}
