package routes

import (
	"whois-client/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(group *gin.RouterGroup, h *handlers.UserHandler) {
	group.GET("/users", h.GetUsers)
	group.GET("/users/:id", h.GetUserByID)
	group.POST("/users", h.CreateUser)
	group.PUT("/users/:id", h.UpdateUser)
	group.DELETE("/users/:id", h.DeleteUser)
}
