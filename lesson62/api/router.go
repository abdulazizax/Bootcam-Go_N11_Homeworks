package api

import (
	"csb/api/handler"
	"csb/internal/storage/acl"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type API struct {
	enforcer *casbin.Enforcer
}

func New() (*API, error) {
	enforcer, err := acl.NewEnforcer()
	if err != nil {
		return nil, err
	}
	return &API{enforcer: enforcer}, nil
}

func (a *API) Run() *gin.Engine {
	router := gin.Default()

	router.Use(a.AuthMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handler.NewHandler()

	user := router.Group("user")
	{
		user.POST("/create", a.CheckPermission("write"), handler.CreateUser)
		user.GET("/get", handler.GetUser)
		user.PUT("/update", a.CheckPermission("write"), handler.UpdateUser)
		user.DELETE("/delete", a.CheckPermission("write"), handler.DeleteUser)
	}

	return router
}

func (a *API) AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user := c.GetHeader("User")
        fmt.Printf("Received User header: %s\n", user) // Debug log
        if user == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Set("user", user)
        fmt.Printf("Set user in context: %v\n", user) // Debug log
        c.Next()
    }
}

func (a *API) CheckPermission(action string) gin.HandlerFunc {
    return func(c *gin.Context) {
        user, _ := c.Get("user")
        userStr, ok := user.(string)
        if !ok {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
            return
        }

        allowed, err := a.enforcer.Enforce(userStr, "users", action)
        fmt.Printf("Permission check: user=%s, obj=users, act=%s, allowed=%v\n", userStr, action, allowed) // Debug log
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error checking permissions"})
            return
        }
        if !allowed {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            return
        }
        c.Next()
    }
}
