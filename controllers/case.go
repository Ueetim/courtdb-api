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
		Visibility: 	"public",
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

func GetRecordsByOtherUsers(c *fiber.Ctx) error {
	_, claims := middleware.AuthenticateUser(c)

	var records []models.Case

	database.DB.Where("court_id <> ? AND visibility = ?", claims, "public").Find(&records)

	if len(records) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "nothing found",
		})
	}

	return c.JSON(records)
}

func UpdateVisibility(c *fiber.Ctx) error {
	var data map[string]string
	
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	_, claims := middleware.AuthenticateUser(c)

	var record models.Case

	// database.DB.Model(record{}).Where("id = ? AND court_id = ?", data["id"], claims).Update(record{"visibility", data["visibility"]})

	database.DB.Where("id = ? AND court_id = ?", data["id"], claims).First(&record).Update("visibility", data["visibility"])

	if record.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Request invalid",
		})
	}

	return c.JSON(record)
}

func EditDocumentation(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	_, claims := middleware.AuthenticateUser(c)

	var record models.Case

	// database.DB.Model(record{}).Where("id = ? AND court_id = ?", data["id"], claims).Update(record{"visibility", data["visibility"]})

	database.DB.Where("id = ? AND court_id = ?", data["id"], claims).First(&record).Update("documentation", data["documentation"])

	if record.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Request invalid",
		})
	}

	return c.JSON(record)
}

func UpdateRecord(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	_, claims := middleware.AuthenticateUser(c)

	var status string
	
	if data["completed"] == "" {
		status = "Open"
	} else {
		status = "Closed"
	}

	var record models.Case

	database.DB.Where("id = ? AND court_id = ?", data["ID"], claims).First(&record).Updates(map[string]interface{}{"record_id": data["record_id"], "title": data["title"], "description": data["description"], "created": data["created"], "completed": data["completed"], "status": status})

	if record.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Could not update record",
		})
	}

	return c.JSON(record)
}