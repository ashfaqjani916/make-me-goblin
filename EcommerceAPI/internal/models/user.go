package models

type User struct{
  user_id uint `json:"user_id"`
  user_type string `json:"user_type"`
  user_name string `json:"user_name"`
  password_hash string `json:"password_hash"`
  jwt_token string `json:"jwt_token"`
}
