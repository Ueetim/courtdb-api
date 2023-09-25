package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

func AuthenticateUser(c *fiber.Ctx) (error, *string) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		}), nil
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return nil, &claims.Issuer
}