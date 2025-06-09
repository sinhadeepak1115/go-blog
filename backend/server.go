package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello, World!")
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello, World!")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello, World!")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello, World!")
}

func main() {
	connStr := "postgresql://postgres:pass@localhost:5432/blog?sslmode=disable"
	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// adding fiber for HTTP server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
