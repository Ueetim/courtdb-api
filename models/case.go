package models

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	CourtID			int		`json:"court_id"`
	Title			string	`json:"title"`
	Description		string	`json:"desc"`
	Created			string	`json:"created"`
	Completed		string	`json:"city"`
	File			string	`json:"doc"`
}