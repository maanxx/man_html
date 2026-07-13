package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123456xyz dbname=crm_gocas port=5447 sslmode= disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Lỗi kết nối Database: \n", err)
	}

	log.Println("Kết nối Postgres Database thành công!")

	DB = db
}