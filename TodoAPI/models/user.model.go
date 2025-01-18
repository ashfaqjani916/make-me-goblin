package models

type User struct{
  id int `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Email string `json:"email"`
    
}
