package handler

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/unblockballot-server/models"
	"github.com/raydwaipayan/unblockballot-server/types"
)

// PollCreate create a poll
func PollCreate(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	log.Println(claims["role"])
	if claims["role"] != float64(1) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	poll := new(types.PollBody)

	// Checking if the org already exists

	if err := c.BodyParser(poll); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err:= poll.CreatePoll(models.DBConfigURL); err !=nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"message": "poll created successfully !", "poll": poll})
}
