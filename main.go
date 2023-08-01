package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ritsuhaaa/go-fiber-postgres/routes"
	"github.com/ritsuhaaa/go-fiber-postgres/storage"
)

func main() {
	err := godotenv.Load(".env") // import the .env file
	if err != nil {
		panic("Error loading .env file")
	}

	// connect to the database
	storage.ConnectDB(&storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMODE:  os.Getenv("DB_SSLMODE"),
	})

	// initialize the fiber app
	app := fiber.New()

	// setup the routes
	routes.Setup(app)

	// listen to port 3000
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
