package config

import (
	"fmt"
	"myapp/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConnection()  {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	var err error
	db , err = gorm.Open(postgres.Open(dns) , &gorm.Config{})
	if err != nil {
		// i have to shutdonw 
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.User{}, &model.EmailRecord{})
	if err != nil {
		panic("failed to migrate database")
	}
	println("connected to database")
}


func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetDB() *gorm.DB {
	return db
}


