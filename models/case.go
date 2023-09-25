package models

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	RecordID		string	`json:"record_id"`
	CourtID			int		`json:"court_id"`
	Title			string	`json:"title"`
	Description		string	`json:"desc"`
	Created			string	`json:"created"`
	Completed		string	`json:"city"`
	Documentation	string	`json:"doc"`
	Status			string	`json:"status"`
}