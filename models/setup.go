package models

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase(){
  dsn := "host=localhost user=steven password=password dbname=fullstack_api port=5432 sslmode=disable"
  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database")
  }
  
  database.AutoMigrate(&Quest{})

  DB = database
}
