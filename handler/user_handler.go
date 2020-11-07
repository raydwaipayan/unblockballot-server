package handler

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/gofiber/fiber/v2"
	"github.com/raydwaipayan/unblockballot-server/types"
	models "github.com/raydwaipayan/unblockballot-server/models"
)

var passowrdKey = []byte("pass_secret")


//Register user registration handler
func Register(c *fiber.Ctx) error {
	u := new(types.User)

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	// Hashing password using bcrypt 
	pass := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
	}
	u.Password = string(hashedPassword)

	log.Println(u.Role)
	if err:= u.Create(models.DBConfigURL); err!=nil {
		log.Println(err)
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

	// Checking if the user exists
	doesExist, err := u.CheckUserExists(models.DBConfigURL) 
	if !doesExist {
		return c.SendStatus(fiber.StatusForbidden)
	}
	log.Println(u)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}

// Update user data
func Update(c *fiber.Ctx) error {
	u := new(types.User)

	if err := c.BodyParser(u); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err:= u.Update(models.DBConfigURL); err!=nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "user updated"})
}

//PollSubmit user poll submission
// func PollSubmit(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	// log.Println(claims["email"])
// 	return c.JSON(fiber.Map{"message": "test"})
// }
