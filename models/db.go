package models

import (
  "os"
  "github.com/joho/godotenv"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {

  err := godotenv.Load(".env")

  if err!=nil {
    panic("Error loading .env file")
  }

  dbHost:=os.Getenv("DB_HOST")
  dbPort:=os.Getenv("DB_PORT")
  dbUsername:=os.Getenv("DB_USERNAME")
  dbPassword:=os.Getenv("DB_PASSWORD")
  dbName:=os.Getenv("DB_NAME")
  connectionString:=dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName

  database, err := gorm.Open("mysql", connectionString)

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Todo{})

  DB = database
}