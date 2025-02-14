package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/savr/database"
	"github.com/savr/models"
	"github.com/savr/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	models.MigrateUser(db)

	routes.UserRoutes(db, app)

	app.Listen(":8080")
}
