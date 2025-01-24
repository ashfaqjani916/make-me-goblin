package services

import (
	"fmt"
	"net/http"
)

func CreateUser(res http.ResponseWriter, req http.Request){
  fmt.Fprintln(res,"Called the create user function")

}
