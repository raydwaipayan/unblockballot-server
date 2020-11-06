package handler

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/unblockballot-server/types"
)

//Register user registration handler
func Register(c *fiber.Ctx) error {
	u := new(types.User)

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

//Login user login handler
func Login(c *fiber.Ctx) error {
	u := new(types.User)

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["firstname"] = u.FirstName
	claims["lastname"] = u.LastName
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
