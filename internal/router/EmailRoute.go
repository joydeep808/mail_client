package router

import (
	"myapp/internal/services"

	"github.com/gofiber/fiber/v2"
)



func EmailRouter(app fiber.Router) {
	app.Post("/request" , services.SendEmailToUser)
	app.Post("/request/all" , services.SendEmailToMultipleUser)
}