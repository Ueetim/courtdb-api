package models

import (

)

type Court struct {
	Id			uint	`jsn:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email" gorm:"unique"`
	Password	[]byte	`json:"-"` //indicates we dont want to return password
	City		string	`json:"city"`
}