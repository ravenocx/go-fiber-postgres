package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ritsuhaaa/go-fiber-postgres/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api") // create a group of routes

	// create a route for the group
	api.Post("/book", controllers.CreateBook)
	api.Get("/books", controllers.GetBooks)
	api.Get("/book/:id", controllers.GetBookById)
	api.Delete("/book/:id", controllers.DeleteBook)

}