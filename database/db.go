package database

import (
	"fmt"
	"golang-project/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DB *gorm.DB

func ConnectDB(){
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil{
		log.Fatal("failed to connect to db")
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	fmt.Println("running migration")
	db.AutoMigrate(&models.User{})
	
	DB = db
}