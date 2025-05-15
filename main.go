package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Sajikan file statis dari folder public
	app.Static("/", "./public")

	// Jalankan server di port 3000
	app.Listen(":3000")
}
