package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/ueetim/court-system/database"
	"github.com/ueetim/court-system/models"
)

const SecretKey = "secret"

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

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Court

	database.DB.Where("email = ?", data["email"]).First(&user)

	// if nothing found
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// generate jwt
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)), //convert user.Id to int, then to string
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name: 		"jwt",
		Value: 		token,
		Expires:	time.Now().Add(time.Hour * 24),
		HTTPOnly: 	true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}