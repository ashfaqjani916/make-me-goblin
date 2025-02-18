package main

import (
	"EcommerceAPI/internal/database"
	"EcommerceAPI/internal/handlers"
	"EcommerceAPI/internal/repositories"
	"EcommerceAPI/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main(){
err := godotenv.Load()
  if err != nil{
    log.Fatal("Error loading .env file")
  }

	db := database.InitDB()

	userRepo := repositories.NewUserRepository(db) 
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)




  r :=  mux.NewRouter()
  r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w,"Hello world")
  })

	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/createUser", userHandler.CreateUser).Methods("POST")

	fmt.Println("server is running on port 8080")
  http.ListenAndServe(":8080",r)
}
