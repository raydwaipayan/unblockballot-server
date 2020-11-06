package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	router "github.com/raydwaipayan/unblockballot-server/router"
)

func main() {
	app := fiber.New()
	app.Use(recover.New(), logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)
	app.Listen(":3000")
}
