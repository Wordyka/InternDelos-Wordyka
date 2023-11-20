package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Wordyka/InternDelos-Wordyka/database"
	"github.com/Wordyka/InternDelos-Wordyka/models"
	"gorm.io/gorm"
)

type ErrorResponse struct {
    Message string `json:"message"`
}

func handleError(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(ErrorResponse{Message: message})
}

// func ListFarms(c *fiber.Ctx) error {
// 	farms := []models.Farms{}
// 	database.DB.Db.Find(&farms)

// 	return c.Status(200).JSON(farms)
// }


// CreateFarm creates a new farm.
func CreateFarm(c *fiber.Ctx) error {
    farm := new(models.Farms)
    if err := c.BodyParser(farm); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Check if the farm with the same name already exists
    existingFarm := models.Farms{}
    if err := database.DB.Db.Where("name = ?", farm.Name).First(&existingFarm).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Farm with the same name not found, create a new one
            database.DB.Db.Create(&farm)
            return c.Status(fiber.StatusCreated).JSON(farm)
        }
        return handleError(c, fiber.StatusInternalServerError, "Error checking existing farm: "+err.Error())
    }

    // Farm with the same name found, return conflict response
    return handleError(c, fiber.StatusConflict, "Farm with the same name already exists")
}

// UpdateFarm updates an existing farm by its ID.
func UpdateFarm(c *fiber.Ctx) error {
    farmID := c.Params("id")
    farm := new(models.Farms)
    if err := c.BodyParser(farm); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Check if the farm with the specified ID exists
    existingFarm := models.Farms{}
    if err := database.DB.Db.First(&existingFarm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding existing farm: "+err.Error())
    }

    // Update the existing farm
    existingFarm.Name = farm.Name
    database.DB.Db.Save(&existingFarm)
    return c.Status(fiber.StatusOK).JSON(existingFarm)
}

// DeleteFarm soft deletes an existing farm.
func DeleteFarm(c *fiber.Ctx) error {
    farmID := c.Params("id")
    farm := models.Farms{}

    // Check if the farm exists
    if err := database.DB.Db.First(&farm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
    }

    // Soft delete the farm
    database.DB.Db.Delete(&farm)
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Farm soft deleted successfully",
    })
}

// ListFarms returns a list of all farms.
func ListFarms(c *fiber.Ctx) error {
    farms := []models.Farms{}
    database.DB.Db.Find(&farms)
    return c.Status(fiber.StatusOK).JSON(farms)
}

// GetFarmByID returns a specific farm by its ID.
func GetFarmByID(c *fiber.Ctx) error {
    farmID := c.Params("id")
    farm := models.Farms{}

    // Check if the farm exists
    if err := database.DB.Db.First(&farm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(farm)
}