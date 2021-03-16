package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realOkeani/wolf-dynasty-auth/models"
	"github.com/realOkeani/wolf-dynasty-auth/sql"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//Get all the request data we will send
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	yahooUser := models.YahooUser{}

	if err := sql.DB.Where("email = ?", data["email"]).First(&yahooUser); err != nil {
		return c.SendFile("Email not associated with league")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	sql.DB.Create(&user)

	return c.JSON(user)
}
