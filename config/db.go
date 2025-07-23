package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx" 
	_ "github.com/lib/pq"     
)


var DB *sqlx.DB


func ConnectDB() (*sqlx.DB, error) {
	
	db, err := sqlx.Connect("postgres", "user=valentinoanugerah dbname=erp sslmode=disable")
	if err != nil {
		log.Println("Gagal koneksi ke database:", err)
		return nil, err
	}

	
	DB = db

	fmt.Println("Koneksi berhasil ke database PostgreSQL")
	return db, nil
}
