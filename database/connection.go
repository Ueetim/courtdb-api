package database

import (
	"gorm.io/gorm"
    "gorm.io/driver/mysql"

    "github.com/ueetim/court-system/models"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("mysql://uduak:quixote4@tcp(faint-banana-cnzs.cf3klctyi2vz.us-west-2.rds.amazonaws.com:3306)/primarydb"), &gorm.Config{})
    if err != nil {
        panic("Could not connect to database")
    }

    DB = conn

    conn.AutoMigrate(&models.Court{}, &models.Case{})
}
