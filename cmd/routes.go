package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Wordyka/InternDelos-Wordyka/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFarms)

	app.Post("/farms", handlers.CreateFarm)
}
