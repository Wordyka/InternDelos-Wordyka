package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Wordyka/InternDelos-Wordyka/handlers"
)

func setupRoutes(app *fiber.App, statsMap map[string]*handlers.EndpointStats) {
    // Group routes with v1 prefix
    v1 := app.Group("/v1")

    // Farm routes
    v1.Get("/farms", handlers.ListFarms)
    v1.Post("/farm", handlers.CreateFarm)
    v1.Put("/farm/:id", handlers.UpdateFarm)
    v1.Delete("/farm/:id", handlers.DeleteFarm)
    v1.Get("/farm/:id", handlers.GetFarmByID)

    // Pond routes
    v1.Post("/pond", handlers.CreatePond)
    v1.Put("/pond/:id", handlers.UpdatePond)
    v1.Delete("/pond/:id", handlers.DeletePond)
    v1.Get("/pond/:id", handlers.GetPondByID)
    v1.Get("/ponds", handlers.ListPonds)

    // API statistics route
    v1.Get("/stats", func(c *fiber.Ctx) error {
        return handlers.GetStats(c, statsMap)
    })
}
