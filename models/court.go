package models

import "gorm.io/gorm"

type Court struct {
	gorm.Model
	Name		string	`json:"name"`
	Email		string	`json:"email" gorm:"unique"`
	Password	[]byte	`json:"-"` //indicates we dont want to return password
	City		string	`json:"city"`
}