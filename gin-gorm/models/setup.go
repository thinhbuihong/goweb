package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "buihongthinh"
	dbname   = "godemo"
)

var DB *gorm.DB

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	database, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connec to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
