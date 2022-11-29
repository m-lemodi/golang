package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://test:sql@localhost:5432/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", "./html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error { return indexHandler(c, db) })

	app.Post("/", func(c *fiber.Ctx) error { return postHandler(c, db) })

	app.Put("/update", func(c *fiber.Ctx) error { return putHandler(c, db) })

	app.Delete("/delete", func(c *fiber.Ctx) error { return deleteHandler(c, db) })

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM persons")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{
		"Person": todos,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}
