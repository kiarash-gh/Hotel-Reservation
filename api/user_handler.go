package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kiarash-gh/Hotel-reservation/types"
)

func HandelGetUsers(c *fiber.Ctx) error{
	u := types.User{
		FirstName: "Kiarash",
		LastName: "Gh",
	}
	return c.JSON(u)
}


func HandelGetUser(c *fiber.Ctx) error{
	return c.JSON(map[string]string{"user":"Kiarash"})
}