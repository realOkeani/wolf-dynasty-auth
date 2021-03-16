package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realOkeani/wolf-dynasty-auth/http"
	"github.com/realOkeani/wolf-dynasty-auth/sql"
)

func main() {
	sql.Connect()

	app := fiber.New()

	http.Setup(app)

	app.Listen(":8000")
}
