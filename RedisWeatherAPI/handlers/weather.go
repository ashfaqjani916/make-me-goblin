package handlers

import (
	"fmt"
	"net/http"
)

func GetWeather(res http.ResponseWriter, req *http.Request){
  fmt.Fprintln(res,"GetWeather function is called")

}
