package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
func (r *Repository) GetBook(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
func (r *Repository) UpdateBook(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/books", r.CreateBook)
	api.Get("/books", r.GetBooks)
	api.Get("/books/:id", r.GetBook)
	api.Put("/books/:id", r.UpdateBook)
	api.Delete("/books/:id", r.DeleteBook)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	repository := Repository{
		// DB: db
	}

	app := fiber.New()
	app.Use(logger.New())
	repository.SetupRoutes(app)
	app.Listen(":8000")
}
