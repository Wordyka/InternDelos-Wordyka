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

// EndpointStats holds the count and unique user agents for each endpoint.
type EndpointStats struct {
	Count            int
	UniqueUserAgents map[string]bool
}


// ============= FARMS ===================

// CreateFarm creates a new farm.
func CreateFarm(c *fiber.Ctx) error {
    farm := new(models.Farms)
    if err := c.BodyParser(farm); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Check for duplicate farm name
    var count int64
    database.DB.Db.Model(&models.Farms{}).Where("name = ?", farm.Name).Count(&count)
    if count > 0 {
        return handleError(c, fiber.StatusConflict, "Farm with the same name already exists")
    }

    // Create a new farm
    if err := database.DB.Db.Create(&farm).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error creating farm: "+err.Error())
    }
    return c.Status(fiber.StatusCreated).JSON(farm)
}

// UpdateFarm updates an existing farm by its ID.
func UpdateFarm(c *fiber.Ctx) error {
    farmID := c.Params("id")
    updateData := new(models.Farms)
    if err := c.BodyParser(updateData); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Check if the farm exists
    var farm models.Farms
    if err := database.DB.Db.First(&farm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
    }

    // Update the farm
    if err := database.DB.Db.Model(&farm).Updates(updateData).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error updating farm: "+err.Error())
    }
    return c.Status(fiber.StatusOK).JSON(farm)
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

// ListFarms returns a list of all farms with their associated ponds.
func ListFarms(c *fiber.Ctx) error {
    farms := []models.Farms{}
    database.DB.Db.Preload("Ponds").Find(&farms)
    return c.Status(fiber.StatusOK).JSON(farms)
}

// GetFarmByID returns a specific farm by its ID.
func GetFarmByID(c *fiber.Ctx) error {
    farmID := c.Params("id")
    farm := models.Farms{}

    // Check if the farm exists and preload ponds
    if err := database.DB.Db.Preload("Ponds").First(&farm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(farm)
}


// =============== PONDS ================

// CreatePond creates a new pond with preloading of related farm data.
func CreatePond(c *fiber.Ctx) error {
    pond := new(models.Ponds)
    if err := c.BodyParser(pond); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Ensure the farm exists
    var existingFarm models.Farms
    if err := database.DB.Db.Preload("Farm").First(&existingFarm, pond.FarmID).Error; err != nil {
        return handleError(c, fiber.StatusNotFound, "Farm not found")
    }

    // Avoid duplicate pond entry
    var count int64
    database.DB.Db.Model(&models.Ponds{}).Where("name = ? AND farm_id = ?", pond.Name, pond.FarmID).Count(&count)
    if count > 0 {
        return handleError(c, fiber.StatusConflict, "Duplicate pond entry")
    }

    // Create the pond
    if err := database.DB.Db.Create(&pond).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error creating pond: "+err.Error())
    }


    return c.Status(fiber.StatusCreated).JSON(pond)
}


// UpdatePond updates an existing pond by its ID.
func UpdatePond(c *fiber.Ctx) error {
    pondID := c.Params("id")
    updatedData := new(models.Ponds)
    if err := c.BodyParser(updatedData); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Check pond existence
    var existingPond models.Ponds
    if err := database.DB.Db.Preload("Farm").First(&existingPond, pondID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Pond not found")
        }   
        return handleError(c, fiber.StatusInternalServerError, "Error finding existing pond: "+err.Error())
    }

    // Update the pond
    if err := database.DB.Db.Model(&existingPond).Updates(updatedData).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error updating pond: "+err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(existingPond)
}

// DeletePond soft deletes an existing pond.
func DeletePond(c *fiber.Ctx) error {
    pondID := c.Params("id")
    var pond models.Ponds

    // Check pond existence
    if err := database.DB.Db.First(&pond, pondID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Pond not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding pond: "+err.Error())
    }

    // Soft delete the pond
    database.DB.Db.Delete(&pond)
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pond soft deleted successfully",
    })
}

// ListPonds returns a list of all ponds.
func ListPonds(c *fiber.Ctx) error {
    var ponds []models.Ponds
    database.DB.Db.Preload("Farm").Find(&ponds)
    return c.Status(fiber.StatusOK).JSON(ponds)
}

// GetPondByID returns a specific pond by its ID.
func GetPondByID(c *fiber.Ctx) error {
    pondID := c.Params("id")
    var pond models.Ponds

    // Check pond existence and preload farm
    if err := database.DB.Db.Preload("Farm").First(&pond, pondID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Pond not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding pond: "+err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(pond)
}


// Tracking API Usage Statistics

// StatsMiddleware updates the statistics for each API request.
func StatsMiddleware(statsMap map[string]*EndpointStats) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Update stats
		endpoint := c.Method() + " " + c.Path()
		userAgent := c.Get("User-Agent")

		stats, exists := statsMap[endpoint]
		if !exists {
			stats = &EndpointStats{
				Count:            0,
				UniqueUserAgents: make(map[string]bool),
			}
			statsMap[endpoint] = stats
		}

		stats.Count++
		stats.UniqueUserAgents[userAgent] = true

		// Proceed with the request
		return c.Next()
	}
}


func GetStats(c *fiber.Ctx, statsMap map[string]*EndpointStats) error {
    // Convert stats to desired format
    formattedStats := make(map[string]interface{})
    for endpoint, stats := range statsMap {
        formattedStats[endpoint] = fiber.Map{
            "count":             stats.Count,
            "unique_user_agent": len(stats.UniqueUserAgents),
            // Include other relevant information
        }
    }

    return c.JSON(formattedStats)
}