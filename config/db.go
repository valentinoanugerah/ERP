package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "user=valentinoanugerah dbname=erp sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Gagal koneksi ke database")
		return nil, err
	}

	fmt.Println("Koneksi berhasil!")
	DB = db
	return db, nil
}
