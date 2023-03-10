package main

import (
	"log"
	"os"
	"time"

	"github.com/RahulMj21/postgres-fiber/models"
	"github.com/RahulMj21/postgres-fiber/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}
	book.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = r.DB.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(201).JSON(&fiber.Map{"status": "success", "message": "book has been added to the DB"})
}
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(200).JSON(&fiber.Map{"status": "success", "data": bookModels})
}
func (r *Repository) GetBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(&fiber.Map{"status": "fail", "message": "id cannot be empty"})
	}

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(200).JSON(&fiber.Map{"status": "success", "data": bookModel})
}
func (r *Repository) UpdateBook(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	bookModel := &models.Books{}

	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(&fiber.Map{"status": "fail", "message": "id cannot be empty"})
	}

	err := r.DB.Delete(bookModel, id).Error
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(200).JSON(&fiber.Map{"status": "success", "message": "item deleted"})
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

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	repository := Repository{
		DB: db,
	}

	app := fiber.New()
	app.Use(logger.New())
	repository.SetupRoutes(app)
	app.Listen(":8000")
}
