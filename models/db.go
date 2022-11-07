package models

import (
	"fmt"
  "os"
  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"

)

var DB *gorm.DB

func ConnectDatabase() {

  err := godotenv.Load(".env")

  if err!=nil {
    panic("Error loading .env file")
  }

  dbHost:=os.Getenv("DB_HOST")
  dbPort:=os.Getenv("DB_PORT")
  dbUsername:=os.Getenv("DB_USER")
  dbPassword:=os.Getenv("DB_PASS")
  dbName:=os.Getenv("DB_NAME")
  connectionString:=dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"
  
  fmt.Println(connectionString)
  database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Todo{})

  DB = database
}