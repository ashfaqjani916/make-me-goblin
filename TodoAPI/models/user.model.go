package models

type User struct{
  id int `json:"id" gorm:"primaryKey"`
  name string `json:"name"`
  email string `json:"email"`
    
}
