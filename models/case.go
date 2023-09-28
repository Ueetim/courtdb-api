package models

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	RecordID		string	`json:"record_id"`
	CourtID			int		`json:"court_id"`
	Title			string	`json:"title"`
	Description		string	`json:"description"`
	Created			string	`json:"created"`
	Completed		string	`json:"completed"`
	Documentation	string	`json:"documentation"`
	Status			string	`json:"status"`
	Visibility		string	`json:"visibility"`
}