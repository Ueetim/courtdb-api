package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/ueetim/court-system/models"
	"github.com/ueetim/court-system/database"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.Court{
		Name:		data["name"],
		Email:		data["email"],
		Password: 	password,
		City:		data["city"],
	}

	// check if email already exists
	var userExists models.Court
	database.DB.Where("email = ?", data["email"]).First(&userExists)

	// if record exists
	if userExists.ID != 0 {
		// c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "a user with the provided email already exists",
		})
	}

	database.DB.Create(&user)
	
	return c.JSON(user)

}