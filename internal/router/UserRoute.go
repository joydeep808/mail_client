package router

import (
	"myapp/internal/services"
	"github.com/gofiber/fiber/v2"
)


func UserRouter(app fiber.Router) {
	app.Post("/" , services.CreateUser)
	app.Get("/" , services.GetUsers)
	app.Get("/:id" , services.CheckPassword)

}