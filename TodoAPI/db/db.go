package db

import (
	"TodoAPI/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB() *gorm.DB{
  url := os.Getenv("dbURL")  

  db, err := gorm.Open(postgres.Open(url),&gorm.Config{})

  if err != nil{
    log.Fatalln(err)
  }
  
  log.Println("database connected successfully")
  db.AutoMigrate(&models.User{})

  return db
}  
