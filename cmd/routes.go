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
    v1.Post("/farms/farm", handlers.CreateFarm)
    v1.Put("/farms/farm/:id", handlers.UpdateFarm)
    v1.Delete("/farms/farm/:id", handlers.DeleteFarm)
    v1.Get("/farms/:id", handlers.GetFarmByID)

    // Pond routes
    // v1.Get("/ponds", handlers.ListPonds)
    // v1.Post("/pond", handlers.CreatePond)
    // v1.Put("/pond", handlers.UpdatePond)
    // v1.Delete("/pond/:id", handlers.DeletePond)
    // v1.Get("/pond/:id", handlers.GetPondByID)

    // API statistics route
    // v1.Get("/stats", handlers.APIStatistics)
}
