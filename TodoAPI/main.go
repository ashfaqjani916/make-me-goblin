package main

import (
	"TodoAPI/db"
	"TodoAPI/handlers"
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
   DB := db.InitDB()
  h := handlers.New(DB)
  r := mux.NewRouter()
  r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w,"Welcome to api")
  }) 
  r.HandleFunc("/newuser",h.CreateUser).Methods(http.MethodPost)
  r.HandleFunc("/deleteuser/{id}",h.DeleteUser).Methods("DELETE")
  r.HandleFunc("/newtodo",h.CreateTodo).Methods(http.MethodPost)
  r.HandleFunc("/todos",h.GetTodos).Methods(http.MethodGet)
  r.HandleFunc("/deletetodo/{id}",h.DeleteTodo).Methods(http.MethodDelete)
  r.HandleFunc("/updatetodo",h.EditTodo).Methods(http.MethodPut)
  http.ListenAndServe(":8080",r)
}
