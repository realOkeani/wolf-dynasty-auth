package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realOkeani/wolf-dynasty-auth/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/v1/register", controllers.Register)
}
