package main

import (
	"react-go-jwt/database"
	"react-go-jwt/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}