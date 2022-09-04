package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-mongo-clean/config"
	"net/http"
)

func main() {
	// Run Database
	database := config.NewMongoDatabase()
	fmt.Println(database)

	// Setup Fiber
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello world!",
		})
	})
	app.Listen(":8081")
}
