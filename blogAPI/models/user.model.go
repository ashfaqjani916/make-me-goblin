package models


type User struct{
  id int `json:"user_id" bson:"user_id"`
  username string `json:"username" bson:"username"`
  email string `json:"email" bson:"email"`
}


