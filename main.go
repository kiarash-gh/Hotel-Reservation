package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/kiarash-gh/Hotel-reservation/api"
)

func main() {
	listenAdder := flag.String("listenAdder", ":4000", "The Listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("api/v1")
	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandelGetUsers)
	apiv1.Get("/user/:id", api.HandelGetUser)
	app.Listen(*listenAdder)
}


func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine"})
}
