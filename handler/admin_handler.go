package handler

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/unblockballot-server/types"
)

// PollCreate create a poll
func PollCreate(c *fiber.Ctx) error {
	poll := new(types.Poll)

	if err := c.BodyParser(poll); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "poll creates successfully !", "poll": poll})
}
