package models

import "gorm.io/gorm"

type Court struct {
	gorm.Model
	Name		string	`json:"name"`
	Location	string	`json:"location"`
	Type		string	`json:"type"`
	Email		string	`json:"email" gorm:"unique"`
	Password	[]byte	`json:"-"`
}