package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var creds struct{
  username string
  password string
}

type LoginResponse struct {
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": username,
    "exp": time.Now().Add(time.Hour *24).Unix(),
  })
  
  secret_token := os.Getenv("jwtSecret")
  tokenString, err := token.SignedString(secret_token)

  if err!= nil{
    return "", err
  }

  return tokenString,nil 
}

func Login(res http.ResponseWriter, req *http.Request){
  err := json.NewDecoder(req.Body).Decode(&creds)
  if err!= nil{
    fmt.Fprintln(res,"Error decoding user details")
    return 
  }
 
  //Get the hashed password from the database , write the function in the repository package
  hashedPassword := ""

  if verifyPassword(hashedPassword, creds.password) == nil{

    token, err := CreateToken(creds.username)
    if err != nil{
    http.Error(res, "Error creating token", http.StatusInternalServerError)
		return
    }

    res.WriteHeader(http.StatusOK)
    response := LoginResponse {
       Message: "Login Successfull!!",
       Token : token,
    }
    json.NewEncoder(res).Encode(response)
  }
  
}

func HashPassword(password string) (string , error){

  hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil{
    return "",err
  }
  return string(hashedBytes),nil
}


func verifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
