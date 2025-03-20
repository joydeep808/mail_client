package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/internal/email"
	"myapp/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
  
	app := fiber.New()
	config.DBConnection()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
}
	email.InitializeEmailClient()
	// app.Mount("/users", router.UserRouter(app))
	users := app.Group("/users")
	emails := app.Group("/emails")
	router.EmailRouter(emails)
	router.UserRouter(users)
	fmt.Printf("Server started on port 3000")
    
	app.Listen(":3000")
}