package handlers

import (
	"TodoAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
  res.WriteHeader(http.StatusCreated) 
  res.WriteHeader(http.StatusNoContent)
}

func (h handler) GetTodos (res http.ResponseWriter, req *http.Request){
    fmt.Fprintln(res,"GetTodos function is called ")
    var todos[] models.Todo
    if result := h.DB.Find(&todos); result.Error != nil {
    http.Error(res,result.Error.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Add("Content-Type","application/json")
  res.WriteHeader(http.StatusOK)
  json.NewEncoder(res).Encode(todos)
}



func (h handler) DeleteTodo(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "DeleteTodo function is called")

    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(res, "Invalid ID format", http.StatusBadRequest)
        return
    }

    if result := h.DB.Delete(&models.Todo{}, id); result.Error != nil {
        http.Error(res, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    res.WriteHeader(http.StatusNoContent)
}


func (h handler) EditTodo(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "EditTodo function is called")

    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(res, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var updatedTodo models.Todo
    err = json.NewDecoder(req.Body).Decode(&updatedTodo)
    if err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        return
    }

    var todo models.Todo
    if result := h.DB.First(&todo, id); result.Error != nil {
        http.Error(res, "Todo not found", http.StatusNotFound)
        return
    }

    todo.Title = updatedTodo.Title
    todo.Description = updatedTodo.Description

    if result := h.DB.Save(&todo); result.Error != nil {
        http.Error(res, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    res.Header().Add("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)
    json.NewEncoder(res).Encode(todo)
}
