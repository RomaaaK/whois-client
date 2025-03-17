package routes

import (
	"whois-client/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandle *handlers.UserHandler) {
	v1 := router.Group("/api/v1")

	{
		SetupUserRoutes(v1, userHandle)
	}
}
