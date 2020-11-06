package handler

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/unblockballot-server/types"
)

// PollCreate create a poll
func PollCreate(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["admin"] != float64(1) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	poll := new(types.Poll)

	if err := c.BodyParser(poll); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "poll created successfully !", "poll": poll})
}
