package main

import(
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "log"
  "blogAPI/db"
  //"blogAPI/modules"
)

func main(){
   db.Connect()
  r := mux.NewRouter()
  r.HandleFunc("/",welcome)
 // r.HandleFunc("GET /users",  modules.getusers)
  http.ListenAndServe(":8080",r)
}
func welcome(res http.ResponseWriter, req *http.Request){
  fmt.Fprintf(res, "Welcome to the blog api")
  log.Println("the server is running on http://localhost:8080")
}
 
