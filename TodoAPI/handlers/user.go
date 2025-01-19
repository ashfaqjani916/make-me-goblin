package handlers

import (
	//	"TodoAPI/models"
	"TodoAPI/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (h handler) DeleteUser (res http.ResponseWriter, req *http.Request){
  fmt.Fprintln(res,"DeleteUser function is called ")
  vars := mux.Vars(req)
  id,err := strconv.Atoi(vars["id"])
  if err != nil{
    http.Error(res, "User not found",http.StatusBadRequest)
    return 
  }

  if result := h.DB.Delete(&models.User{},id); result.Error != nil{
    http.Error(res,result.Error.Error(),http.StatusInternalServerError)
    return
  }

  res.WriteHeader(http.StatusNoContent)
  fmt.Fprintln(res,"The user has been deleted")
}
