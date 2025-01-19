package main

import (
	"RedisWeatherAPI/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){
err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  handlers.InitClient()
  r := mux.NewRouter()
  r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w,"welcome to Weather API powered by redis") 
  })
  r.HandleFunc("/getweather/{location}",handlers.GetWeather).Methods(http.MethodGet)
 http.ListenAndServe(":8080",r)
}
