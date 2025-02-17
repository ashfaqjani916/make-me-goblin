package models

type User struct{
  User_id uint `json:"user_id"`
  User_type string `json:"user_type"`
  User_name string `json:"user_name"`
  Password_hash string `json:"password_hash"`
  Jwt_token string `json:"jwt_token"`
}
