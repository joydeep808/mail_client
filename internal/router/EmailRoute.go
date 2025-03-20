package router

import (
	"myapp/internal/services"
	"github.com/gofiber/fiber/v2"
)



func EmailRouter(app fiber.Router) {
	app.Post("/" , services.SendEmailToUser)
	app.Post("/all" , services.SendEmailToMultipleUser)
}