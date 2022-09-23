package database

import (
	"log"

	"github.com/bandrade/golang-simple-restapi-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDB() {
	con := "user=admin dbname=students password=Passw0rd host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(con))
	if err != nil {
		log.Panic("db connection error: ", err)
	}
	DB.AutoMigrate(&models.Student{})
}
