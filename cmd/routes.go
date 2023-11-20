package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Wordyka/InternDelos-Wordyka/handlers"
)

func setupRoutes(app *fiber.App) {
    // Group routes with v1 prefix
    v1 := app.Group("/v1")

    // Farm routes
    v1.Get("/farms", handlers.ListFarms)
    v1.Post("/farm", handlers.CreateFarm)

 
}
