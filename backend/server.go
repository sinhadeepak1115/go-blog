package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// makeing handlers for each HTTP method

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var title string
	var posts []string

	rows, err := db.Query("SELECT title FROM posts")
	if err != nil {
		log.Println("Error querying database:", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&title); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		posts = append(posts, title)
	}

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	var input struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString("Invalid input")
	}
	if input.Title == "" {
		return c.Status(400).SendString("Title is required")
	}

	_, err := db.Exec("INSERT INTO posts (title) VALUES ($1)", input.Title)
	if err != nil {
		log.Println("Error inserting into DB:", err)
		return c.Status(500).SendString("Failed to create post")
	}

	return c.SendStatus(201)
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
	defer db.Close()

	// adding fiber for HTTP server
	app := fiber.New()

	app.Static("/", "../frontend")

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	// Get all posts
	app.Get("/api/posts", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT title FROM posts ORDER BY created_at DESC")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "DB error"})
		}
		defer rows.Close()

		var posts []string
		for rows.Next() {
			var title string
			rows.Scan(&title)
			posts = append(posts, title)
		}
		return c.JSON(fiber.Map{"posts": posts})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	// Add new post
	app.Post("/api/posts", func(c *fiber.Ctx) error {
		var data struct {
			Title string `json:"title"`
		}
		if err := c.BodyParser(&data); err != nil || data.Title == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		_, err := db.Exec("INSERT INTO posts(title) VALUES($1)", data.Title)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Insert failed"})
		}

		return c.SendStatus(201)
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
