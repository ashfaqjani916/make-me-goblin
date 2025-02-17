package main

import (
	"EcommerceAPI/internal/database"
	"EcommerceAPI/internal/handlers"
	"EcommerceAPI/internal/repositories"
	"EcommerceAPI/internal/services"
 "github.com/google/wire"
)


func InitialiseUserHandler() *handlers.UserHandler {

	 wire.Build(
		database.InitDB,
		repositories.NewUserRepository,
		services.NewUserService,
		handlers.NewUserHandler,
		)
	return &handlers.UserHandler{}
}
