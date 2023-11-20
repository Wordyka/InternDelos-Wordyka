package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Wordyka/InternDelos-Wordyka/database"
	"github.com/Wordyka/InternDelos-Wordyka/models"
)

func ListFarms(c *fiber.Ctx) error {
	farms := []models.Farm{}
	database.DB.Db.Find(&farms)

	return c.Status(200).JSON(farms)
}

func CreateFarm(c *fiber.Ctx) error {
	farm := new(models.Farms)
	if err := c.BodyParser(farm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&farm)

	return c.Status(200).JSON(farm)
}
