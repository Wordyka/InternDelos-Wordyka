package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Wordyka/InternDelos-Wordyka/database"
	"github.com/Wordyka/InternDelos-Wordyka/handlers"
)




func main() {
	database.ConnectDb()

	app := fiber.New()

	// Initialize stats map
	statsMap := make(map[string]*handlers.EndpointStats)

	// Use stats middleware
	app.Use(handlers.StatsMiddleware(statsMap))

	// Setup other routes
	setupRoutes(app, statsMap)

	app.Listen(":3000")
}
