package handlers

import (
	"TodoAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) CreateTodo(res http.ResponseWriter, req *http.Request){
  fmt.Fprintln(res, "CreateTodo function is called ")  
  var todo models.Todo
   err := json.NewDecoder(req.Body).Decode(&todo)
  if err != nil{
    http.Error(res,err.Error(),http.StatusBadRequest)
    return
  }

  if todo.Title == ""{
    http.Error(res, "Title is required",http.StatusBadRequest)
    return
  }

  if result:= h.DB.Create(&todo); result.Error != nil{
    fmt.Println(result.Error)
  }

  res.Header().Add("Content-Type","application/json")
  
  res.WriteHeader(http.StatusNoContent)
}

func (h handler) getTodos (res http.ResponseWriter, req *http.Request){
  fmt.Fprintln(res,"get todos function is called")
  
}
