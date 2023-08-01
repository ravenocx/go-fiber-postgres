package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ritsuhaaa/go-fiber-postgres/dto"
	"github.com/ritsuhaaa/go-fiber-postgres/models"
	"github.com/ritsuhaaa/go-fiber-postgres/storage"
)

func CreateBook(context *fiber.Ctx) error {
	// with fiber.context we can access to the http body
	request := dto.BookRequest{}

	// parse the body to the request
	if err := context.BodyParser(&request); err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing request body",
		})

	}

	// create the book
	book := models.Books{
		Author:    request.Author,
		Title:     request.Title,
		Publisher: request.Publisher,
	}

	// save the book
	err := storage.DB.Create(&book).Error

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not create book",
		})
	}

	// return response
	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "book created successfully",
		"data":    book,
	})

}

func GetBooks(context *fiber.Ctx) error {
	bookModel := []models.Books{}

	err := storage.DB.Find(&bookModel).Error

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not fetch books",
		})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModel,
	})
}

func DeleteBook(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := models.Books{}

	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "id cannot be empty",
		})
	}

	err := storage.DB.Where("id = ?", id).Delete(bookModel).Error

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not delete book",
		})

	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book deleted",
	})

}

func GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "id cannot be empty",
		})
	}

	bookModel := models.Books{}

	err := storage.DB.Where("id = ?", id).First(&bookModel).Error

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not fetch book",
		})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "book fetched successfully",
		"data":    bookModel,
	})
}
