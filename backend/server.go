package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()

	app.Get("/", indexHandler)

	app.Post("/", postHandler)

	app.Put("/", putHandler)

	app.Delete("/", deleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
