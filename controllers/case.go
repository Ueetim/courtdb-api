package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/ueetim/court-system/database"
	"github.com/ueetim/court-system/models"
	"github.com/ueetim/court-system/middleware"
)

func CreateRecord(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	_, claims := middleware.AuthenticateUser(c)
	claimString := *claims

	courtId, err := strconv.Atoi(claimString)
	if err != nil {
		return err
	}

	var status string
	
	if data["completed"] == "" {
		status = "Open"
	} else {
		status = "Closed"
	}

	record := models.Case {
		CourtID:		courtId,
		RecordID:		data["record_id"],
		Title:			data["title"],
		Description:	data["description"],
		Created:		data["created"],
		Completed:		data["completed"],
		Status:			status,
	}

	database.DB.Create(&record)
	
	return c.JSON(record)
}

func GetRecordsByUser(c *fiber.Ctx) error {
	_, claims := middleware.AuthenticateUser(c)

	var records []models.Case

	database.DB.Where("court_id = ?", claims).Find(&records)

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