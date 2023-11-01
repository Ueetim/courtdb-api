package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/ueetim/court-system/database"
	"github.com/ueetim/court-system/models"
)

type Data struct {
	Id   int                   `json:"id"`
	File *multipart.FileHeader `json:"file"`
}

func UploadFile(c *fiber.Ctx) error {
	var data Data

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// get original filename
	originalFilename := file.Filename

	// generate random filename
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	ext := filepath.Ext(file.Filename)
	filename := hex.EncodeToString(randBytes) + ext

	file.Filename = filename

	// save locally
	c.SaveFile(file, "public/uploads/"+file.Filename)

	// save to db
	newFile := models.Documents{
		Filename:	originalFilename,
		Filepath: 	"public/" + file.Filename,
		CaseId:   	c.FormValue("case_id"),
	}

	database.DB.Create(&newFile)

	return c.JSON(newFile)
}

func GetCaseFiles(c *fiber.Ctx) error {
	caseId := c.Params("id")

	var docs []models.Documents

	database.DB.Where("case_id = ? AND deleted_at IS NULL", caseId).Find(&docs)

	return c.JSON(docs)
}