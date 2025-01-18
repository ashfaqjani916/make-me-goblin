package handlers

import (
	//	"TodoAPI/models"
	"TodoAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) CreateUser(res http.ResponseWriter, req *http.Request){
 
  fmt.Fprintln(res, "create user function is called")
  var user models.User
  err := json.NewDecoder(req.Body).Decode(&user)
  if err != nil{
    http.Error(res,err.Error(),http.StatusBadRequest)
    return
  }

  if user.Name == ""{
    http.Error(res, "name is required",http.StatusBadRequest)
    return
  }

  if result:= h.DB.Create(&user); result.Error != nil{
    fmt.Println(result.Error)
  }

  res.Header().Add("Content-Type","application/json")
  
  res.WriteHeader(http.StatusNoContent)

} 

