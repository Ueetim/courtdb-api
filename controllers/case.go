package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ueetim/court-system/database"
	"github.com/ueetim/court-system/models"
)

func CreateRecord(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	courtId, err := strconv.Atoi(data["court_id"])
	if err != nil {
		return err
	}

	record := models.Case {
		CourtID:		courtId,
		Title:			data["title"],
		Description:	data["description"],
		Created:		data["created"],
		Completed:		data["completed"],
	}

	database.DB.Create(&record)
	
	return c.JSON(record)
}

func GetRecordsByUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var records []models.Case

	database.DB.Where("court_id = ?", data["id"]).Find(&records)

	if len(records) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "nothing found",
		})
	}

	return c.JSON(records)
}

func GetOneRecordByUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var record models.Case

	database.DB.Where("id = ? AND court_id = ?", data["id"], data["court_id"]).First(&record)

	if record.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "nothing found",
		})
	}

	return c.JSON(record)
}