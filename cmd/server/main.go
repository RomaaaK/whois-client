package main

import (
	"whois-client/internal/handlers"
	"whois-client/internal/repository"
	"whois-client/internal/routes"
	"whois-client/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserService(userService)

	router := gin.Default()

	routes.SetupRoutes(router, userHandler)

	router.Run(":8080")
}
