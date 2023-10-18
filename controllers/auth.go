package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/ueetim/court-system/database"
	"github.com/ueetim/court-system/models"
	"github.com/ueetim/court-system/middleware"
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
		Location:	data["location"],
		Type:		data["type"],
		Email:		data["email"],
		Password: 	password,
	}

	// check if email already exists
	var userExists models.Court
	database.DB.Where("email = ?", data["email"]).First(&userExists)

	// if record exists
	if userExists.ID != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "A user with the provided email already exists",
		})
	}

	database.DB.Create(&user)
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Account created successfully. Please log in",
	})

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
		SameSite:	"None",
	}

	c.Cookie(&cookie)

	c.Status(fiber.StatusAccepted)
	return c.JSON(fiber.Map{
		"token": 	&cookie,
		"expires":	&cookie.Expires,
	})
}

func GetLoggedInUser(c *fiber.Ctx) error {
	_, claims := middleware.AuthenticateUser(c)	

	var court models.Court

	database.DB.Where("id = ?", claims).First(&court)

	return c.JSON(court)
}

func Logout(c *fiber.Ctx) error {
	// create a cookie and set the expiration to a time in the past
	cookie := fiber.Cookie{
		Name: 		"jwt",
		Value: 		"",
		Expires: 	time.Now().Add(-time.Hour),
		HTTPOnly:	true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logged out successfully",
	})
}
