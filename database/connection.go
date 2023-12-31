package database

import (
	"gorm.io/gorm"
    "gorm.io/driver/mysql"

    "github.com/ueetim/court-system/models"
)

var DB *gorm.DB

func Connect() {
    // conn, err := gorm.Open(mysql.Open("uduak:quixote4@/court_db?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	conn, err := gorm.Open(mysql.Open("uduak:quixote4@tcp(faint-banana-cnzs.cf3klctyi2vz.us-west-2.rds.amazonaws.com:3306)/primarydb?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
    if err != nil {
        panic("Could not connect to database")
    }

    DB = conn

    conn.AutoMigrate(&models.Court{}, &models.Case{}, &models.Documents{})
}
