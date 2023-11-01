package models

import "gorm.io/gorm"

type Documents struct {
	gorm.Model
	Filename	string	`json:"filename"`
	Filepath	string	`json:"filepath"`
	CaseId		string	`json:"case_id"`
}