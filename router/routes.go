package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	jwtware "github.com/gofiber/jwt/v2"
	handler "github.com/raydwaipayan/unblockballot-server/handler"
)

//SetupRoutes initiates the fiber router
func SetupRoutes(app *fiber.App) {
	user := app.Group("/user", logger.New(), limiter.New())

	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)
	user.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    "test",
	}))

}
