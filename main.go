package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/ritsuhaaa/go-fiber-postgres/models"
	"github.com/ritsuhaaa/go-fiber-postgres/storage"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"` //  json -> "author" : Author
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB // it gives ability to interact to database
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	// with fiber.context we can access to the http body
	book := Book{}

	err := context.BodyParser(&book)
	// binds the request body to a struct (decode json to struct that we want)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON( // HTTP status for the response
			&fiber.Map{"message": "request failed"}) // we got http request from user and response it if failed
		// Map is a shortcut for map[string]interface{}, useful for JSON return
		return err
	}

	err = r.DB.Create(&book).Error // insert into database
	if err != nil {
		context.Status(http.StatusBadRequest).JSON( // http response jika gagal insert
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{ // http response jika berhasil
		"message": "book has been added"})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error // find the book inside the database
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id") // get parameter from request http

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	err := r.DB.Delete(bookModel, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete book"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book deleted successfully"})

	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	fmt.Println("the id is", id)

	err := r.DB.Where("id = ?", id).Find(bookModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book fetched",
		"data":    bookModel,
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) { // it needed access to repository for http routes request
	api := app.Group("/api")                // all the api will start with path (prefix) -> /api
	api.Post("/create_books", r.CreateBook) // r.Method -> it is not calling function but calling the struct method
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env") // import the .env file
	if err != nil {
		panic("Error loading .env file")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config) // connect to postgres
	// config file -> .env
	if err != nil {
		panic("Could not load the database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		panic("Could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New() // create fiber instance
	r.SetupRoutes(app)
	app.Listen(":8080") // Listen serves HTTP requests from the given addr.
}
